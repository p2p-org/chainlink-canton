package sequences

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Masterminds/semver/v3"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ds "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/go-daml/pkg/types"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/common"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/executor"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/feequoter"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/mcms"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/splice/splice_api_token_holding_v1"
	"github.com/smartcontractkit/chainlink-canton/contracts"
	executor2 "github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/executor"
	feequoterop "github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/fee_quoter"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/global_config"
	dsutils "github.com/smartcontractkit/chainlink-canton/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-canton/deployment/utils/operations/contract"
)

func appendProposalExercise(proposalDriven bool, outs []contract.ExerciseOutput, o contract.ExerciseOutput) []contract.ExerciseOutput {
	if !proposalDriven {
		return outs
	}

	return append(outs, o)
}

// cantonFeeQuoterUSDPerUnitGas formats V2Params.USDPerUnitGas for Canton FeeQuoter UpdatePrices.
// DAML stores this as Decimal (CCIP.FeeQuoterTypes.GasPriceUpdate). chainlink-ccip models it as *big.Int
// for cross-family tooling; on Canton that integer is scaled by 1e10 USD per gas unit (integration
// parity: 38 -> 0.0000000038, matching historical ApplyFeeTokenUpdates+UpdatePrices tests).
func cantonFeeQuoterUSDPerUnitGas(v *big.Int) types.NUMERIC {
	if v == nil || v.Sign() == 0 {
		return types.NUMERIC("0")
	}
	const scale int64 = 10_000_000_000 // 1e10
	r := new(big.Rat).SetFrac(new(big.Int).Set(v), big.NewInt(scale))
	s := strings.TrimRight(strings.TrimRight(r.FloatString(20), "0"), ".")
	if s == "" || s == "-" {
		return types.NUMERIC("0")
	}

	return types.NUMERIC(s)
}

// feeQuoterInstanceAndRef resolves the FeeQuoter instance for direct ledger calls.
// MCMS proposals require a labeled AddressRef (raw instance id); pinned chainlink-ccip
// CantonLaneConfig does not carry FeeQuoter refs, so ref is often empty — callers skip
// FeeQuoter MCMS exercises in that case.
func feeQuoterInstanceAndRef(source *lanes.ChainDefinition) (contracts.InstanceAddress, cldf_ds.AddressRef, error) {
	inst := contracts.BytesToInstanceAddress(source.FeeQuoter)

	return inst, cldf_ds.AddressRef{}, nil
}

// hardcodedFeeQuoterMCMSRef is a temporary shim until lane input carries a labeled FeeQuoter AddressRef
// (e.g. from topology / datastore). Values match domains/ccv/staging_testnet/datastore/address_refs.json.
func hardcodedFeeQuoterMCMSRef(chainSelector uint64) (cldf_ds.AddressRef, bool) {
	const (
		stagingCantonFeeQuoterChainSelector uint64 = 10109143320554840099
		stagingCantonFeeQuoterAddress              = "0x3891327bf89b1621f67a720f73f8478777f2c106d95e570c5fa388f138bc0728"
		stagingCantonFeeQuoterRawLabel            = "feequoter-shywn@ccipOwner::1220644bd9e52834e8fba90d4607beed37b65991cc2b5377d5d40d07d3db36d4ed51"
	)
	if chainSelector != stagingCantonFeeQuoterChainSelector {
		return cldf_ds.AddressRef{}, false
	}

	return cldf_ds.AddressRef{
		Address:       stagingCantonFeeQuoterAddress,
		Labels:        cldf_ds.NewLabelSet(stagingCantonFeeQuoterRawLabel),
		ChainSelector: chainSelector,
	}, true
}

func cantonMCMSChoiceInput[ARGS any](
	contractRef cldf_ds.AddressRef,
	instanceAddress contracts.InstanceAddress,
	args ARGS,
	proposalDriven bool,
) (contract.ChoiceInput[ARGS], error) {
	if proposalDriven {
		raw, err := dsutils.GetRawInstanceAddressFromAddressRef(contractRef)
		if err != nil {
			return contract.ChoiceInput[ARGS]{}, fmt.Errorf("MCMS raw instance address from ref: %w", err)
		}

		return contract.ChoiceInput[ARGS]{
			InstanceAddress:    instanceAddress,
			RawInstanceAddress: raw.String(),
			Args:               args,
			MCMSEnabled:        true,
		}, nil
	}

	return contract.ChoiceInput[ARGS]{
		InstanceAddress: instanceAddress,
		Args:            args,
		MCMSEnabled:     false,
	}, nil
}

