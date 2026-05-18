package changesets

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/feequoter"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/lockreleasetokenpool"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/rmn"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/tokenadminregistry"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/splice/splice_api_token_holding_v1"
	"github.com/smartcontractkit/chainlink-canton/contracts"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/fee_quoter"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/rmn_remote"
	"github.com/smartcontractkit/chainlink-canton/deployment/operations/ccip/token_admin_registry"
	cantonmcms "github.com/smartcontractkit/chainlink-canton/deployment/utils/mcms"
	opcontract "github.com/smartcontractkit/chainlink-canton/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/canton"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/go-daml/pkg/types"
	cantonsdk "github.com/smartcontractkit/mcms/sdk/canton"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDeployLockReleaseTokenPoolFactory_MCMSProposal deploys CCIPFactory, LRTP dependencies,
// and MCMS on a CTF Canton ledger, then runs DeployLockReleaseTokenPoolFactory with MCMSEnabled.
// The sequence builds a batch of factory exercises (pool + rate limiters); the changeset turns
// that batch into an MCMS timelock proposal without executing those exercises on-chain.
//
// Like TestConfigureGlobalConfig_MCMSProposal, this requires a local Canton CTF stack (see setupCantonEnv).
func TestDeployLockReleaseTokenPoolFactory_MCMSProposal(t *testing.T) {
	t.Parallel()

	cantonChain, bundle, env := setupCantonEnv(t)
	participant := cantonChain.Participants[0]
	party := participant.PartyID

	uploadChainContractDARs(t, participant)
	uploadDARs(t, participant, contracts.CCIPFactory, contracts.MCMS)

	deps := deployLRTPDeps(t, bundle, *cantonChain, party)

	factoryOut, err := DeployCCIPFactory{}.Apply(*env, CantonCSDeps[DeployCCIPFactoryConfig]{
		ChainSelector: chainsel.CANTON_LOCALNET.Selector,
		Participant:   0,
		Config: DeployCCIPFactoryConfig{
			Params: DeployCCIPFactoryParams{
				OwnerParty: party,
			},
		},
	})
	require.NoError(t, err)

	factoryRefs := factoryOut.DataStore.Addresses().Filter()
	require.Len(t, factoryRefs, 1, "DeployCCIPFactory should persist exactly one address ref")
	factoryRaw, err := contracts.RawInstanceAddressFromString(factoryRefs[0].Labels.List()[0])
	require.NoError(t, err)

	mcmsAddrRef := deployMCMSContract(t, bundle, *cantonChain, party)
	mcmsRawAddr, err := contracts.RawInstanceAddressFromString(mcmsAddrRef.Labels.List()[0])
	require.NoError(t, err)

	merged := datastore.NewMemoryDataStore()
	for _, ref := range factoryRefs {
		require.NoError(t, merged.AddressRefStore.Add(ref))
	}
	require.NoError(t, merged.AddressRefStore.Add(mcmsAddrRef))
	env.DataStore = merged.Seal()

	csDeps := CantonCSDeps[DeployLockReleaseTokenPoolFactoryConfig]{
		ChainSelector: chainsel.CANTON_LOCALNET.Selector,
		Participant:   0,
		Config: DeployLockReleaseTokenPoolFactoryConfig{
			CcipOwner:            party,
			PoolOwner:            party,
			InstrumentId:         splice_api_token_holding_v1.InstrumentId{Admin: types.PARTY(party), Id: types.TEXT("LRTP-FACTORY-TEST")},
			Decimals:             6,
			MCMSEnabled:          true,
			Deps:                 deps,
			RemoteChainSelectors: []uint64{999},
			MCMSTimelock: MCMSTimelockConfig{
				MinDelay:         10 * time.Minute,
				Description:      "deploy LRTP and rate limiters via CCIPFactory (MCMS)",
				OverridePrevRoot: true,
				Action:           mcms_types.TimelockActionSchedule,
				MCMSContract: cantonmcms.MCMSContractInfo{
					RawInstanceAddress: mcmsRawAddr,
					InstanceAddress:    mcmsRawAddr.InstanceAddress(),
				},
				Role: cantonsdk.TimelockRoleProposer,
			},
		},
	}

	require.NoError(t, DeployLockReleaseTokenPoolFactory{}.VerifyPreconditions(*env, csDeps))

	out, err := DeployLockReleaseTokenPoolFactory{}.Apply(*env, csDeps)
	require.NoError(t, err)

	fmt.Println("Out: ", out.MCMSTimelockProposals)

	require.Len(t, out.MCMSTimelockProposals, 1)
	proposal := out.MCMSTimelockProposals[0]
	assert.Equal(t, "v1", proposal.Version)
	assert.Equal(t, "deploy LRTP and rate limiters via CCIPFactory (MCMS)", proposal.Description)
	assert.Equal(t, mcms_types.TimelockActionSchedule, proposal.Action)
	require.Len(t, proposal.Operations, 1)

	batchOp := proposal.Operations[0]
	assert.Equal(t, mcms_types.ChainSelector(chainsel.CANTON_LOCALNET.Selector), batchOp.ChainSelector)

	// One DeployLockReleaseTokenPool + three DeployRateLimiter per remote chain selector.
	require.Len(t, batchOp.Transactions, 4)

	wantFuncs := []string{
		"DeployLockReleaseTokenPool",
		"DeployRateLimiter",
		"DeployRateLimiter",
		"DeployRateLimiter",
	}
	for i, tx := range batchOp.Transactions {
		assert.Equal(t, factoryRaw.InstanceAddress().Hex(), tx.To)
		assert.NotEmpty(t, tx.Data)

		var af cantonsdk.AdditionalFields
		require.NoError(t, json.Unmarshal(tx.AdditionalFields, &af))
		assert.Equal(t, factoryRaw.String(), af.TargetInstanceAddress)
		assert.Equal(t, wantFuncs[i], af.FunctionName)
		assert.NotEmpty(t, af.OperationData)
	}

	require.Len(t, out.DataStore.Addresses().Filter(), 4, "pool + 3 rate limiter address refs")
}

