package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-canton/deployment/sequences"
)

var _ cldf.ChangeSetV2[CantonCSDeps[DeployChainContractsConfig]] = DeployChainContractsFromFactory{}

type DeployChainContractsFromFactory struct{}

func (d DeployChainContractsFromFactory) VerifyPreconditions(e cldf.Environment, config CantonCSDeps[DeployChainContractsConfig]) error {
	return verifyDeployChainContractsPreconditions(e, config)
}

func (d DeployChainContractsFromFactory) Apply(e cldf.Environment, config CantonCSDeps[DeployChainContractsConfig]) (cldf.ChangesetOutput, error) {
	ds := datastore.NewMemoryDataStore()

	chain := e.BlockChains.CantonChains()[config.ChainSelector]

	out, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployChainContractsFromFactory, chain, config.Config.Params)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to execute DeployChainContractsFromFactory sequence: %w", err)
	}

	for _, addrRef := range out.Output.Addresses {
		if err := ds.AddressRefStore.Add(addrRef); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to store address ref %v: %w", addrRef, err)
		}
	}

	return cldf.ChangesetOutput{
		DataStore: ds,
		Reports:   []operations.Report[any, any]{},
	}, nil
}
