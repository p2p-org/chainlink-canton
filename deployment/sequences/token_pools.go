package sequences

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	gethcommon "github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	tokenadaptersfinality "github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	tokenadapters "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	ccipsequences "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfcanton "github.com/smartcontractkit/chainlink-deployments-framework/chain/canton"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/go-daml/pkg/types"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-canton/bindings"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/burnminttokenpool"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/common"
	factorybindings "github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/factory"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/lockreleasetokenpool"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/mcms"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/splice/splice_api_token_holding_v1"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/splice/splice_api_token_metadata_v1"
	"github.com/smartcontractkit/chainlink-canton/contracts"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/committee_verifier"
	factoryops "github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/factory"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/fee_quoter"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/rate_limiter"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/rmn_remote"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/token_admin_registry"
	dsutils "github.com/smartcontractkit/chainlink-canton/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-canton/deployment/utils/operations/contract"
)

var (
	lockReleasePoolType = datastore.ContractType("LockReleaseTokenPool")
	burnMintPoolType    = datastore.ContractType("BurnMintTokenPool")
)

type configuredCantonTokenPool struct {
	InstrumentId       splice_api_token_holding_v1.InstrumentId
	InstanceId         types.TEXT
	CcipOwner          types.PARTY
	PoolOwner          types.PARTY
	Decimals           types.INT64
	RemoteChainConfigs map[types.NUMERIC]any
}

type tokenPoolChainUpdate struct {
	RemoteChainSelector                        types.NUMERIC
	RemotePools                                []types.TEXT
	RemoteTokenAddress                         types.TEXT
	InboundCCVs                                []mcms.RawInstanceAddress
	OutboundCCVs                               []mcms.RawInstanceAddress
	FinalityConfig                             common.FinalityConfig
	InboundRateLimiter                         mcms.RawInstanceAddress
	InboundCustomBlockConfirmationsRateLimiter mcms.RawInstanceAddress
	OutboundRateLimiter                        mcms.RawInstanceAddress
}

type rateLimiterPoolMeta struct {
	InstanceId types.TEXT
	PoolOwner  types.PARTY
}

// appendTokenPoolDatastoreRefsForConfigure adds logical pool, canonical pool, and token refs so
// datastore merges match what DeployTokenPoolForToken used to return; deploy-only flow no longer registers or emits these.
func appendTokenPoolDatastoreRefsForConfigure(
	out *ccipsequences.OnChainOutput,
	input tokenadapters.ConfigureTokenForTransfersInput,
	logicalPoolType datastore.ContractType,
	parsedPool *configuredCantonTokenPool,
) {
	if parsedPool == nil {
		return
	}
	poolAddr := strings.TrimSpace(input.TokenPoolAddress)
	if poolAddr == "" {
		return
	}

	var canonPoolType datastore.ContractType
	var canonPoolVersion *semver.Version
	switch logicalPoolType {
	case lockReleasePoolType:
		canonPoolType = datastore.ContractType(lock_release_token_pool.ContractType)
		canonPoolVersion = lock_release_token_pool.Version
	case burnMintPoolType:
		canonPoolType = datastore.ContractType(burn_mint_token_pool.ContractType)
		canonPoolVersion = burn_mint_token_pool.Version
	default:
		return
	}

	rawPool := contracts.InstanceID(strings.TrimSpace(string(parsedPool.InstanceId))).RawInstanceAddress(parsedPool.PoolOwner)

	poolVersion := input.TokenRef.Version
	if poolVersion == nil {
		poolVersion = semver.MustParse("2.0.0")
	}

	logicalRef := datastore.AddressRef{
		Address:       poolAddr,
		Labels:        datastore.NewLabelSet(rawPool.String()),
		Type:          logicalPoolType,
		Version:       poolVersion,
		Qualifier:     strings.TrimSpace(input.TokenRef.Qualifier),
		ChainSelector: input.ChainSelector,
	}
	canonicalRef := datastore.AddressRef{
		Address:       poolAddr,
		Labels:        logicalRef.Labels,
		Type:          canonPoolType,
		Version:       canonPoolVersion,
		Qualifier:     "",
		ChainSelector: input.ChainSelector,
	}
	out.Addresses = append(out.Addresses, logicalRef, canonicalRef)

	if strings.TrimSpace(input.TokenRef.Address) != "" {
		out.Addresses = append(out.Addresses, input.TokenRef)
	}
}

