package contracts

import (
	"embed"
	"fmt"
	"slices"
)

//go:embed dars
var Dars embed.FS

//go:embed dependencies/splice
var SpliceDependencies embed.FS

type Package string

const (
	Coin = Package("coin")
	Link = Package("link")

	ChainlinkAPI         = Package("chainlink-api")
	ChainlinkInstanceAPI = ChainlinkAPI

	MCMSAPI      = Package("mcms-api")
	MCMSCore     = Package("mcms-core")
	MCMS         = MCMSCore
	MCMSTest     = Package("mcms-test")
	GlobalConfig = Package("globalconfig")

	CCIPCore                 = Package("ccip-core")
	CCIPExtensionAPI         = Package("ccip-extension-api")
	CCIPRuntime              = Package("ccip-runtime")
	CCIPMain                 = CCIPRuntime
	CCIPSender               = Package("ccip-sender")
	CCIPReceiver             = Package("ccip-receiver")
	CCIPCommitteeVerifier    = Package("ccip-committee-verifier")
	CCIPExecutor             = Package("ccip-executor")
	CCIPLockReleaseTokenPool = Package("ccip-lock-release-token-pool")
	CCIPBurnMintTokenPool    = Package("ccip-burn-mint-token-pool")
	CCIPFactory              = Package("ccip-factory")
	CCIPTest                 = Package("ccip-test")

	CCIPClient             = CCIPCore
	CCIPCommon             = CCIPCore
	CCIPFeeQuoter          = CCIPCore
	CCIPTokenAdminRegistry = CCIPCore
	CCIPRMN                = CCIPCore
	CCIPOnRamp             = CCIPRuntime
	CCIPOffRamp            = CCIPRuntime
	CCIPPerPartyRouter     = CCIPRuntime
	CCIPPoolInterfaces     = CCIPExtensionAPI

	SpliceApiFeaturedAppV1                = Package("splice-api-featured-app-v1")
	SpliceApiTokenAllocationV1            = Package("splice-api-token-allocation-v1")
	SpliceApiTokenAllocationInstructionV1 = Package("splice-api-token-allocation-instruction-v1")
	SpliceApiTokenBurnMintV1              = Package("splice-api-token-burn-mint-v1")
	SpliceApiTokenHoldingV1               = Package("splice-api-token-holding-v1")
	SpliceApiTokenMetadataV1              = Package("splice-api-token-metadata-v1")
	SpliceApiTokenTransferInstructionV1   = Package("splice-api-token-transfer-instruction-v1")

	// Canton Network Utility DARs (bundle 0.12.5). Package IDs pinned in dar-versions.md.
	UtilityCredentialV0      = Package("utility-credential-v0")
	UtilityRegistryV0        = Package("utility-registry-v0")
	UtilityRegistryHoldingV0 = Package("utility-registry-holding-v0")
	UtilityRegistryAppV0     = Package("utility-registry-app-v0")
)

// Pinned package IDs from canton-network-utility-dars-0.12.5 / dar-versions.md.
//
//nolint:gosec // G101: These are safe package IDs
const (
	UtilityCommercialsV0PackageID     = "fa5b1cc5c8368dff7c2e6a74aa2af9d520d755e2a508f44acd17343326e41839"
	UtilityCredentialAppV0PackageID   = "e9a3b7df354dfd2f15c7d015328c34256308c90ba96f86f185dad58ffca8299b"
	UtilityCredentialV0PackageID      = "5a29ead611a0abd5f5b3fc3caf7d0f67c0ff802032ab6d392824aa9060e56d70"
	UtilityRegistryAppV0PackageID     = "7a75ef6e69f69395a4e60919e228528bb8f3881150ccfde3f31bcc73864b18ab"
	UtilityRegistryV0PackageID        = "a236e8e22a3b5f199e37d5554e82bafd2df688f901de02b00be3964bdfa8c1ab"
	UtilityRegistryHoldingV0PackageID = "8107899ac4723ce986bf7d27416534e576e54b92161e46150a595fb78ff3d3a1"
)

const CurrentVersion = "current"

// ReleaseDir is the frozen production DAR snapshot (e.g. v2_0_0).
// Individual packages keep their own semver in the filename; multiple package
// versions (e.g. globalconfig-1.0.0 and globalconfig-2.0.0) live in the same
// release directory.
const ReleaseDir = "v2_0_0"

