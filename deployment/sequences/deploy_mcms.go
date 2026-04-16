package sequences

import (
	"fmt"
	"time"

	"github.com/Masterminds/semver/v3"
	ccipdeploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/canton"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/go-daml/pkg/types"

	mcmsbindings "github.com/smartcontractkit/chainlink-canton/bindings/generated/mcms"
	"github.com/smartcontractkit/chainlink-canton/contracts"
	mcmsops "github.com/smartcontractkit/chainlink-canton/deployment/operations/mcms"
	opcontract "github.com/smartcontractkit/chainlink-canton/deployment/utils/operations/contract"
)

const mcmsGroupCount = 32

type MCMSConfigParams struct {
	Signers      []mcmsbindings.SignerInfo
	GroupQuorums []types.INT64
	GroupParents []types.INT64
	ClearRoot    bool
}

type MCMSRoleConfigParams struct {
	Role   mcmsbindings.Role
	Config MCMSConfigParams
}

type DeployAndConfigureMCMSParams struct {
	OwnerParty       string
	InstanceID       string
	ChainID          int64
	Qualifier        string
	MinDelay         time.Duration
	BlockedFunctions []mcmsbindings.BlockedFunction
	InitialConfig    MCMSConfigParams
	RoleConfigs      []MCMSRoleConfigParams
}

var DeployAndConfigureMCMS = operations.NewSequence(
	"canton/mcms/deploy_and_configure",
	semver.MustParse("0.1.0"),
	"Deploys and configures a Canton MCMS contract",
	func(b operations.Bundle, deps canton.Chain, input DeployAndConfigureMCMSParams) (sequences.OnChainOutput, error) {
		if input.OwnerParty == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("owner party is required")
		}
		if input.InstanceID == "" {
			return sequences.OnChainOutput{}, fmt.Errorf("instance ID is required")
		}
		if input.ChainID <= 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("chain ID must be greater than zero")
		}

		initialConfig, err := buildMultisigConfig(input.InitialConfig)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("build initial MCMS config: %w", err)
		}

		roleState := emptyRoleState(initialConfig)
		ownerParty := types.PARTY(input.OwnerParty)

		deployReport, err := operations.ExecuteOperation(b, mcmsops.Deploy, deps, opcontract.DeployInput[mcmsbindings.MCMS]{
			Template: mcmsbindings.MCMS{
				Owner:              ownerParty,
				InstanceId:         types.TEXT(input.InstanceID),
				ChainId:            types.INT64(input.ChainID),
				Proposer:           roleState,
				Canceller:          roleState,
				Bypasser:           roleState,
				MinDelay:           types.RELTIME(input.MinDelay),
				BlockedFunctions:   input.BlockedFunctions,
				TimelockTimestamps: types.GENMAP{},
			},
			OwnerParty: ownerParty,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("deploy MCMS: %w", err)
		}

		rawInstanceAddress := contracts.InstanceID(input.InstanceID).RawInstanceAddress(ownerParty)
		for i, roleConfig := range input.RoleConfigs {
			groupConfig, err := buildNormalizedConfig(roleConfig.Config)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("build MCMS config for role %s: %w", roleConfig.Role, err)
			}

			_, err = operations.ExecuteOperation(b, mcmsops.SetConfig, deps, opcontract.ChoiceInput[mcmsbindings.SetConfig]{
				InstanceAddress: rawInstanceAddress.InstanceAddress(),
				Args: mcmsbindings.SetConfig{
					TargetRole:      roleConfig.Role,
					NewSigners:      groupConfig.Signers,
					NewGroupQuorums: groupConfig.GroupQuorums,
					NewGroupParents: groupConfig.GroupParents,
					ClearRoot:       types.BOOL(roleConfig.Config.ClearRoot),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("configure MCMS role %s at index %d: %w", roleConfig.Role, i, err)
			}
		}

		refs := []datastore.AddressRef{
			deployReport.Output,
			newMCMSRoleAddressRef(deps.ChainSelector(), rawInstanceAddress, datastore.ContractType(ccipdeploymentutils.ProposerManyChainMultisig), qualifierOrDefault(input.Qualifier)),
			newMCMSRoleAddressRef(deps.ChainSelector(), rawInstanceAddress, datastore.ContractType(ccipdeploymentutils.CancellerManyChainMultisig), qualifierOrDefault(input.Qualifier)),
			newMCMSRoleAddressRef(deps.ChainSelector(), rawInstanceAddress, datastore.ContractType(ccipdeploymentutils.BypasserManyChainMultisig), qualifierOrDefault(input.Qualifier)),
			newMCMSRoleAddressRef(deps.ChainSelector(), rawInstanceAddress, datastore.ContractType(ccipdeploymentutils.RBACTimelock), qualifierOrDefault(input.Qualifier)),
		}

		return sequences.OnChainOutput{Addresses: refs}, nil
	},
)