var ConfigureTokenForTransfers = operations.NewSequence(
	"canton/token-adapter/configure-token-for-transfers",
	semver.MustParse("2.0.0"),
	"Registers the pool with token admin registry, applies remote-chain updates, deploys rate limiters, and merges logical/canonical pool + token datastore refs",
	func(b operations.Bundle, chains chain.BlockChains, input tokenadapters.ConfigureTokenForTransfersInput) (ccipsequences.OnChainOutput, error) {
		if input.ExistingDataStore == nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("existing datastore is required")
		}

		cantonChain, ok := chains.CantonChains()[input.ChainSelector]
		if !ok || len(cantonChain.Participants) == 0 {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("canton chain with selector %d not found", input.ChainSelector)
		}
		participant := cantonChain.Participants[0]
		poolAddress := contracts.HexToInstanceAddress(input.TokenPoolAddress)
		logicalPoolType, err := resolveCantonTokenPoolType(input.PoolType)
		if err != nil {
			return ccipsequences.OnChainOutput{}, err
		}

		parsedPool, err := loadConfiguredCantonTokenPool(b.GetContext(), participant, logicalPoolType, poolAddress)
		if err != nil {
			return ccipsequences.OnChainOutput{}, err
		}

		registryRef, err := input.ExistingDataStore.Addresses().Get(datastore.NewAddressRefKey(
			input.ChainSelector,
			datastore.ContractType(token_admin_registry.ContractType),
			token_admin_registry.Version,
			"",
		))
		if err != nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("resolve token admin registry: %w", err)
		}
		_, err = operations.ExecuteSequence(b, RegisterTokenPool, cantonChain, RegisterTokenPoolInput{
			TokenAdminRegistryInstanceAddress: contracts.HexToInstanceAddress(registryRef.Address),
			InstrumentId:                      parsedPool.InstrumentId,
			PoolInstanceID:                    string(parsedPool.InstanceId),
			CcipParty:                         string(parsedPool.CcipOwner),
			PoolOwnerParty:                    string(parsedPool.PoolOwner),
		})
		if err != nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("register token pool with token admin registry: %w", err)
		}

		out := ccipsequences.OnChainOutput{}
		appendTokenPoolDatastoreRefsForConfigure(&out, input, logicalPoolType, parsedPool)

		committeeVerifierRefs := input.ExistingDataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
			datastore.AddressRefByType(datastore.ContractType(committee_verifier.ContractType)),
			datastore.AddressRefByVersion(committee_verifier.Version),
		)

		updates := make([]tokenPoolChainUpdate, 0, len(input.RemoteChains))
		for remoteSelector, remoteCfg := range input.RemoteChains {
			remoteSelectorKeyStr := strconv.FormatUint(remoteSelector, 10)
			remoteSelectorKey := types.NUMERIC(remoteSelectorKeyStr)
			if _, found := parsedPool.RemoteChainConfigs[remoteSelectorKey]; found {
				return out, fmt.Errorf("remote chain %d is already configured on token pool", remoteSelector)
			}

			inboundCCVs, err := resolveCommitteeVerifierRawAddresses(committeeVerifierRefs, remoteCfg.InboundCCVs)
			if err != nil {
				return out, fmt.Errorf("resolve inbound CCVs for remote chain %d: %w", remoteSelector, err)
			}
			outboundCCVs, err := resolveCommitteeVerifierRawAddresses(committeeVerifierRefs, remoteCfg.OutboundCCVs)
			if err != nil {
				return out, fmt.Errorf("resolve outbound CCVs for remote chain %d: %w", remoteSelector, err)
			}

			outboundDefaultCfg, inboundDefaultCfg := tokenadapters.GenerateTPRLConfigs(
				remoteCfg.OutboundRateLimiterConfig,
				remoteCfg.InboundRateLimiterConfig,
				uint8(parsedPool.Decimals),
				remoteCfg.RemoteDecimals,
				"canton",
				semver.MustParse("2.0.0"),
				lockReleasePoolType.String(),
			)
			_, inboundCustomCfg := tokenadapters.GenerateTPRLConfigs(
				remoteCfg.OutboundRateLimiterConfig, // TODO: how do we get the "custom finality" config?
				remoteCfg.InboundRateLimiterConfig,  // TODO: how do we get the "custom finality" config?
				uint8(parsedPool.Decimals),
				remoteCfg.RemoteDecimals,
				"canton",
				semver.MustParse("2.0.0"),
				lockReleasePoolType.String(),
			)

			meta := rateLimiterPoolMeta{InstanceId: parsedPool.InstanceId, PoolOwner: parsedPool.PoolOwner}
			outboundRef, outboundRaw, err := deployTokenPoolRateLimiter(
				b,
				cantonChain,
				input.ExistingDataStore,
				meta,
				input.TokenPoolAddress,
				remoteSelectorKeyStr,
				"outbound",
				outboundDefaultCfg,
				common.RateLimitModeRateLimitMode_DefaultFinality,
			)
			if err != nil {
				return out, fmt.Errorf("deploy outbound rate limiter for remote chain %d: %w", remoteSelector, err)
			}
			out.Addresses = append(out.Addresses, outboundRef)

			inboundRef, inboundRaw, err := deployTokenPoolRateLimiter(
				b,
				cantonChain,
				input.ExistingDataStore,
				meta,
				input.TokenPoolAddress,
				remoteSelectorKeyStr,
				"inbound",
				inboundDefaultCfg,
				common.RateLimitModeRateLimitMode_DefaultFinality,
			)
			if err != nil {
				return out, fmt.Errorf("deploy inbound rate limiter for remote chain %d: %w", remoteSelector, err)
			}
			out.Addresses = append(out.Addresses, inboundRef)

			customRef, customInboundRaw, err := deployTokenPoolRateLimiter(
				b,
				cantonChain,
				input.ExistingDataStore,
				meta,
				input.TokenPoolAddress,
				remoteSelectorKeyStr,
				"inbound-custom",
				inboundCustomCfg,
				common.RateLimitModeRateLimitMode_CustomFinality,
			)
			if err != nil {
				return out, fmt.Errorf("deploy custom inbound rate limiter for remote chain %d: %w", remoteSelector, err)
			}
			out.Addresses = append(out.Addresses, customRef)

			remoteFamily, err := chain_selectors.GetSelectorFamily(remoteSelector)
			if err != nil {
				return out, fmt.Errorf("get remote chain family for %d: %w", remoteSelector, err)
			}
			remotePoolAddress := strings.ToLower(hex.EncodeToString(remoteCfg.RemotePool))
			remoteTokenAddress := strings.ToLower(hex.EncodeToString(remoteCfg.RemoteToken))
			if remoteFamily == chain_selectors.FamilyEVM {
				remotePoolAddress = strings.TrimPrefix(strings.ToLower(gethcommon.BytesToHash(remoteCfg.RemotePool).Hex()), "0x")
				remoteTokenAddress = strings.TrimPrefix(strings.ToLower(gethcommon.BytesToAddress(remoteCfg.RemoteToken).Hex()), "0x")
			}

			updates = append(updates, tokenPoolChainUpdate{
				RemoteChainSelector: remoteSelectorKey,
				RemotePools:         []types.TEXT{types.TEXT(remotePoolAddress)},
				RemoteTokenAddress:  types.TEXT(remoteTokenAddress),
				InboundCCVs:         inboundCCVs,
				OutboundCCVs:        outboundCCVs,
				FinalityConfig:      toCantonFinalityConfig(input.AllowedFinalityConfig),
				InboundRateLimiter:  inboundRaw,
				InboundCustomBlockConfirmationsRateLimiter: customInboundRaw,
				OutboundRateLimiter:                        outboundRaw,
			})
		}

		if len(updates) == 0 {
			return out, nil
		}

		switch logicalPoolType {
		case lockReleasePoolType:
			lockReleaseUpdates := make([]lockreleasetokenpool.ChainUpdate, 0, len(updates))
			for _, update := range updates {
				lockReleaseUpdates = append(lockReleaseUpdates, lockreleasetokenpool.ChainUpdate{
					RemoteChainSelector: update.RemoteChainSelector,
					RemotePools:         update.RemotePools,
					RemoteTokenAddress:  update.RemoteTokenAddress,
					InboundCCVs:         update.InboundCCVs,
					OutboundCCVs:        update.OutboundCCVs,
					FinalityConfig:      update.FinalityConfig,
					InboundRateLimiter:  update.InboundRateLimiter,
					InboundCustomBlockConfirmationsRateLimiter: update.InboundCustomBlockConfirmationsRateLimiter,
					OutboundRateLimiter:                        update.OutboundRateLimiter,
				})
			}
			_, err = operations.ExecuteOperation(b, lock_release_token_pool.ApplyChainUpdates, cantonChain, contract.ChoiceInput[lockreleasetokenpool.ApplyChainUpdates]{
				InstanceAddress: poolAddress,
				Args: lockreleasetokenpool.ApplyChainUpdates{
					RemoteChainSelectorsToRemove: []types.NUMERIC{},
					ChainsToAdd:                  lockReleaseUpdates,
				},
			})
			if err != nil {
				if strings.Contains(err.Error(), "ApplyChainUpdates: chain already exists:") {
					return out, nil
				}

				return out, fmt.Errorf("apply remote chain updates to lock/release pool: %w", err)
			}
		case burnMintPoolType:
			burnMintUpdates := make([]burnminttokenpool.ChainUpdate, 0, len(updates))
			for _, update := range updates {
				burnMintUpdates = append(burnMintUpdates, burnminttokenpool.ChainUpdate{
					RemoteChainSelector: update.RemoteChainSelector,
					RemotePools:         update.RemotePools,
					RemoteTokenAddress:  update.RemoteTokenAddress,
					InboundCCVs:         update.InboundCCVs,
					OutboundCCVs:        update.OutboundCCVs,
					FinalityConfig:      update.FinalityConfig,
					InboundRateLimiter:  update.InboundRateLimiter,
					InboundCustomBlockConfirmationsRateLimiter: update.InboundCustomBlockConfirmationsRateLimiter,
					OutboundRateLimiter:                        update.OutboundRateLimiter,
				})
			}
			_, err = operations.ExecuteOperation(b, burn_mint_token_pool.ApplyChainUpdates, cantonChain, contract.ChoiceInput[burnminttokenpool.ApplyChainUpdates]{
				InstanceAddress: poolAddress,
				Args: burnminttokenpool.ApplyChainUpdates{
					RemoteChainSelectorsToRemove: []types.NUMERIC{},
					ChainsToAdd:                  burnMintUpdates,
				},
			})
			if err != nil {
				if strings.Contains(err.Error(), "ApplyChainUpdates: chain already exists:") {
					return out, nil
				}

				return out, fmt.Errorf("apply remote chain updates to burn/mint pool: %w", err)
			}
		default:
			return out, fmt.Errorf("unsupported Canton token pool type %q", logicalPoolType)
		}

		return out, nil
	},
)