func deployLRTPDeps(t *testing.T, bundle cld_ops.Bundle, chain canton.Chain, party string) lockreleasetokenpool.LockReleaseTokenPoolDeps {
	t.Helper()

	rmnReport, err := cld_ops.ExecuteOperation(bundle, rmn_remote.Deploy, chain, opcontract.DeployInput[rmn.RMNRemote]{
		Template: rmn.RMNRemote{
			RmnOwner:       types.PARTY(party),
			CcipOwner:      types.PARTY(party),
			CursedSubjects: nil,
		},
		OwnerParty: types.PARTY(party),
	})
	require.NoError(t, err)

	tarReport, err := cld_ops.ExecuteOperation(bundle, token_admin_registry.Deploy, chain, opcontract.DeployInput[tokenadminregistry.TokenAdminRegistry]{
		Template: tokenadminregistry.TokenAdminRegistry{
			Owner:      types.PARTY(party),
			EntryCount: 0,
		},
		OwnerParty: types.PARTY(party),
	})
	require.NoError(t, err)

	linkTokenID := splice_api_token_holding_v1.InstrumentId{
		Admin: types.PARTY(party),
		Id:    types.TEXT("link-token"),
	}
	fqReport, err := cld_ops.ExecuteOperation(bundle, fee_quoter.Deploy, chain, opcontract.DeployInput[feequoter.FeeQuoter]{
		Template: feequoter.FeeQuoter{
			Owner:                            types.PARTY(party),
			FeeTokens:                        types.SET{},
			DestChainConfigs:                 nil,
			TokenTransferFeeConfigs:          nil,
			UsdPerUnitGasByDestChainSelector: nil,
			UsdPerToken:                      nil,
			LinkTokenInstrumentId:            linkTokenID,
			PriceUpdaters:                    nil,
		},
		OwnerParty: types.PARTY(party),
	})
	require.NoError(t, err)

	rmnRaw, err := contracts.RawInstanceAddressFromString(rmnReport.Output.Labels.List()[0])
	require.NoError(t, err)
	tarRaw, err := contracts.RawInstanceAddressFromString(tarReport.Output.Labels.List()[0])
	require.NoError(t, err)
	fqRaw, err := contracts.RawInstanceAddressFromString(fqReport.Output.Labels.List()[0])
	require.NoError(t, err)

	return lockreleasetokenpool.LockReleaseTokenPoolDeps{
		TokenAdminRegistry: tarRaw.Binding(),
		RmnRemote:          rmnRaw.Binding(),
		FeeQuoter:          fqRaw.Binding(),
	}
}