var Versions map[Package][]string = map[Package][]string{
	Coin:         []string{CurrentVersion},
	Link:         []string{"2.0.0", CurrentVersion},
	ChainlinkAPI: []string{"2.0.0", CurrentVersion},
	MCMSAPI:      []string{"2.0.0", CurrentVersion},
	MCMSCore:     []string{"2.0.0", CurrentVersion},
	MCMSTest:     []string{CurrentVersion},
	GlobalConfig: []string{"1.0.0", "2.0.0", CurrentVersion},

	CCIPCore:                 []string{"2.0.0", CurrentVersion},
	CCIPExtensionAPI:         []string{"2.0.0", CurrentVersion},
	CCIPRuntime:              []string{"2.0.0", CurrentVersion},
	CCIPSender:               []string{"2.0.0", CurrentVersion},
	CCIPReceiver:             []string{"2.0.0", CurrentVersion},
	CCIPCommitteeVerifier:    []string{"2.0.0", CurrentVersion},
	CCIPExecutor:             []string{"2.0.0", CurrentVersion},
	CCIPLockReleaseTokenPool: []string{"2.0.0", CurrentVersion},
	CCIPBurnMintTokenPool:    []string{"2.0.0", CurrentVersion},
	CCIPFactory:              []string{"2.0.0", CurrentVersion},
	CCIPTest:                 []string{CurrentVersion},

	SpliceApiFeaturedAppV1:                []string{"1.0.0"},
	SpliceApiTokenAllocationV1:            []string{"1.0.0"},
	SpliceApiTokenAllocationInstructionV1: []string{"1.0.0"},
	SpliceApiTokenBurnMintV1:              []string{"1.0.0"},
	SpliceApiTokenHoldingV1:               []string{"1.0.0"},
	SpliceApiTokenMetadataV1:              []string{"1.0.0"},
	SpliceApiTokenTransferInstructionV1:   []string{"1.0.0"},

	// Vendored from canton-network-utility-dars-0.12.5; semver pinned in Utility*PackageID constants.
	UtilityCredentialV0:      []string{CurrentVersion},
	UtilityRegistryV0:        []string{CurrentVersion},
	UtilityRegistryHoldingV0: []string{CurrentVersion},
	UtilityRegistryAppV0:     []string{CurrentVersion},
}

// VersionDir maps a DAR version string to its artifact subdirectory.
func VersionDir(version string) string {
	if version == CurrentVersion {
		return CurrentVersion
	}

	return ReleaseDir
}

func darPath(packageName Package, version string) string {
	return fmt.Sprintf("dars/%s/%s-%s.dar", VersionDir(version), packageName, version)
}

func GetDar(packageName Package, version string) ([]byte, error) {
	availableVersions, ok := Versions[packageName]
	if !ok {
		return nil, fmt.Errorf("no available versions for package %s", packageName)
	}

	if !slices.Contains(availableVersions, version) {
		return nil, fmt.Errorf("version %s not found for package %s", version, packageName)
	}

	path := darPath(packageName, version)
	data, err := Dars.ReadFile(path)
	if err != nil {
		// Try to read from Splice dependencies if not found in dars
		data, err = SpliceDependencies.ReadFile(fmt.Sprintf("dependencies/splice/%s-%s.dar", packageName, version))
		if err != nil {
			return nil, fmt.Errorf("failed to read embedded DAR file %s: %w", path, err)
		}
	}

	return data, nil
}

var OutputDirs = map[Package][]string{
	Coin: []string{"coin"},
	Link: []string{"link"},

	ChainlinkAPI: []string{"chainlink", "chainlinkapi"},
	MCMSAPI:      []string{"mcms", "api"},
	MCMSCore:     []string{"mcms", "core"},
	MCMSTest:     []string{"mcms", "mcmstest"},

	CCIPCore:                 []string{"ccip", "core"},
	CCIPExtensionAPI:         []string{"ccip", "extensionapi"},
	CCIPRuntime:              []string{"ccip", "ccipruntime"},
	CCIPReceiver:             []string{"ccip", "receiver"},
	CCIPSender:               []string{"ccip", "sender"},
	CCIPCommitteeVerifier:    []string{"ccip", "committeeverifier"},
	CCIPExecutor:             []string{"ccip", "executor"},
	CCIPLockReleaseTokenPool: []string{"ccip", "lockreleasetokenpool"},
	CCIPBurnMintTokenPool:    []string{"ccip", "burnminttokenpool"},
	CCIPFactory:              []string{"ccip", "factory"},

	SpliceApiFeaturedAppV1:                []string{"splice", "splice_api_featured_app_v1"},
	SpliceApiTokenAllocationV1:            []string{"splice", "splice_api_token_allocation_v1"},
	SpliceApiTokenAllocationInstructionV1: []string{"splice", "splice_api_token_allocation_instruction_v1"},
	SpliceApiTokenBurnMintV1:              []string{"splice", "splice_api_token_burn_mint_v1"},
	SpliceApiTokenHoldingV1:               []string{"splice", "splice_api_token_holding_v1"},
	SpliceApiTokenMetadataV1:              []string{"splice", "splice_api_token_metadata_v1"},
	SpliceApiTokenTransferInstructionV1:   []string{"splice", "splice_api_token_transfer_instruction_v1"},

	UtilityCredentialV0:      []string{"utility", "credential_v0"},
	UtilityRegistryV0:        []string{"utility", "registry_v0"},
	UtilityRegistryHoldingV0: []string{"utility", "registry_holding_v0"},
	UtilityRegistryAppV0:     []string{"utility", "registry_app_v0"},
}