func buildMultisigConfig(input MCMSConfigParams) (mcmsbindings.MultisigConfig, error) {
	cfg, err := buildNormalizedConfig(input)
	if err != nil {
		return mcmsbindings.MultisigConfig{}, err
	}

	return mcmsbindings.MultisigConfig{
		Signers:      cfg.Signers,
		GroupQuorums: cfg.GroupQuorums,
		GroupParents: cfg.GroupParents,
	}, nil
}

func buildNormalizedConfig(input MCMSConfigParams) (MCMSConfigParams, error) {
	groupQuorums, err := normalizeGroups(input.GroupQuorums, "group quorums")
	if err != nil {
		return MCMSConfigParams{}, err
	}
	groupParents, err := normalizeGroups(input.GroupParents, "group parents")
	if err != nil {
		return MCMSConfigParams{}, err
	}

	return MCMSConfigParams{
		Signers:      input.Signers,
		GroupQuorums: groupQuorums,
		GroupParents: groupParents,
		ClearRoot:    input.ClearRoot,
	}, nil
}

func normalizeGroups(input []types.INT64, name string) ([]types.INT64, error) {
	if len(input) > mcmsGroupCount {
		return nil, fmt.Errorf("%s length %d exceeds max %d", name, len(input), mcmsGroupCount)
	}

	out := make([]types.INT64, mcmsGroupCount)
	copy(out, input)

	return out, nil
}

func emptyRoleState(config mcmsbindings.MultisigConfig) mcmsbindings.RoleState {
	return mcmsbindings.RoleState{
		Config:     config,
		SeenHashes: types.GENMAP{},
		ExpiringRoot: mcmsbindings.ExpiringRoot{
			Root:       types.TEXT(""),
			ValidUntil: types.TIMESTAMP(time.Unix(0, 0)),
			OpCount:    types.INT64(0),
		},
		RootMetadata: mcmsbindings.RootMetadata{
			ChainId:              types.INT64(0),
			MultisigId:           types.TEXT(""),
			PreOpCount:           types.INT64(0),
			PostOpCount:          types.INT64(0),
			OverridePreviousRoot: types.BOOL(false),
		},
	}
}

func qualifierOrDefault(qualifier string) string {
	if qualifier != "" {
		return qualifier
	}

	return ccipdeploymentutils.CLLQualifier
}

func newMCMSRoleAddressRef(
	chainSelector uint64,
	rawInstanceAddress contracts.RawInstanceAddress,
	contractType datastore.ContractType,
	qualifier string,
) datastore.AddressRef {
	return datastore.AddressRef{
		Address:       rawInstanceAddress.InstanceAddress().String(),
		Labels:        datastore.NewLabelSet(rawInstanceAddress.String()),
		ChainSelector: chainSelector,
		Type:          contractType,
		Version:       mcmsops.Version,
		Qualifier:     qualifier,
	}
}
