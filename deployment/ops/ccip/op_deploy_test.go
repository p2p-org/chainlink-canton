package ccip

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	apiv2 "github.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v2"
	admin "github.com/digital-asset/dazl-client/v8/go/api/com/daml/ledger/api/v2/admin"
	"github.com/smartcontractkit/chainlink-canton-internal/contracts"
	"github.com/smartcontractkit/chainlink-canton-internal/deployment/client"
	participantv30 "github.com/smartcontractkit/chainlink-canton-internal/pb/gen/com/digitalasset/canton/admin/participant/v30"
)

const UserName = "ledger-api-user"

// TODO import these from testhelpers
type Participant struct {
	// Admin API
	PackageServiceClient participantv30.PackageServiceClient

	// Ledger API
	PartyManagementServiceClient admin.PartyManagementServiceClient
	UserManagementServiceClient  admin.UserManagementServiceClient

	StateServiceClient   apiv2.StateServiceClient
	CommandServiceClient apiv2.CommandServiceClient
	UpdateServiceClient  apiv2.UpdateServiceClient
	VersionServiceClient apiv2.VersionServiceClient
}

func NewParticipant(adminApiURL, ledgerApiURL string) (*Participant, error) {
	adminApiClient, err := grpc.NewClient(adminApiURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	ledgerApiClient, err := grpc.NewClient(ledgerApiURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Participant{
		PackageServiceClient: participantv30.NewPackageServiceClient(adminApiClient),

		PartyManagementServiceClient: admin.NewPartyManagementServiceClient(ledgerApiClient),
		UserManagementServiceClient:  admin.NewUserManagementServiceClient(ledgerApiClient),
		StateServiceClient:           apiv2.NewStateServiceClient(ledgerApiClient),
		CommandServiceClient:         apiv2.NewCommandServiceClient(ledgerApiClient),
		UpdateServiceClient:          apiv2.NewUpdateServiceClient(ledgerApiClient),
		VersionServiceClient:         apiv2.NewVersionServiceClient(ledgerApiClient),
	}, nil
}

func UploadDARstoMultipleParticipants(ctx context.Context, dars [][]byte, participants ...*Participant) ([]string, error) {
	var darData []*participantv30.UploadDarRequest_UploadDarData
	for _, dar := range dars {
		darData = append(darData, &participantv30.UploadDarRequest_UploadDarData{
			Bytes: dar,
		})
	}

	var packageIDs []string
	for i, participant := range participants {
		res, err := participant.PackageServiceClient.UploadDar(ctx, &participantv30.UploadDarRequest{
			Dars:               darData,
			VetAllPackages:     true,
			SynchronizeVetting: true,
		})
		if err != nil {
			return nil, fmt.Errorf("uploadDAR to participant %d failed: %w", i+1, err)
		}
		packageIDs = append(packageIDs, res.GetDarIds()...)
	}
	return packageIDs, nil
}

func EnsurePartyOnMultipleParticipants(ctx context.Context, participants ...*Participant) ([]string, error) {
	partyHints := []string{"alice", "bob", "charlie", "dave", "erin"}
	var partyIDs []string
	for i, participant := range participants {
		knownParties, err := participant.PartyManagementServiceClient.ListKnownParties(ctx, &admin.ListKnownPartiesRequest{})
		if err != nil {
			return nil, fmt.Errorf("listing known parties failed on participant %d: %w", i+1, err)
		}
		for _, details := range knownParties.GetPartyDetails() {
			if details.GetIsLocal() && strings.HasPrefix(details.GetParty(), fmt.Sprintf("%s::", partyHints[i])) {
				fmt.Printf("Found existing local party: %s\n", details.GetParty())
				partyIDs = append(partyIDs, details.GetParty())
				break
			}
		}
		if len(partyIDs) <= i {
			res, err := participant.PartyManagementServiceClient.AllocateParty(ctx, &admin.AllocatePartyRequest{
				PartyIdHint: partyHints[i],
			})
			if err != nil {
				return nil, fmt.Errorf("allocating party %s on participant %d failed: %w", partyHints[i], i+1, err)
			}
			fmt.Printf("Allocated new party on participant %v: %s\n", i+1, res.PartyDetails.Party)
			partyIDs = append(partyIDs, res.PartyDetails.Party)
		}

		// Grant user rights on party
		grantUserRightsResult, err := participant.UserManagementServiceClient.GrantUserRights(ctx, &admin.GrantUserRightsRequest{
			UserId: UserName,
			Rights: []*admin.Right{
				{Kind: &admin.Right_CanExecuteAsAnyParty_{CanExecuteAsAnyParty: &admin.Right_CanExecuteAsAnyParty{}}},
				{Kind: &admin.Right_CanReadAsAnyParty_{CanReadAsAnyParty: &admin.Right_CanReadAsAnyParty{}}},
				{Kind: &admin.Right_CanActAs_{CanActAs: &admin.Right_CanActAs{Party: partyIDs[i]}}},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("grantUserRights failed on participant %d: %w", i+1, err)
		}
		fmt.Printf("Granted user %q rights to act as party %q: %v\n", UserName, partyIDs[i], grantUserRightsResult.GetNewlyGrantedRights())
	}
	return partyIDs, nil
}

func getJWT() (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "",
		Subject:   UserName,
		Audience:  []string{"https://canton.network.global"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        "",
	})
	return t.SignedString([]byte("unsafe"))
}

func TestDeployCCIPContracts(t *testing.T) {
	t.Parallel()

	// setupResult, err := compileClient.Setup(ctx, compileClient.Config{
	// 	LedgerAPIURL:      "participant1.grpc-ledger-api.localhost:8080",
	// 	AdminAPIURL:       "participant1.admin-api.localhost:8080",
	// 	JWTSecret:         "unsafe",
	// 	DeployerParty:     "", // Empty to use primary party or allocate new one
	// 	DeployerPartyHint: "ledger-api-user",
	// })
	// require.NoError(t, err, "Failed to setup Canton client")

	// t.Cleanup(func() { setupResult.BindingClient.Close() })

	jwToken, err := getJWT()
	require.NoError(t, err)
	md := metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", jwToken))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	participant1, err := NewParticipant("participant1.admin-api.localhost:8080", "participant1.grpc-ledger-api.localhost:8080")
	require.NoError(t, err)

	version, err := participant1.VersionServiceClient.GetLedgerApiVersion(ctx, &apiv2.GetLedgerApiVersionRequest{})
	require.NoError(t, err)
	fmt.Println(version.Version)

	// Upload the DARs to all participants
	commonDar, err := contracts.GetDar(contracts.CCIPCommon, contracts.CurrentVersion)
	require.NoError(t, err)
	packageIDs, err := UploadDARstoMultipleParticipants(ctx, [][]byte{commonDar}, participant1)
	require.NoError(t, err)
	fmt.Printf("Uploaded CCIP Common DARs to all participants: %v\n", packageIDs)

	parties, err := EnsurePartyOnMultipleParticipants(ctx, participant1)
	require.NoError(t, err)
	fmt.Printf("Ensured parties on all participants: %v\n", parties)

	deps := client.CantonOpDepsForGRPC{
		BindingClient: participant1.CommandServiceClient,
		Party:         parties[0],
		UserID:        UserName,
	}

	reporter := cld_ops.NewMemoryReporter()

	bundle := cld_ops.NewBundle(
		func() context.Context { return ctx },
		logger.Test(t),
		reporter,
	)

	instanceID := "test-ccip-instance"
	chainSelectorValue := "1111111111"
	onRampAddress := "0000000000000000000000000000000000000001"

	// --------------------------
	// Test CCIP Common Deployment
	// --------------------------
	t.Run("DeployCCIPCommon", func(t *testing.T) {
		t.Parallel()

		result, err := cld_ops.ExecuteOperation(bundle, DeployCCIPCommonOp, deps, DeployCCIPCommonInput{
			InstanceID:         instanceID,
			ChainSelectorValue: chainSelectorValue,
			OnRampAddress:      onRampAddress,
		})
		require.NoError(t, err, "failed to deploy CCIP Common")
		require.NotEmpty(t, result.Output.Output.GlobalConfigContractID, "GlobalConfig contract ID should not be empty")
		require.NotEmpty(t, result.Output.Output.GlobalConfigTemplateID, "GlobalConfig template ID should not be empty")
		t.Logf("Deployed GlobalConfig contract ID: %s", result.Output.Output.GlobalConfigContractID)
	})

	// --------------------------
	// Test Token Admin Registry Deployment
	// --------------------------
	// t.Run("DeployTokenAdminRegistry", func(t *testing.T) {
	// 	t.Parallel()

	// 	result, err := cld_ops.ExecuteOperation(bundle, DeployTokenAdminRegistryOp, deps, DeployTokenAdminRegistryInput{
	// 		InstanceID: instanceID,
	// 	})
	// 	require.NoError(t, err, "failed to deploy TokenAdminRegistry")
	// 	require.NotEmpty(t, result.Output.Output.TokenAdminRegistryContractID, "TokenAdminRegistry contract ID should not be empty")
	// 	require.NotEmpty(t, result.Output.Output.TokenAdminRegistryTemplateID, "TokenAdminRegistry template ID should not be empty")
	// 	t.Logf("Deployed TokenAdminRegistry contract ID: %s", result.Output.Output.TokenAdminRegistryContractID)
	// })

	// // --------------------------
	// // Test Committee Verifier Deployment
	// // --------------------------
	// t.Run("DeployCommitteeVerifier", func(t *testing.T) {
	// 	t.Parallel()

	// 	result, err := cld_ops.ExecuteOperation(bundle, DeployCommitteeVerifierOp, deps, DeployCommitteeVerifierInput{
	// 		InstanceID:          instanceID,
	// 		VersionTag:          "1.0.0",
	// 		StorageLocation:     "ipfs://test-ccv",
	// 		Threshold:           2,
	// 		Signers:             []string{"signer1", "signer2", "signer3"},
	// 		MessageSentObserver: "", // Will default to deployer party
	// 	})
	// 	require.NoError(t, err, "failed to deploy CommitteeVerifier")
	// 	require.NotEmpty(t, result.Output.Output.CommitteeVerifierContractID, "CommitteeVerifier contract ID should not be empty")
	// 	require.NotEmpty(t, result.Output.Output.CommitteeVerifierTemplateID, "CommitteeVerifier template ID should not be empty")
	// 	t.Logf("Deployed CommitteeVerifier contract ID: %s", result.Output.Output.CommitteeVerifierContractID)
	// })

	// // --------------------------
	// // Test CCV Registry Deployment
	// // --------------------------
	// t.Run("DeployCCVRegistry", func(t *testing.T) {
	// 	t.Parallel()

	// 	result, err := cld_ops.ExecuteOperation(bundle, DeployCCVRegistryOp, deps, DeployCCVRegistryInput{
	// 		InstanceID: instanceID,
	// 	})
	// 	require.NoError(t, err, "failed to deploy CCVRegistry")
	// 	require.NotEmpty(t, result.Output.Output.CCVRegistryContractID, "CCVRegistry contract ID should not be empty")
	// 	require.NotEmpty(t, result.Output.Output.CCVRegistryTemplateID, "CCVRegistry template ID should not be empty")
	// 	t.Logf("Deployed CCVRegistry contract ID: %s", result.Output.Output.CCVRegistryContractID)
	// })

	// // --------------------------
	// // Test Fee Quoter Deployment
	// // --------------------------
	// t.Run("DeployFeeQuoter", func(t *testing.T) {
	// 	t.Parallel()

	// 	result, err := cld_ops.ExecuteOperation(bundle, DeployFeeQuoterOp, deps, DeployFeeQuoterInput{
	// 		InstanceID: instanceID,
	// 	})
	// 	require.NoError(t, err, "failed to deploy FeeQuoter")
	// 	require.NotEmpty(t, result.Output.Output.FeeQuoterContractID, "FeeQuoter contract ID should not be empty")
	// 	require.NotEmpty(t, result.Output.Output.FeeQuoterTemplateID, "FeeQuoter template ID should not be empty")
	// 	t.Logf("Deployed FeeQuoter contract ID: %s", result.Output.Output.FeeQuoterContractID)
	// })

	// // --------------------------
	// // Test OffRamp Deployment
	// // --------------------------
	// t.Run("DeployOffRamp", func(t *testing.T) {
	// 	t.Parallel()

	// 	result, err := cld_ops.ExecuteOperation(bundle, DeployOffRampOp, deps, DeployOffRampInput{
	// 		InstanceID: instanceID,
	// 	})
	// 	require.NoError(t, err, "failed to deploy OffRamp")
	// 	require.NotEmpty(t, result.Output.Output.OffRampContractID, "OffRamp contract ID should not be empty")
	// 	require.NotEmpty(t, result.Output.Output.OffRampTemplateID, "OffRamp template ID should not be empty")
	// 	t.Logf("Deployed OffRamp contract ID: %s", result.Output.Output.OffRampContractID)
	// })

	// // --------------------------
	// // Test PerPartyRouter Deployment
	// // --------------------------
	// t.Run("DeployPerPartyRouter", func(t *testing.T) {
	// 	t.Parallel()

	// 	result, err := cld_ops.ExecuteOperation(bundle, DeployPerPartyRouterOp, deps, DeployPerPartyRouterInput{
	// 		InstanceID: instanceID,
	// 	})
	// 	require.NoError(t, err, "failed to deploy PerPartyRouter")
	// 	require.NotEmpty(t, result.Output.Output.PerPartyRouterContractID, "PerPartyRouter contract ID should not be empty")
	// 	require.NotEmpty(t, result.Output.Output.PerPartyRouterTemplateID, "PerPartyRouter template ID should not be empty")
	// 	t.Logf("Deployed PerPartyRouter contract ID: %s", result.Output.Output.PerPartyRouterContractID)
	// })

	// // --------------------------
	// // Test OnRamp Deployment
	// // --------------------------
	// t.Run("DeployOnRamp", func(t *testing.T) {
	// 	t.Parallel()

	// 	result, err := cld_ops.ExecuteOperation(bundle, DeployOnRampOp, deps, DeployOnRampInput{
	// 		InstanceID:           instanceID,
	// 		DestChainSelector:    "2222222222",
	// 		DestChainOnRampBytes: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
	// 	})
	// 	require.NoError(t, err, "failed to deploy OnRamp")
	// 	require.NotEmpty(t, result.Output.Output.OnRampContractID, "OnRamp contract ID should not be empty")
	// 	require.NotEmpty(t, result.Output.Output.OnRampTemplateID, "OnRamp template ID should not be empty")
	// 	t.Logf("Deployed OnRamp contract ID: %s", result.Output.Output.OnRampContractID)
	// })
}