// Canton wires rate limiters during initial remote-chain setup in ConfigureTokenForTransfers
// via ApplyChainUpdates, so this follow-up sequence is intentionally unused.
var SetTokenPoolRateLimits = operations.NewSequence(
	"canton/token-adapter/set-token-pool-rate-limits",
	semver.MustParse("2.0.0"),
	"Updates Canton lock/release pool rate limiter config for a remote chain",
	func(b operations.Bundle, chains chain.BlockChains, input tokenadapters.TPRLRemotes) (ccipsequences.OnChainOutput, error) {
		_ = b
		_ = chains
		_ = input

		return ccipsequences.OnChainOutput{}, fmt.Errorf(
			"SetTokenPoolRateLimits is disabled: initial Canton rate limiter setup happens during ConfigureTokenForTransfers",
		)
	},
)

var DeployTokenPoolForToken = operations.NewSequence(
	"canton/token-adapter/deploy-token-pool-for-token",
	semver.MustParse("2.0.0"),
	"Deploys a Canton token pool contract only (no TAR registration or logical/token datastore refs); run ConfigureTokenForTransfers for that",
	func(b operations.Bundle, chains chain.BlockChains, input tokenadapters.DeployTokenPoolInput) (ccipsequences.OnChainOutput, error) {
		if input.TokenPoolVersion == nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("TokenPoolVersion is required")
		}
		if input.ExistingDataStore == nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("existing datastore is required")
		}

		logicalPoolType, err := resolveCantonTokenPoolType(input.PoolType)
		if err != nil {
			return ccipsequences.OnChainOutput{}, err
		}

		qualifier := strings.TrimSpace(input.TokenPoolQualifier)
		if qualifier == "" {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("token pool qualifier is required")
		}
		if input.TokenRef == nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("tokenRef is required and must include instrument labels")
		}

		instrumentID, instrumentAdmin, err := parseInstrumentIDFromTokenRefLabels(*input.TokenRef)
		if err != nil {
			return ccipsequences.OnChainOutput{}, err
		}

		matches := input.ExistingDataStore.Addresses().Filter(
			datastore.AddressRefByType(logicalPoolType),
			datastore.AddressRefByChainSelector(input.ChainSelector),
			datastore.AddressRefByQualifier(qualifier),
			datastore.AddressRefByVersion(input.TokenPoolVersion),
		)
		if len(matches) > 1 {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("multiple Canton token pools found with qualifier %q", qualifier)
		}
		if len(matches) == 1 {
			b.Logger.Infof("Canton token pool already deployed at %s", matches[0].Address)
			return ccipsequences.OnChainOutput{}, nil
		}

		cantonChain, ok := chains.CantonChains()[input.ChainSelector]
		if !ok || len(cantonChain.Participants) == 0 {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("canton chain with selector %d not found", input.ChainSelector)
		}
		participant := cantonChain.Participants[0]
		if logicalPoolType == burnMintPoolType && instrumentAdmin != participant.PartyID {
			return ccipsequences.OnChainOutput{}, fmt.Errorf(
				"burn/mint pools require instrument-admin %q to match pool owner %q",
				instrumentAdmin,
				participant.PartyID,
			)
		}

		resolveRefAndRaw := func(name string, contractType datastore.ContractType, version *semver.Version) (datastore.AddressRef, contracts.RawInstanceAddress, error) {
			ref, err := input.ExistingDataStore.Addresses().Get(datastore.NewAddressRefKey(
				input.ChainSelector,
				contractType,
				version,
				"",
			))
			if err != nil {
				return datastore.AddressRef{}, "", fmt.Errorf("resolve %s: %w", name, err)
			}
			raw, err := dsutils.GetRawInstanceAddressFromAddressRef(ref)
			if err != nil {
				return datastore.AddressRef{}, "", fmt.Errorf("resolve %s raw address: %w", name, err)
			}

			return ref, raw, nil
		}

		_, tokenAdminRegistryRaw, err := resolveRefAndRaw("token admin registry", datastore.ContractType(token_admin_registry.ContractType), token_admin_registry.Version)
		if err != nil {
			return ccipsequences.OnChainOutput{}, err
		}
		_, rmnRemoteRaw, err := resolveRefAndRaw("rmn remote", datastore.ContractType(rmn_remote.ContractType), rmn_remote.Version)
		if err != nil {
			return ccipsequences.OnChainOutput{}, err
		}
		_, feeQuoterRaw, err := resolveRefAndRaw("fee quoter", datastore.ContractType(fee_quoter.ContractType), fee_quoter.Version)
		if err != nil {
			return ccipsequences.OnChainOutput{}, err
		}

		var deployOutput datastore.AddressRef
		switch logicalPoolType {
		case lockReleasePoolType:
			factorySeqReport, err := operations.ExecuteSequence(b, DeployLockReleaseTokenPoolFromFactory, cantonChain, DeployLockReleaseTokenPoolFromFactoryInput{
				ExistingDataStore: input.ExistingDataStore,
				ChainSelector:     input.ChainSelector,
				ProposalDriven:    false,
				CcipOwner:         participant.PartyID,
				PoolOwner:         participant.PartyID,
				InstrumentId:      instrumentID,
				Decimals:          10,
				Deps: lockreleasetokenpool.LockReleaseTokenPoolDeps{
					TokenAdminRegistry: tokenAdminRegistryRaw.Binding(),
					RmnRemote:          rmnRemoteRaw.Binding(),
					FeeQuoter:          feeQuoterRaw.Binding(),
				},
				RemoteChainSelectors: nil,
			})
			if err != nil {
				return ccipsequences.OnChainOutput{}, fmt.Errorf("deploy Canton lock/release pool via CCIPFactory: %w", err)
			}
			for _, ref := range factorySeqReport.Output.Addresses {
				if ref.Type == datastore.ContractType(lock_release_token_pool.ContractType) {
					deployOutput = ref
					break
				}
			}
			if deployOutput.Address == "" {
				return ccipsequences.OnChainOutput{}, fmt.Errorf("lock/release pool address ref missing from factory sequence output")
			}
		case burnMintPoolType:
			deployReport, err := operations.ExecuteOperation(b, burn_mint_token_pool.Deploy, cantonChain, contract.DeployInput[burnminttokenpool.BurnMintTokenPool]{
				Qualifier: new(qualifier),
				Template: burnminttokenpool.BurnMintTokenPool{
					CcipOwner:               types.PARTY(participant.PartyID),
					PoolOwner:               types.PARTY(participant.PartyID),
					InstrumentId:            instrumentID,
					Decimals:                types.INT64(10),
					RemoteChainConfigs:      map[types.NUMERIC]burnminttokenpool.RemoteChainConfig{},
					TokenTransferFeeConfigs: map[types.NUMERIC]burnminttokenpool.TokenTransferFeeConfig{},
					PoolReceiveContext: splice_api_token_metadata_v1.ChoiceContext{
						Values: map[string]splice_api_token_metadata_v1.AnyValue{},
					},
					TransferTimeout: burnminttokenpool.TransferTimeout{
						RelativeHours: new(types.INT64(24)),
					},
					Deps: burnminttokenpool.BurnMintTokenPoolDeps{
						TokenAdminRegistry: tokenAdminRegistryRaw.Binding(),
						RmnRemote:          rmnRemoteRaw.Binding(),
						FeeQuoter:          feeQuoterRaw.Binding(),
					},
				},
				OwnerParty: types.PARTY(participant.PartyID),
			})
			if err != nil {
				return ccipsequences.OnChainOutput{}, fmt.Errorf("deploy Canton burn/mint pool: %w", err)
			}
			deployOutput = deployReport.Output
		default:
			return ccipsequences.OnChainOutput{}, fmt.Errorf("unsupported Canton token pool type %q", logicalPoolType)
		}

		if len(deployOutput.Labels.List()) == 0 {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("missing raw token pool label in deploy output")
		}

		return ccipsequences.OnChainOutput{
			Addresses: []datastore.AddressRef{deployOutput},
		}, nil
	},
)