func mergeCantonLaneProposalBatches(
	_ uint64,
	proposalOutputs []contract.ExerciseOutput,
	committeeExtra []mcms_types.BatchOperation,
) (sequences.OnChainOutput, error) {
	mainBatch, err := contract.NewBatchOperationFromExercises(proposalOutputs)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("building MCMS batch from lane exercises: %w", err)
	}

	var batchOps []mcms_types.BatchOperation
	if len(mainBatch.Transactions) > 0 {
		batchOps = append(batchOps, mainBatch)
	}
	batchOps = append(batchOps, committeeExtra...)

	return sequences.OnChainOutput{BatchOps: batchOps}, nil
}

var ConfigureLaneLegAsSource = operations.NewSequence(
	"CantonConfigureLaneLegAsSource",
	semver.MustParse("2.0.0"),
	"Configures a lane leg as source on CCIP 2.0.0",
	func(b operations.Bundle, deps chain.BlockChains, input lanes.UpdateLanesInput) (output sequences.OnChainOutput, err error) {
		b.Logger.Infof("Canton Configuring lane leg as source. src: %+v, dest: %+v", input.Source, input.Dest)

		chain, ok := deps.CantonChains()[input.Source.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.Source.Selector)
		}

		sourceChain := input.Source
		destChain := input.Dest
		// HACK: always collect MCMS proposal transactions for this sequence (ignores input.ProposalDriven).
		pd := true

		if sourceChain.CantonLaneConfig == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("CantonLaneConfig is required for proposal-driven Canton lane configuration (source chain %d)", sourceChain.Selector)
		}
		globalConfigAddress, err := dsutils.ToInstanceAddress(sourceChain.CantonLaneConfig.GlobalConfig)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("global config instance address: %w", err)
		}

		isEnabled := len(destChain.Router) > 0
		defaultExecutor, err := dsutils.GetRawInstanceAddressFromAddressRef(sourceChain.DefaultExecutor)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("getting default executor: %w", err)
		}
		laneMandatedOutboundCCVs := make([]mcms.RawInstanceAddress, 0, len(sourceChain.LaneMandatedOutboundCCVs))
		for _, ccv := range sourceChain.LaneMandatedOutboundCCVs {
			outboundCCV, err := dsutils.GetRawInstanceAddressFromAddressRef(ccv)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("getting lane mandated outbound CCV: %w", err)
			}
			laneMandatedOutboundCCVs = append(laneMandatedOutboundCCVs, outboundCCV.Binding())
		}
		defaultOutboundCCVs := make([]mcms.RawInstanceAddress, 0, len(sourceChain.DefaultOutboundCCVs))
		for _, ccv := range sourceChain.DefaultOutboundCCVs {
			outboundCCV, err := dsutils.GetRawInstanceAddressFromAddressRef(ccv)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("getting default outbound CCV: %w", err)
			}
			defaultOutboundCCVs = append(defaultOutboundCCVs, outboundCCV.Binding())
		}

		var proposalOutputs []contract.ExerciseOutput

		gcIn, err := cantonMCMSChoiceInput(
			sourceChain.CantonLaneConfig.GlobalConfig,
			globalConfigAddress,
			common.ApplyDestChainConfigUpdates{
				DestChainConfigUpdates: []common.DestChainConfigArgs{
					{
						DestChainSelector:         types.NUMERIC(strconv.FormatUint(destChain.Selector, 10)),
						IsEnabled:                 types.BOOL(isEnabled),
						AddressBytesLength:        types.INT64(destChain.AddressBytesLength),
						TokenReceiverAllowed:      true, // TODO: missing from input
						BaseExecutionGasCost:      types.INT64(destChain.BaseExecutionGasCost),
						OffRampAddress:            types.TEXT(hex.EncodeToString(destChain.OffRamp)),
						DefaultExecutor:           new(defaultExecutor.Binding()),
						LaneMandatedCCVs:          laneMandatedOutboundCCVs,
						DefaultCCVs:               defaultOutboundCCVs,
						MessageNetworkFeeUSDCents: types.NUMERIC(strconv.FormatUint(uint64(destChain.MessageNetworkFeeUSDCents), 10)),
						TokenNetworkFeeUSDCents:   types.NUMERIC(strconv.FormatUint(uint64(destChain.TokenNetworkFeeUSDCents), 10)),
					},
				},
			},
			pd,
		)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		gcRep, err := operations.ExecuteOperation(b, global_config.ApplyDestChainConfigUpdates, chain, gcIn)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("applying dest chain config updates to global config: %w", err)
		}
		proposalOutputs = appendProposalExercise(pd, proposalOutputs, gcRep.Output)

		executorAddress, err := dsutils.ToInstanceAddress(sourceChain.DefaultExecutor)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("executor instance address: %w", err)
		}
		exIn, err := cantonMCMSChoiceInput(
			sourceChain.DefaultExecutor,
			executorAddress,
			executor.ApplyDestChainUpdates{
				DestChainSelectorsToRemove: nil,
				DestChainSelectorsToAdd: []executor.RemoteChainConfigArgs{
					{
						DestChainSelector: types.NUMERIC(strconv.FormatUint(destChain.Selector, 10)),
						Config: executor.RemoteChainConfig{
							FeeUSDCents: types.NUMERIC(strconv.FormatUint(uint64(destChain.ExecutorDestChainConfig.USDCentsFee), 10)),
							Enabled:     types.BOOL(destChain.ExecutorDestChainConfig.Enabled),
						},
					},
				},
			},
			pd,
		)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		exRep, err := operations.ExecuteOperation(b, executor2.ApplyDestChainUpdates, chain, exIn)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("applying dest chain config updates to executor: %w", err)
		}
		proposalOutputs = appendProposalExercise(pd, proposalOutputs, exRep.Output)

		feeQuoterAddress, feeQuoterRef, err := feeQuoterInstanceAndRef(sourceChain)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("fee quoter target: %w", err)
		}
		if hackRef, ok := hardcodedFeeQuoterMCMSRef(sourceChain.Selector); ok {
			feeQuoterRef = hackRef
			hackAddr, err := dsutils.ToInstanceAddress(hackRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("hardcoded fee quoter instance address: %w", err)
			}
			feeQuoterAddress = hackAddr
		}
		feeQuoterMCMS := feeQuoterRef.Address != "" || len(feeQuoterRef.Labels.List()) > 0
		if !feeQuoterMCMS {
			b.Logger.Warnf("skipping FeeQuoter MCMS exercises (no labeled AddressRef on chain %d); global config and executor proposal txs are still emitted", sourceChain.Selector)
		}

		if feeQuoterMCMS {
			fqDestIn, err := cantonMCMSChoiceInput(
				feeQuoterRef,
				feeQuoterAddress,
				feequoter.ApplyDestChainConfigUpdates2{
					DestChainConfigArgs: []feequoter.DestChainConfigArgs2{
						{
							DestChainSelector: types.NUMERIC(strconv.FormatUint(destChain.Selector, 10)),
							DestChainConfig: feequoter.DestChainConfig2{
								IsEnabled:                   types.BOOL(destChain.FeeQuoterDestChainConfig.IsEnabled),
								MaxDataBytes:                types.INT64(destChain.FeeQuoterDestChainConfig.MaxDataBytes),
								MaxPerMsgGasLimit:           types.INT64(destChain.FeeQuoterDestChainConfig.MaxPerMsgGasLimit),
								DestGasOverhead:             types.INT64(destChain.FeeQuoterDestChainConfig.DestGasOverhead),
								DestGasPerPayloadByteBase:   types.INT64(destChain.FeeQuoterDestChainConfig.DestGasPerPayloadByteBase),
								DefaultTxGasLimit:           types.INT64(destChain.FeeQuoterDestChainConfig.DefaultTxGasLimit),
								LinkFeeMultiplierPercent:    types.NUMERIC(strconv.FormatUint(uint64(destChain.FeeQuoterDestChainConfig.V2Params.LinkFeeMultiplierPercent), 10)),
								DefaultTokenFeeUSD:          types.NUMERIC(strconv.FormatUint(uint64(destChain.FeeQuoterDestChainConfig.DefaultTokenFeeUSDCents), 10)),
								DefaultTokenDestGasOverhead: types.INT64(destChain.FeeQuoterDestChainConfig.DefaultTokenDestGasOverhead),
							},
						},
					},
				},
				pd,
			)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			fqDestRep, err := operations.ExecuteOperation(b, feequoterop.ApplyDestChainConfigUpdates, chain, fqDestIn)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("applying dest chain config updates to fee quoter: %w", err)
			}
			proposalOutputs = appendProposalExercise(pd, proposalOutputs, fqDestRep.Output)

			fqPriceIn, err := cantonMCMSChoiceInput(
				feeQuoterRef,
				feeQuoterAddress,
				feequoter.ApplyPriceUpdatersUpdate{
					AddedPriceUpdaters:   []types.PARTY{types.PARTY(chain.Participants[0].PartyID)},
					RemovedPriceUpdaters: nil,
				},
				pd,
			)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			fqPriceRep, err := operations.ExecuteOperation(b, feequoterop.ApplyPriceUpdatersUpdate, chain, fqPriceIn)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("ensuring price updater on fee quoter: %w", err)
			}
			proposalOutputs = appendProposalExercise(pd, proposalOutputs, fqPriceRep.Output)

			tokenPriceUpdates, err := tokenPriceUpdatesFromParams(destChain.TokenPrices)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("building token price updates from lane params: %w", err)
			}

			fqUpdIn, err := cantonMCMSChoiceInput(
				feeQuoterRef,
				feeQuoterAddress,
				feequoter.UpdatePrices{
					PriceUpdates: feequoter.PriceUpdates{
						TokenPriceUpdates: tokenPriceUpdates,
						GasPriceUpdates: []feequoter.GasPriceUpdate{
							{
								DestChainSelector: types.NUMERIC(strconv.FormatUint(destChain.Selector, 10)),
								UsdPerUnitGas: cantonFeeQuoterUSDPerUnitGas(
									destChain.FeeQuoterDestChainConfig.V2Params.USDPerUnitGas),
							},
						},
					},
				},
				pd,
			)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			fqUpdRep, err := operations.ExecuteOperation(b, feequoterop.UpdatePrices, chain, fqUpdIn)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("updating prices in fee quoter: %w", err)
			}
			proposalOutputs = appendProposalExercise(pd, proposalOutputs, fqUpdRep.Output)
		}

		var committeeExtra []mcms_types.BatchOperation
		for _, verifierConfig := range input.Source.CommitteeVerifiers {
			subReport, err := operations.ExecuteSequence(b, ConfigureCommitteeVerifierAsSource, deps, ConfigureCommitteeVerifierAsSourceInput{
				ChainSelector:           chain.Selector,
				CommitteeVerifierConfig: verifierConfig,
				ProposalDriven:          pd,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("configuring committee verifier as source: %w", err)
			}
			committeeExtra = append(committeeExtra, subReport.Output.BatchOps...)
		}

		return mergeCantonLaneProposalBatches(chain.Selector, proposalOutputs, committeeExtra)
	},
)

