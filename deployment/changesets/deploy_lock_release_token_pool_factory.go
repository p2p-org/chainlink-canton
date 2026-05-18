package changesets

import (
	"fmt"

	"github.com/smartcontractkit/mcms"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/lockreleasetokenpool"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/splice/splice_api_token_holding_v1"
	factoryops "github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/factory"
	mcmsops "github.com/smartcontractkit/chainlink-canton/deployment/operations/mcms"
	"github.com/smartcontractkit/chainlink-canton/deployment/sequences"
	cantonmcms "github.com/smartcontractkit/chainlink-canton/deployment/utils/mcms"
)

// DeployLockReleaseTokenPoolFactoryConfig is the config for deploying a LockReleaseTokenPool
// via the CCIPFactory. When MCMSEnabled is set, MCMSTimelock configures the timelock proposal.
type DeployLockReleaseTokenPoolFactoryConfig struct {
	CcipOwner    string
	PoolOwner    string
	InstrumentId splice_api_token_holding_v1.InstrumentId
	Decimals     int64
	MCMSEnabled  bool
	Deps         lockreleasetokenpool.LockReleaseTokenPoolDeps
	// RemoteChainSelectors selects CCIP remote chains for which the factory deploys outbound/inbound rate limiters after the pool.
	RemoteChainSelectors []uint64
	MCMSTimelock         MCMSTimelockConfig
}

var _ cldf.ChangeSetV2[CantonCSDeps[DeployLockReleaseTokenPoolFactoryConfig]] = DeployLockReleaseTokenPoolFactory{}

type DeployLockReleaseTokenPoolFactory struct{}

func (d DeployLockReleaseTokenPoolFactory) VerifyPreconditions(e cldf.Environment, config CantonCSDeps[DeployLockReleaseTokenPoolFactoryConfig]) error {
	chain, ok := e.BlockChains.CantonChains()[config.ChainSelector]
	if !ok {
		return fmt.Errorf("canton chain %v not found", config.ChainSelector)
	}
	if config.Participant < 0 || config.Participant >= len(chain.Participants) {
		return fmt.Errorf("participant index %d out of range for canton chain %d with %d participants", config.Participant, config.ChainSelector, len(chain.Participants))
	}

	if _, err := factoryops.ResolveFromDatastore(e.DataStore, config.ChainSelector); err != nil {
		return fmt.Errorf("CCIPFactory not found in datastore: %w", err)
	}

	if config.Config.MCMSEnabled {
		if config.Config.MCMSTimelock.MCMSContract.RawInstanceAddress == "" {
			return fmt.Errorf("MCMSTimelock.MCMSContract is required when MCMSEnabled is true")
		}
		if _, err := e.DataStore.Addresses().Get(datastore.NewAddressRefKey(
			config.ChainSelector,
			datastore.ContractType(mcmsops.ContractType),
			mcmsops.Version,
			"",
		)); err != nil {
			return fmt.Errorf("MCMS contract not found in datastore (deploy MCMS first): %w", err)
		}
	}

	return nil
}

func (d DeployLockReleaseTokenPoolFactory) Apply(e cldf.Environment, config CantonCSDeps[DeployLockReleaseTokenPoolFactoryConfig]) (cldf.ChangesetOutput, error) {
	ds := datastore.NewMemoryDataStore()

	chain := e.BlockChains.CantonChains()[config.ChainSelector]
	cfg := config.Config

	seqOut, err := cld_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployLockReleaseTokenPoolFromFactory, chain, sequences.DeployLockReleaseTokenPoolFromFactoryInput{
		ExistingDataStore:    e.DataStore,
		ChainSelector:        config.ChainSelector,
		ProposalDriven:       cfg.MCMSEnabled,
		CcipOwner:            cfg.CcipOwner,
		PoolOwner:            cfg.PoolOwner,
		InstrumentId:         cfg.InstrumentId,
		Decimals:             cfg.Decimals,
		Deps:                 cfg.Deps,
		RemoteChainSelectors: cfg.RemoteChainSelectors,
	})
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("deploy lock/release token pool from factory sequence: %w", err)
	}

	for _, addrRef := range seqOut.Output.Addresses {
		if err := ds.AddressRefStore.Upsert(addrRef); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to store address ref: %w", err)
		}
	}

	if cfg.MCMSEnabled {
		if len(seqOut.Output.BatchOps) == 0 {
			return cldf.ChangesetOutput{}, fmt.Errorf("MCMS enabled but sequence returned no batch operations")
		}
		tl := cfg.MCMSTimelock
		participant := chain.Participants[config.Participant]
		proposal, err := cantonmcms.GenerateTimelockProposal(
			e.GetContext(),
			participant.LedgerServices.State,
			participant.PartyID,
			cantonmcms.ProposalConfig{
				MCMSContract:         tl.MCMSContract,
				ChainSelector:        mcms_types.ChainSelector(config.ChainSelector),
				Description:          tl.Description,
				MinDelay:             tl.MinDelay,
				OverridePreviousRoot: tl.OverridePrevRoot,
				Action:               tl.Action,
				Role:                 tl.Role,
			},
			seqOut.Output.BatchOps,
		)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("generate timelock proposal: %w", err)
		}
		return cldf.ChangesetOutput{
			DataStore:             ds,
			MCMSTimelockProposals: []mcms.TimelockProposal{*proposal},
		}, nil
	}

	return cldf.ChangesetOutput{
		DataStore: ds,
		Reports:   []cld_ops.Report[any, any]{},
	}, nil
}