func toCantonFinalityConfig(cfg tokenadaptersfinality.Config) common.FinalityConfig {
	switch {
	case cfg.WaitForFinality:
		return common.FinalityConfig{WaitForFinality: &types.UNIT{}}
	case cfg.WaitForSafe:
		return common.FinalityConfig{WaitForSafe: &types.UNIT{}}
	case cfg.BlockDepth > 0:
		return common.FinalityConfig{BlockDepth: new(types.INT64(cfg.BlockDepth))}
	default:
		return common.FinalityConfig{}
	}
}

func deployTokenPoolRateLimiter(
	b operations.Bundle,
	cantonChain cldfcanton.Chain,
	existingDataStore datastore.DataStore,
	poolMeta rateLimiterPoolMeta,
	tokenPoolAddress string,
	remoteSelectorKey string,
	direction string,
	cfg tokenadapters.RateLimiterConfig,
	mode common.RateLimitMode,
) (datastore.AddressRef, mcms.RawInstanceAddress, error) {
	deployOp := rate_limiter.DeployInbound
	dirEnum := common.RateLimitDirectionRateLimitDirection_Inbound
	if direction == "outbound" {
		deployOp = rate_limiter.DeployOutbound
		dirEnum = common.RateLimitDirectionRateLimitDirection_Outbound
	}

	capacity := types.NUMERIC("0")
	if cfg.Capacity != nil {
		capacity = types.NUMERIC(cfg.Capacity.String())
	}
	rate := types.NUMERIC("0")
	if cfg.Rate != nil {
		rate = types.NUMERIC(cfg.Rate.String())
	}

	report, err := operations.ExecuteOperation(b, deployOp, cantonChain, contract.DeployInput[common.RateLimiter]{
		Qualifier: new(fmt.Sprintf("%s-%s-%s", tokenPoolAddress, direction, remoteSelectorKey)),
		Template: common.RateLimiter{
			PoolInstanceId:      poolMeta.InstanceId,
			PoolOwner:           poolMeta.PoolOwner,
			RemoteChainSelector: types.NUMERIC(remoteSelectorKey),
			Direction:           dirEnum,
			Mode:                mode,
			IsEnabled:           types.BOOL(cfg.IsEnabled),
			Capacity:            capacity,
			Rate:                rate,
			Tokens:              capacity,
			LastUpdated:         types.TIMESTAMP(time.Now()),
		},
		OwnerParty: poolMeta.PoolOwner,
	})
	if err != nil {
		return datastore.AddressRef{}, mcms.RawInstanceAddress{}, err
	}
	addresses, ok := existingDataStore.Addresses().(datastore.MutableAddressRefStore)
	if !ok {
		return datastore.AddressRef{}, mcms.RawInstanceAddress{}, fmt.Errorf("existing datastore addresses are not mutable")
	}
	if err := addresses.Add(report.Output); err != nil {
		return datastore.AddressRef{}, mcms.RawInstanceAddress{}, err
	}

	rawAddr, err := dsutils.GetRawInstanceAddressFromAddressRef(report.Output)
	if err != nil {
		return datastore.AddressRef{}, mcms.RawInstanceAddress{}, err
	}

	return report.Output, rawAddr.Binding(), nil
}

