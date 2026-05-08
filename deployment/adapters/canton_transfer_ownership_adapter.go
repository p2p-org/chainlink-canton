package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// CantonTransferOwnershipAdapter satisfies deploy.TransferOwnershipAdapter for Canton chains.
//
// CCIP's DeployChainContracts always invokes deploy.TransferToTimelock when DeployerKeyOwned is false.
// On Canton, product ownership moves to governance inside proposal-driven DAML flows (see
// DeployChainContractsFromFactory); the generic deploy adapter does not populate RefsToTransferOwnership.
// This adapter is therefore a deliberate no-op for the empty-ref case so the CCIP changeset can complete.
// If RefsToTransferOwnership is ever populated for Canton, implement real transfer batching here.
type CantonTransferOwnershipAdapter struct{}

func (a *CantonTransferOwnershipAdapter) InitializeTimelockAddress(_ deployment.Environment, _ mcms.Input) error {
	return nil
}

func (a *CantonTransferOwnershipAdapter) SequenceTransferOwnershipViaMCMS() *cldf_ops.Sequence[deploy.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"canton-seq-transfer-ownership-via-mcms",
		semver.MustParse("1.0.0"),
		"No-op: Canton ownership for CCIP contracts is handled in Canton deploy sequences / MCMS proposals, not via the EVM timelock transfer path.",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, in deploy.TransferOwnershipPerChainInput) (sequences.OnChainOutput, error) {
			if len(in.ContractRef) > 0 {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"canton: generic TransferToTimelock is not implemented for %d refs on chain %d (RefsToTransferOwnership should be empty for Canton v2 deploy)",
					len(in.ContractRef), in.ChainSelector,
				)
			}
			return sequences.OnChainOutput{}, nil
		},
	)
}

func (a *CantonTransferOwnershipAdapter) ShouldAcceptOwnershipWithTransferOwnership(_ deployment.Environment, _ deploy.TransferOwnershipPerChainInput) (bool, error) {
	return false, nil
}

func (a *CantonTransferOwnershipAdapter) SequenceAcceptOwnership() *cldf_ops.Sequence[deploy.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"canton-seq-accept-ownership",
		semver.MustParse("1.0.0"),
		"No-op: Canton does not use the EVM accept-ownership MCMS batch path.",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ deploy.TransferOwnershipPerChainInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, nil
		},
	)
}
