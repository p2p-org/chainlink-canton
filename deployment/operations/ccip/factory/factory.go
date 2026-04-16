package factory

import (
	"errors"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	factorybindings "github.com/smartcontractkit/chainlink-canton/bindings/generated/ccip/factory"
	"github.com/smartcontractkit/chainlink-canton/contracts"
	"github.com/smartcontractkit/chainlink-canton/deployment/utils/operations/contract"
)

var ContractType = deployment.ContractType("CCIPFactory")

var Version = semver.MustParse("0.1.0")

var factoryEncoder = factorybindings.NewContract("", "CCIP.Factory", "CCIPFactory").Encoder()

var Deploy = contract.NewDeploy(contract.DeployParams[factorybindings.CCIPFactory]{
	Name:           "canton/ccip/factory/deploy",
	TypeAndVersion: deployment.NewTypeAndVersion(ContractType, *Version),
	Description:    "Deploys a CCIPFactory contract on Canton",
	Validate: func(template factorybindings.CCIPFactory) error {
		if template.Owner == "" {
			return errors.New("owner cannot be empty")
		}
		if template.McmsParty == "" {
			return errors.New("mcmsParty cannot be empty")
		}

		return nil
	},
	PackageName: string(contracts.CCIPFactory),
	Prefix:      "factory",
})

var DeployRMNRemote = contract.NewExercise(contract.ExerciseParams[factorybindings.DeployRMNRemote]{
	Name:         "canton/ccip/factory/deploy_rmn_remote",
	Version:      Version,
	Description:  "Deploys an RMNRemote through the CCIPFactory",
	ContractType: ContractType,
	Template:     factorybindings.CCIPFactory{},
	Method:       factorybindings.CCIPFactory{}.DeployRMNRemote,
	EncodeMethod: factoryEncoder.DeployRMNRemote,
})

var DeployGlobalConfig = contract.NewExercise(contract.ExerciseParams[factorybindings.DeployGlobalConfig]{
	Name:         "canton/ccip/factory/deploy_global_config",
	Version:      Version,
	Description:  "Deploys a GlobalConfig through the CCIPFactory",
	ContractType: ContractType,
	Template:     factorybindings.CCIPFactory{},
	Method:       factorybindings.CCIPFactory{}.DeployGlobalConfig,
	EncodeMethod: factoryEncoder.DeployGlobalConfig,
})

var DeployTokenAdminRegistry = contract.NewExercise(contract.ExerciseParams[factorybindings.DeployTokenAdminRegistry]{
	Name:         "canton/ccip/factory/deploy_token_admin_registry",
	Version:      Version,
	Description:  "Deploys a TokenAdminRegistry through the CCIPFactory",
	ContractType: ContractType,
	Template:     factorybindings.CCIPFactory{},
	Method:       factorybindings.CCIPFactory{}.DeployTokenAdminRegistry,
	EncodeMethod: factoryEncoder.DeployTokenAdminRegistry,
})

var DeployFeeQuoter = contract.NewExercise(contract.ExerciseParams[factorybindings.DeployFeeQuoter]{
	Name:         "canton/ccip/factory/deploy_fee_quoter",
	Version:      Version,
	Description:  "Deploys a FeeQuoter through the CCIPFactory",
	ContractType: ContractType,
	Template:     factorybindings.CCIPFactory{},
	Method:       factorybindings.CCIPFactory{}.DeployFeeQuoter,
	EncodeMethod: factoryEncoder.DeployFeeQuoter,
})

var DeployCommitteeVerifier = contract.NewExercise(contract.ExerciseParams[factorybindings.DeployCommitteeVerifier]{
	Name:         "canton/ccip/factory/deploy_committee_verifier",
	Version:      Version,
	Description:  "Deploys a CommitteeVerifier through the CCIPFactory",
	ContractType: ContractType,
	Template:     factorybindings.CCIPFactory{},
	Method:       factorybindings.CCIPFactory{}.DeployCommitteeVerifier,
	EncodeMethod: factoryEncoder.DeployCommitteeVerifier,
})

var DeployOffRamp = contract.NewExercise(contract.ExerciseParams[factorybindings.DeployOffRamp]{
	Name:         "canton/ccip/factory/deploy_offramp",
	Version:      Version,
	Description:  "Deploys an OffRamp through the CCIPFactory",
	ContractType: ContractType,
	Template:     factorybindings.CCIPFactory{},
	Method:       factorybindings.CCIPFactory{}.DeployOffRamp,
	EncodeMethod: factoryEncoder.DeployOffRamp,
})

var DeployOnRamp = contract.NewExercise(contract.ExerciseParams[factorybindings.DeployOnRamp]{
	Name:         "canton/ccip/factory/deploy_onramp",
	Version:      Version,
	Description:  "Deploys an OnRamp through the CCIPFactory",
	ContractType: ContractType,
	Template:     factorybindings.CCIPFactory{},
	Method:       factorybindings.CCIPFactory{}.DeployOnRamp,
	EncodeMethod: factoryEncoder.DeployOnRamp,
})

var DeployPerPartyRouterFactory = contract.NewExercise(contract.ExerciseParams[factorybindings.DeployPerPartyRouterFactory]{
	Name:         "canton/ccip/factory/deploy_per_party_router_factory",
	Version:      Version,
	Description:  "Deploys a PerPartyRouterFactory through the CCIPFactory",
	ContractType: ContractType,
	Template:     factorybindings.CCIPFactory{},
	Method:       factorybindings.CCIPFactory{}.DeployPerPartyRouterFactory,
	EncodeMethod: factoryEncoder.DeployPerPartyRouterFactory,
})