func resolveCommitteeVerifierRawAddresses(refs []datastore.AddressRef, addresses []string) ([]mcms.RawInstanceAddress, error) {
	result := make([]mcms.RawInstanceAddress, 0, len(addresses))
	for _, address := range addresses {
		var matchedRef *datastore.AddressRef
		for _, ref := range refs {
			if strings.EqualFold(ref.Address, address) {
				matchedRef = new(ref)
				break
			}
		}
		if matchedRef == nil {
			return nil, fmt.Errorf("resolve committee verifier ref for address %s", address)
		}
		rawAddress, err := dsutils.GetRawInstanceAddressFromAddressRef(*matchedRef)
		if err != nil {
			return nil, fmt.Errorf("resolve committee verifier raw address for %s: %w", address, err)
		}
		result = append(result, rawAddress.Binding())
	}

	return result, nil
}

func resolveCantonTokenPoolType(poolType string) (datastore.ContractType, error) {
	trimmed := strings.TrimSpace(poolType)
	switch datastore.ContractType(trimmed) {
	case "":
		return lockReleasePoolType, nil
	case lockReleasePoolType:
		return lockReleasePoolType, nil
	case burnMintPoolType:
		return burnMintPoolType, nil
	default:
		return "", fmt.Errorf("unsupported Canton token pool type %q", poolType)
	}
}