var ConfigureLaneLegAsDest = operations.NewSequence(
	"CantonConfigureLaneLegAsDest",
	semver.MustParse("2.0.0"),
	"Configures a lane leg as dest on CCIP 2.0.0",
	func(b operations.Bundle, deps chain.BlockChains, input lanes.UpdateLanesInput) (output sequences.OnChainOutput, err error) {
		b.Logger.Infof("Canton Configuring lane leg as dest. src: %+v, dest: %+v", input.Source, input.Dest)

		chain, ok := deps.CantonChains()[input.Dest.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.Dest.Selector)
		}

		sourceChain := input.Source
		destChain := input.Dest
		// HACK: always collect MCMS proposal transactions for this sequence (ignores input.ProposalDriven).
		pd := true

		if destChain.CantonLaneConfig == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("CantonLaneConfig is required for proposal-driven Canton lane configuration (dest chain %d)", destChain.Selector)
		}

		globalConfigAddress, err := dsutils.ToInstanceAddress(destChain.CantonLaneConfig.GlobalConfig)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("global config instance address: %w", err)
		}

		laneMandatedInboundCCVs := make([]mcms.RawInstanceAddress, 0, len(destChain.LaneMandatedInboundCCVs))
		for _, ccv := range destChain.LaneMandatedInboundCCVs {
			inboundCCV, err := dsutils.GetRawInstanceAddressFromAddressRef(ccv)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("getting lane mandated inbound CCV: %w", err)
			}
			laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, inboundCCV.Binding())
		}
		defaultInboundCCVs := make([]mcms.RawInstanceAddress, 0, len(destChain.DefaultInboundCCVs))
		for _, ccv := range destChain.DefaultInboundCCVs {
			inboundCCV, err := dsutils.GetRawInstanceAddressFromAddressRef(ccv)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("getting default inbound CCV: %w", err)
			}
			defaultInboundCCVs = append(defaultInboundCCVs, inboundCCV.Binding())
		}

		inboundOnRampAddresses := []types.TEXT{
			types.TEXT(hex.EncodeToString(gethcommon.LeftPadBytes(sourceChain.OnRamp, 32))),
		}

		var proposalOutputs []contract.ExerciseOutput

		gcIn, err := cantonMCMSChoiceInput(
			destChain.CantonLaneConfig.GlobalConfig,
			globalConfigAddress,
			common.ApplySourceChainConfigUpdates{
				SourceChainConfigUpdates: []common.SourceChainConfigArgs{
					{
						SourceChainSelector: types.NUMERIC(strconv.FormatUint(sourceChain.Selector, 10)),
						IsEnabled:           types.BOOL(!input.IsDisabled),
						OnRampAddresses:     inboundOnRampAddresses,
						DefaultCCVs:         defaultInboundCCVs,
						LaneMandatedCCVs:    laneMandatedInboundCCVs,
					},
				},
			},
			pd,
		)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		gcRep, err := operations.ExecuteOperation(b, global_config.ApplySourceChainConfigUpdates, chain, gcIn)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("applying source chain config updates to global config: %w", err)
		}
		proposalOutputs = appendProposalExercise(pd, proposalOutputs, gcRep.Output)

		var committeeExtra []mcms_types.BatchOperation
		for _, verifierConfig := range input.Dest.CommitteeVerifiers {
			subReport, err := operations.ExecuteSequence(b, ConfigureCommitteeVerifierAsDest, deps, ConfigureCommitteeVerifierAsDestInput{
				ChainSelector:           chain.Selector,
				CommitteeVerifierConfig: verifierConfig,
				ProposalDriven:          pd,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("configuring committee verifier as dest: %w", err)
			}
			committeeExtra = append(committeeExtra, subReport.Output.BatchOps...)
		}

		return mergeCantonLaneProposalBatches(chain.Selector, proposalOutputs, committeeExtra)
	},
)

func tokenPriceUpdatesFromParams(tokenPrices map[string]*big.Int) ([]feequoter.TokenPriceUpdate, error) {
	if len(tokenPrices) == 0 {
		return nil, nil
	}
	updates := make([]feequoter.TokenPriceUpdate, 0, len(tokenPrices))
	for instrument, price := range tokenPrices {
		if price == nil {
			continue
		}
		parts := strings.SplitN(instrument, ":", 2)
		if len(parts) != 2 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
			return nil, fmt.Errorf("invalid token price instrument key %q, expected format <admin>:<id>", instrument)
		}
		updates = append(updates, feequoter.TokenPriceUpdate{
			InstrumentId: splice_api_token_holding_v1.InstrumentId{
				Admin: types.PARTY(strings.TrimSpace(parts[0])),
				Id:    types.TEXT(strings.TrimSpace(parts[1])),
			},
			UsdPerToken: types.NUMERIC(price.String()),
		})
	}

	return updates, nil
}
