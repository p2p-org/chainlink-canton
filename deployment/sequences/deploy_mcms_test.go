package sequences

import (
	"sync"
	"testing"
	"time"

	participantv30 "github.com/digital-asset/dazl-client/v8/go/api/com/digitalasset/canton/admin/participant/v30"
	chainsel "github.com/smartcontractkit/chain-selectors"
	ccipdeploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/canton"
	cantonProvider "github.com/smartcontractkit/chainlink-deployments-framework/chain/canton/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/go-daml/pkg/types"
	"github.com/stretchr/testify/require"

	mcmsbindings "github.com/smartcontractkit/chainlink-canton/bindings/generated/mcms"
	"github.com/smartcontractkit/chainlink-canton/contracts"
	mcmsops "github.com/smartcontractkit/chainlink-canton/deployment/operations/mcms"
)

func TestDeployAndConfigureMCMS(t *testing.T) {
	t.Parallel()

	bc, err := cantonProvider.NewCTFChainProvider(t, chainsel.CANTON_LOCALNET.Selector, cantonProvider.CTFChainProviderConfig{
		NumberOfValidators: 1,
		Once:               &sync.Once{},
	}).Initialize(t.Context())
	require.NoError(t, err)

	cantonChain := bc.(*canton.Chain)
	participant := cantonChain.Participants[0]

	mcmsDar, err := contracts.GetDar(contracts.MCMS, contracts.CurrentVersion)
	require.NoError(t, err)
	_, err = participant.AdminServices.Package.UploadDar(t.Context(), &participantv30.UploadDarRequest{
		Dars: []*participantv30.UploadDarRequest_UploadDarData{
			{Bytes: mcmsDar},
		},
		VetAllPackages:     true,
		SynchronizeVetting: true,
	})
	require.NoError(t, err)

	bundle := cld_ops.NewBundle(t.Context, logger.Test(t), cld_ops.NewMemoryReporter())
	owner := participant.PartyID

	initial := MCMSConfigParams{
		Signers: []mcmsbindings.SignerInfo{
			{SignerAddress: types.TEXT("0x1111111111111111111111111111111111111111"), SignerIndex: 0, SignerGroup: 0},
			{SignerAddress: types.TEXT("0x2222222222222222222222222222222222222222"), SignerIndex: 1, SignerGroup: 0},
			{SignerAddress: types.TEXT("0x3333333333333333333333333333333333333333"), SignerIndex: 2, SignerGroup: 0},
		},
		GroupQuorums: []types.INT64{2},
		GroupParents: []types.INT64{0},
	}
	roleUpdate := MCMSRoleConfigParams{
		Role: mcmsbindings.RoleProposer,
		Config: MCMSConfigParams{
			Signers: []mcmsbindings.SignerInfo{
				{SignerAddress: types.TEXT("0x1111111111111111111111111111111111111111"), SignerIndex: 0, SignerGroup: 0},
				{SignerAddress: types.TEXT("0x2222222222222222222222222222222222222222"), SignerIndex: 1, SignerGroup: 0},
				{SignerAddress: types.TEXT("0x3333333333333333333333333333333333333333"), SignerIndex: 2, SignerGroup: 0},
				{SignerAddress: types.TEXT("0x4444444444444444444444444444444444444444"), SignerIndex: 3, SignerGroup: 0},
			},
			GroupQuorums: []types.INT64{3},
			GroupParents: []types.INT64{0},
		},
	}

	report, err := cld_ops.ExecuteSequence(bundle, DeployAndConfigureMCMS, *cantonChain, DeployAndConfigureMCMSParams{
		OwnerParty:    owner,
		InstanceID:    "mcms-seq-test",
		ChainID:       1,
		Qualifier:     "test",
		MinDelay:      time.Second,
		InitialConfig: initial,
		RoleConfigs:   []MCMSRoleConfigParams{roleUpdate},
	})
	require.NoError(t, err)

	require.Len(t, report.Output.Addresses, 5)
	require.True(t, hasRef(report.Output.Addresses, datastore.ContractType(mcmsops.ContractType), ""))
	require.True(t, hasRef(report.Output.Addresses, datastore.ContractType(ccipdeploymentutils.ProposerManyChainMultisig), "test"))
	require.True(t, hasRef(report.Output.Addresses, datastore.ContractType(ccipdeploymentutils.CancellerManyChainMultisig), "test"))
	require.True(t, hasRef(report.Output.Addresses, datastore.ContractType(ccipdeploymentutils.BypasserManyChainMultisig), "test"))
	require.True(t, hasRef(report.Output.Addresses, datastore.ContractType(ccipdeploymentutils.RBACTimelock), "test"))
}

func hasRef(refs []datastore.AddressRef, contractType datastore.ContractType, qualifier string) bool {
	for _, ref := range refs {
		if ref.Type == contractType && ref.Qualifier == qualifier {
			return true
		}
	}

	return false
}