func parseInstrumentIDFromTokenRefLabels(tokenRef datastore.AddressRef) (splice_api_token_holding_v1.InstrumentId, string, error) {
	var instrumentAdmin, instrumentIDText string
	for _, label := range tokenRef.Labels.List() {
		switch {
		case strings.HasPrefix(label, "instrument-admin:"):
			instrumentAdmin = strings.TrimSpace(strings.TrimPrefix(label, "instrument-admin:"))
		case strings.HasPrefix(label, "instrument-id:"):
			instrumentIDText = strings.TrimSpace(strings.TrimPrefix(label, "instrument-id:"))
		}
	}
	if instrumentAdmin == "" || instrumentIDText == "" {
		return splice_api_token_holding_v1.InstrumentId{}, "", fmt.Errorf(
			"tokenRef labels must include instrument-admin:<party> and instrument-id:<id>",
		)
	}

	return splice_api_token_holding_v1.InstrumentId{
		Admin: types.PARTY(instrumentAdmin),
		Id:    types.TEXT(instrumentIDText),
	}, instrumentAdmin, nil
}

func loadConfiguredCantonTokenPool(
	ctx context.Context,
	participant cldfcanton.Participant,
	logicalPoolType datastore.ContractType,
	poolAddress contracts.InstanceAddress,
) (*configuredCantonTokenPool, error) {
	switch logicalPoolType {
	case lockReleasePoolType:
		activePool, err := contract.FindActiveContractByInstanceAddress(
			ctx,
			participant.LedgerServices.State,
			participant.PartyID,
			lockreleasetokenpool.LockReleaseTokenPool{}.GetTemplateID(),
			poolAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("find active lock/release pool %s: %w", poolAddress.Hex(), err)
		}
		parsedPool, err := bindings.UnmarshalCreatedEvent[lockreleasetokenpool.LockReleaseTokenPool](activePool.GetCreatedEvent())
		if err != nil {
			return nil, fmt.Errorf("parse active lock/release pool %s: %w", poolAddress.Hex(), err)
		}

		remoteChainConfigsAny := make(map[types.NUMERIC]any, len(parsedPool.RemoteChainConfigs))
		for numeric, config := range parsedPool.RemoteChainConfigs {
			remoteChainConfigsAny[numeric] = any(config)
		}

		return &configuredCantonTokenPool{
			InstrumentId:       parsedPool.InstrumentId,
			InstanceId:         parsedPool.InstanceId,
			CcipOwner:          parsedPool.CcipOwner,
			PoolOwner:          parsedPool.PoolOwner,
			Decimals:           parsedPool.Decimals,
			RemoteChainConfigs: remoteChainConfigsAny,
		}, nil
	case burnMintPoolType:
		activePool, err := contract.FindActiveContractByInstanceAddress(
			ctx,
			participant.LedgerServices.State,
			participant.PartyID,
			burnminttokenpool.BurnMintTokenPool{}.GetTemplateID(),
			poolAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("find active burn/mint pool %s: %w", poolAddress.Hex(), err)
		}
		parsedPool, err := bindings.UnmarshalCreatedEvent[burnminttokenpool.BurnMintTokenPool](activePool.GetCreatedEvent())
		if err != nil {
			return nil, fmt.Errorf("parse active burn/mint pool %s: %w", poolAddress.Hex(), err)
		}

		remoteChainConfigsAny := make(map[types.NUMERIC]any, len(parsedPool.RemoteChainConfigs))
		for numeric, config := range parsedPool.RemoteChainConfigs {
			remoteChainConfigsAny[numeric] = any(config)
		}

		return &configuredCantonTokenPool{
			InstrumentId:       parsedPool.InstrumentId,
			InstanceId:         parsedPool.InstanceId,
			CcipOwner:          parsedPool.CcipOwner,
			PoolOwner:          parsedPool.PoolOwner,
			Decimals:           parsedPool.Decimals,
			RemoteChainConfigs: remoteChainConfigsAny,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported Canton token pool type %q", logicalPoolType)
	}
}

type DeployLockReleaseTokenPoolFromFactoryInput struct {
	ExistingDataStore datastore.DataStore
	ChainSelector     uint64
	ProposalDriven    bool
	CcipOwner         string
	PoolOwner         string
	InstrumentId      splice_api_token_holding_v1.InstrumentId
	Decimals          int64
	Deps              lockreleasetokenpool.LockReleaseTokenPoolDeps
	// RemoteChainSelectors lists CCIP remote chain selectors for which factory DeployRateLimiter runs
	// (outbound, inbound default-finality, inbound custom-finality) after DeployLockReleaseTokenPool.
	RemoteChainSelectors []uint64
}

func factoryDeployRateLimiterChoice(
	poolInstanceID contracts.InstanceID,
	poolOwner types.PARTY,
	remoteSelector types.NUMERIC,
	direction common.RateLimitDirection,
	mode common.RateLimitMode,
	instancePrefix string,
) (factorybindings.DeployRateLimiter, contracts.RawInstanceAddress, deployment.ContractType, error) {
	instID, err := ensureInstanceID("", instancePrefix)
	if err != nil {
		return factorybindings.DeployRateLimiter{}, "", "", err
	}
	tmpl := common.RateLimiter{
		InstanceId:          types.TEXT(instID),
		PoolInstanceId:      types.TEXT(poolInstanceID),
		PoolOwner:           poolOwner,
		RemoteChainSelector: remoteSelector,
		Direction:           direction,
		Mode:                mode,
		IsEnabled:           types.BOOL(false),
		Capacity:            types.NUMERIC("0"),
		Rate:                types.NUMERIC("0"),
		Tokens:              types.NUMERIC("0"),
		LastUpdated:         types.TIMESTAMP(time.Now()),
	}
	var ct deployment.ContractType
	if direction == common.RateLimitDirectionRateLimitDirection_Outbound {
		ct = rate_limiter.ContractTypeOutbound
	} else {
		ct = rate_limiter.ContractTypeInbound
	}
	raw := instID.RawInstanceAddress(poolOwner)
	return factorybindings.DeployRateLimiter{Contract: tmpl}, raw, ct, nil
}

var DeployLockReleaseTokenPoolFromFactory = operations.NewSequence(
	"canton/ccip/deploy_lock_release_token_pool_from_factory",
	semver.MustParse("2.0.0"),
	"Deploys a LockReleaseTokenPool from a CCIPFactory",
	func(b operations.Bundle, cantonChain cldfcanton.Chain, input DeployLockReleaseTokenPoolFromFactoryInput) (ccipsequences.OnChainOutput, error) {
		if input.ExistingDataStore == nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("existing datastore is required")
		}
		if input.ChainSelector != cantonChain.ChainSelector() {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("input chain selector %d does not match canton chain %d", input.ChainSelector, cantonChain.ChainSelector())
		}

		factoryRef, err := factoryops.ResolveFromDatastore(input.ExistingDataStore, input.ChainSelector)
		if err != nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("resolve CCIPFactory: %w", err)
		}
		factoryRaw, err := dsutils.GetRawInstanceAddressFromAddressRef(factoryRef)
		if err != nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("factory raw instance address: %w", err)
		}

		poolInstanceID, err := ensureInstanceID("", "lockreleasetokenpool")
		if err != nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("ensure pool instance ID: %w", err)
		}

		poolReceiveContext := splice_api_token_metadata_v1.ChoiceContext{Values: map[string]splice_api_token_metadata_v1.AnyValue{}}
		transferTimeout := lockreleasetokenpool.TransferTimeout{RelativeHours: new(types.INT64(24))}
		remoteChainConfigs := map[types.NUMERIC]lockreleasetokenpool.RemoteChainConfig{}
		tokenTransferFeeConfigs := map[types.NUMERIC]lockreleasetokenpool.TokenTransferFeeConfig2{}

		factoryDeployArgs := factorybindings.DeployLockReleaseTokenPool{
			Contract: lockreleasetokenpool.LockReleaseTokenPool{
				InstanceId:              types.TEXT(poolInstanceID),
				PoolOwner:               types.PARTY(input.PoolOwner),
				CcipOwner:               types.PARTY(input.CcipOwner),
				InstrumentId:            input.InstrumentId,
				Decimals:                types.INT64(input.Decimals),
				RemoteChainConfigs:      remoteChainConfigs,
				TokenTransferFeeConfigs: tokenTransferFeeConfigs,
				PoolReceiveContext:      poolReceiveContext,
				TransferTimeout:         transferTimeout,
				Deps:                    input.Deps,
			},
		}

		var proposalOutputs []contract.ExerciseOutput

		poolReport, err := operations.ExecuteOperation(b, factoryops.DeployLockReleaseTokenPool, cantonChain, newChoiceInput(factoryRaw, factoryDeployArgs, input.ProposalDriven))
		if err != nil {
			return ccipsequences.OnChainOutput{}, fmt.Errorf("deploy lock/release pool via factory: %w", err)
		}
		proposalOutputs = appendExerciseOutput(proposalOutputs, poolReport.Output, input.ProposalDriven)

		poolOwnerParty := types.PARTY(input.PoolOwner)
		poolRawAddr := poolInstanceID.RawInstanceAddress(poolOwnerParty)
		addresses := []datastore.AddressRef{
			newAddressRef(cantonChain.ChainSelector(), poolRawAddr, lock_release_token_pool.ContractType, lock_release_token_pool.Version, ""),
		}

		for _, remoteSel := range input.RemoteChainSelectors {
			remoteStr := strconv.FormatUint(remoteSel, 10)
			rs := types.NUMERIC(remoteStr)

			type rlStep struct {
				prefix string
				dir    common.RateLimitDirection
				mode   common.RateLimitMode
			}
			steps := []rlStep{
				{"outbound-rate-limiter", common.RateLimitDirectionRateLimitDirection_Outbound, common.RateLimitModeRateLimitMode_DefaultFinality},
				{"inbound-rate-limiter", common.RateLimitDirectionRateLimitDirection_Inbound, common.RateLimitModeRateLimitMode_DefaultFinality},
				{"inbound-custom-rate-limiter", common.RateLimitDirectionRateLimitDirection_Inbound, common.RateLimitModeRateLimitMode_CustomFinality},
			}
			for _, step := range steps {
				args, rlRaw, rlType, ferr := factoryDeployRateLimiterChoice(poolInstanceID, poolOwnerParty, rs, step.dir, step.mode, step.prefix)
				if ferr != nil {
					return ccipsequences.OnChainOutput{}, ferr
				}
				rep, opErr := operations.ExecuteOperation(b, factoryops.DeployRateLimiter, cantonChain, newChoiceInput(factoryRaw, args, input.ProposalDriven))
				if opErr != nil {
					return ccipsequences.OnChainOutput{}, fmt.Errorf("deploy rate limiter via factory (remote %s %s): %w", rs, step.prefix, opErr)
				}
				proposalOutputs = appendExerciseOutput(proposalOutputs, rep.Output, input.ProposalDriven)
				if !input.ProposalDriven && !rep.Output.Executed() {
					return ccipsequences.OnChainOutput{}, fmt.Errorf("factory deploy rate limiter was not executed on ledger (remote %s %s)", rs, step.prefix)
				}
				// Keys ignore Address; inbound default vs custom share Type. Qualifier must be unique per RL instance.
				rlQualifier := fmt.Sprintf("%s-%s-%s", remoteStr, step.prefix, rlRaw.InstanceID())
				addresses = append(addresses, newAddressRef(cantonChain.ChainSelector(), rlRaw, rlType, rate_limiter.Version, rlQualifier))
			}
		}

		if !input.ProposalDriven {
			if !poolReport.Output.Executed() {
				return ccipsequences.OnChainOutput{}, fmt.Errorf("factory deploy lock/release pool was not executed on ledger")
			}
		}

		out := ccipsequences.OnChainOutput{Addresses: addresses}
		if input.ProposalDriven {
			batchOp, berr := contract.NewBatchOperationFromExercises(proposalOutputs)
			if berr != nil {
				return ccipsequences.OnChainOutput{}, fmt.Errorf("build MCMS batch for factory pool and rate limiters: %w", berr)
			}
			if len(batchOp.Transactions) > 0 {
				out.BatchOps = []mcms_types.BatchOperation{batchOp}
			}
		}

		return out, nil
	},
)
