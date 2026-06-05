package commercials_v0

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/smartcontractkit/go-daml/pkg/bind"
	"github.com/smartcontractkit/go-daml/pkg/codec"
	"github.com/smartcontractkit/go-daml/pkg/model"
	"github.com/smartcontractkit/go-daml/pkg/types"
)

var (
	_ = fmt.Sprintf
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = model.Command{}
	_ bind.BoundTemplate
)

const (
	PackageName = "utility-commercials-v0"
	PackageID   = "fa5b1cc5c8368dff7c2e6a74aa2af9d520d755e2a508f44acd17343326e41839"
	SDKVersion  = "3.4.9"
)

type Template interface {
	CreateCommand() *model.CreateCommand
	GetTemplateID() string
}

func argsToMap(args any) map[string]any {
	if args == nil {
		return map[string]any{}
	}

	if m, ok := args.(map[string]any); ok {
		return m
	}

	type mapper interface {
		ToMap() map[string]any
	}
	if mapper, ok := args.(mapper); ok {
		return mapper.ToMap()
	}

	return map[string]any{"args": args}
}

// BillingContext is a Record type
type BillingContext struct {
	OpenRoundCid              types.CONTRACT_ID  `json:"openRoundCid"`
	OpenRound                 OpenMiningRound    `json:"openRound"`
	FeaturedTransferContext   AppTransferContext `json:"featuredTransferContext"`
	UnfeaturedTransferContext AppTransferContext `json:"unfeaturedTransferContext"`
}

// ToMap converts BillingContext to a map for DAML arguments
func (t BillingContext) ToMap() map[string]any {
	m := make(map[string]any)

	m["openRoundCid"] = model.NestedToDAMLValue(t.OpenRoundCid)

	m["openRound"] = model.NestedToDAMLValue(t.OpenRound)

	m["featuredTransferContext"] = model.NestedToDAMLValue(t.FeaturedTransferContext)

	m["unfeaturedTransferContext"] = model.NestedToDAMLValue(t.UnfeaturedTransferContext)

	return m
}

func (t BillingContext) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BillingContext) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BillingContext to hex string (Canton MCMS format)
func (t BillingContext) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BillingContext from hex string (Canton MCMS format)
func (t *BillingContext) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// BillingCycleParams is a Record type
type BillingCycleParams struct {
	AmuletPrice      types.NUMERIC   `json:"amuletPrice"`
	FeeAmountCc      types.NUMERIC   `json:"feeAmountCc"`
	NewBilledUntil   types.TIMESTAMP `json:"newBilledUntil"`
	DepositExpiresAt types.TIMESTAMP `json:"depositExpiresAt"`
}

// ToMap converts BillingCycleParams to a map for DAML arguments
func (t BillingCycleParams) ToMap() map[string]any {
	m := make(map[string]any)

	m["amuletPrice"] = t.AmuletPrice

	m["feeAmountCc"] = t.FeeAmountCc

	m["newBilledUntil"] = t.NewBilledUntil

	m["depositExpiresAt"] = t.DepositExpiresAt

	return m
}

func (t BillingCycleParams) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BillingCycleParams) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BillingCycleParams to hex string (Canton MCMS format)
func (t BillingCycleParams) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BillingCycleParams from hex string (Canton MCMS format)
func (t *BillingCycleParams) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// BillingState is a Record type
type BillingState struct {
	Status       BillingStatus   `json:"status"`
	LastBilledAt types.TIMESTAMP `json:"lastBilledAt"`
	BilledUntil  types.TIMESTAMP `json:"billedUntil"`
}

// ToMap converts BillingState to a map for DAML arguments
func (t BillingState) ToMap() map[string]any {
	m := make(map[string]any)

	m["status"] = model.NestedToDAMLValue(t.Status)

	m["lastBilledAt"] = t.LastBilledAt

	m["billedUntil"] = t.BilledUntil

	return m
}

func (t BillingState) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BillingState) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BillingState to hex string (Canton MCMS format)
func (t BillingState) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BillingState from hex string (Canton MCMS format)
func (t *BillingState) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// BillingStatus is a variant/union type
type BillingStatus struct {
	Success *types.UNIT `json:"Success,omitempty"`
	Failure *Failure    `json:"Failure,omitempty"`
	New     *types.UNIT `json:"New,omitempty"`
}

// MarshalJSON implements custom JSON marshaling for BillingStatus
func (v BillingStatus) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(v)
}

// UnmarshalJSON implements custom JSON unmarshalling for BillingStatus
func (v *BillingStatus) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, v)
}

// MarshalHex encodes BillingStatus to hex string (Canton MCMS format)
func (v BillingStatus) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(v)
}

// UnmarshalHex decodes BillingStatus from hex string (Canton MCMS format)
func (v *BillingStatus) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, v)
}

// GetVariantTag implements types.VARIANT interface
func (v BillingStatus) GetVariantTag() string {

	if v.Success != nil {
		return "Success"
	}

	if v.Failure != nil {
		return "Failure"
	}

	if v.New != nil {
		return "New"
	}

	return ""
}

// GetVariantValue implements types.VARIANT interface
func (v BillingStatus) GetVariantValue() any {

	if v.Success != nil {
		return v.Success
	}

	if v.Failure != nil {
		return v.Failure
	}

	if v.New != nil {
		return v.New
	}

	return nil
}

var _ types.VARIANT = (*BillingStatus)(nil)

// CommercialAgreement is a Template type
type CommercialAgreement struct {
	Operator                    types.PARTY         `json:"operator"`
	User                        types.PARTY         `json:"user"`
	FeeReceiver                 types.PARTY         `json:"feeReceiver"`
	LockedAmuletCids            []types.CONTRACT_ID `json:"lockedAmuletCids"`
	CurrentLockedAmuletAmountCc types.NUMERIC       `json:"currentLockedAmuletAmountCc"`
	UtilityFees                 UtilityFees         `json:"utilityFees"`
	Dso                         types.PARTY         `json:"dso"`
	BaseFeeBillingState         *BillingState       `json:"baseFeeBillingState" hex:"optional"`
	CredentialFeeBillingState   *EventBillingState  `json:"credentialFeeBillingState" hex:"optional"`
	AccruedFeesCc               *types.NUMERIC      `json:"accruedFeesCc" hex:"optional"`
	RewardReceiver              *types.PARTY        `json:"rewardReceiver" hex:"optional"`
	DataPublishingConsent       *types.BOOL         `json:"dataPublishingConsent" hex:"optional"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t CommercialAgreement) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t CommercialAgreement) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t CommercialAgreement) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["feeReceiver"] = t.FeeReceiver.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["lockedAmuletCids"] = func() []any {
		res := make([]any, 0, len(t.LockedAmuletCids))
		for _, e := range t.LockedAmuletCids {
			res = append(res, e)
		}
		return res
	}()

	if t.CurrentLockedAmuletAmountCc != "" {
		args["currentLockedAmuletAmountCc"] = t.CurrentLockedAmuletAmountCc
	}

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["utilityFees"] = model.NestedToDAMLValue(t.UtilityFees)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	if t.BaseFeeBillingState != nil {
		args["baseFeeBillingState"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.BaseFeeBillingState),
		}
	} else {
		args["baseFeeBillingState"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.CredentialFeeBillingState != nil {
		args["credentialFeeBillingState"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.CredentialFeeBillingState),
		}
	} else {
		args["credentialFeeBillingState"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.AccruedFeesCc != nil {
		args["accruedFeesCc"] = map[string]any{
			"_type": "optional",
			"value": *t.AccruedFeesCc,
		}
	} else {
		args["accruedFeesCc"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.RewardReceiver != nil {
		args["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": (*t.RewardReceiver).ToMap(),
		}
	} else {
		args["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.DataPublishingConsent != nil {
		args["dataPublishingConsent"] = map[string]any{
			"_type": "optional",
			"value": bool(*t.DataPublishingConsent),
		}
	} else {
		args["dataPublishingConsent"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t CommercialAgreement) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["feeReceiver"] = t.FeeReceiver.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["lockedAmuletCids"] = func() []any {
		res := make([]any, 0, len(t.LockedAmuletCids))
		for _, e := range t.LockedAmuletCids {
			res = append(res, e)
		}
		return res
	}()

	if t.CurrentLockedAmuletAmountCc != "" {
		args["currentLockedAmuletAmountCc"] = t.CurrentLockedAmuletAmountCc
	}

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["utilityFees"] = model.NestedToDAMLValue(t.UtilityFees)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	if t.BaseFeeBillingState != nil {
		args["baseFeeBillingState"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.BaseFeeBillingState),
		}
	} else {
		args["baseFeeBillingState"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.CredentialFeeBillingState != nil {
		args["credentialFeeBillingState"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.CredentialFeeBillingState),
		}
	} else {
		args["credentialFeeBillingState"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.AccruedFeesCc != nil {
		args["accruedFeesCc"] = map[string]any{
			"_type": "optional",
			"value": *t.AccruedFeesCc,
		}
	} else {
		args["accruedFeesCc"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.RewardReceiver != nil {
		args["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": (*t.RewardReceiver).ToMap(),
		}
	} else {
		args["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.DataPublishingConsent != nil {
		args["dataPublishingConsent"] = map[string]any{
			"_type": "optional",
			"value": bool(*t.DataPublishingConsent),
		}
	} else {
		args["dataPublishingConsent"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t CommercialAgreement) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreement) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreement to hex string (Canton MCMS format)
func (t CommercialAgreement) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreement from hex string (Canton MCMS format)
func (t *CommercialAgreement) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for CommercialAgreement

// CommercialAgreementRevoke exercises the CommercialAgreement_Revoke choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementRevoke(contractID string, args CommercialAgreementRevoke) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_Revoke",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementRevokeWithPackageID exercises the CommercialAgreement_Revoke choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementRevokeWithPackageID(contractID string, packageID string, args CommercialAgreementRevoke) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_Revoke",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementBillCredentialFeeMultiUnfeatured exercises the CommercialAgreement_BillCredentialFeeMultiUnfeatured choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementBillCredentialFeeMultiUnfeatured(contractID string, args CommercialAgreementBillCredentialFeeMultiUnfeatured) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_BillCredentialFeeMultiUnfeatured",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementBillCredentialFeeMultiUnfeaturedWithPackageID exercises the CommercialAgreement_BillCredentialFeeMultiUnfeatured choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementBillCredentialFeeMultiUnfeaturedWithPackageID(contractID string, packageID string, args CommercialAgreementBillCredentialFeeMultiUnfeatured) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_BillCredentialFeeMultiUnfeatured",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementBillCredentialFeeMulti exercises the CommercialAgreement_BillCredentialFeeMulti choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementBillCredentialFeeMulti(contractID string, args CommercialAgreementBillCredentialFeeMulti) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_BillCredentialFeeMulti",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementBillCredentialFeeMultiWithPackageID exercises the CommercialAgreement_BillCredentialFeeMulti choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementBillCredentialFeeMultiWithPackageID(contractID string, packageID string, args CommercialAgreementBillCredentialFeeMulti) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_BillCredentialFeeMulti",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementBillBaseFee exercises the CommercialAgreement_BillBaseFee choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementBillBaseFee(contractID string, args CommercialAgreementBillBaseFee) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_BillBaseFee",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementBillBaseFeeWithPackageID exercises the CommercialAgreement_BillBaseFee choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementBillBaseFeeWithPackageID(contractID string, packageID string, args CommercialAgreementBillBaseFee) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_BillBaseFee",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementLockCoin exercises the CommercialAgreement_LockCoin choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementLockCoin(contractID string, args CommercialAgreementLockCoin) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_LockCoin",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementLockCoinWithPackageID exercises the CommercialAgreement_LockCoin choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementLockCoinWithPackageID(contractID string, packageID string, args CommercialAgreementLockCoin) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_LockCoin",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementModify exercises the CommercialAgreement_Modify choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementModify(contractID string, args CommercialAgreementModify) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_Modify",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementModifyWithPackageID exercises the CommercialAgreement_Modify choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementModifyWithPackageID(contractID string, packageID string, args CommercialAgreementModify) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_Modify",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementFlushExpiredDeposit exercises the CommercialAgreement_FlushExpiredDeposit choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementFlushExpiredDeposit(contractID string, args CommercialAgreementFlushExpiredDeposit) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_FlushExpiredDeposit",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementFlushExpiredDepositWithPackageID exercises the CommercialAgreement_FlushExpiredDeposit choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementFlushExpiredDepositWithPackageID(contractID string, packageID string, args CommercialAgreementFlushExpiredDeposit) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_FlushExpiredDeposit",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementSetDefaultCredentialFeeBillingState exercises the CommercialAgreement_SetDefaultCredentialFeeBillingState choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementSetDefaultCredentialFeeBillingState(contractID string, args CommercialAgreementSetDefaultCredentialFeeBillingState) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_SetDefaultCredentialFeeBillingState",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementSetDefaultCredentialFeeBillingStateWithPackageID exercises the CommercialAgreement_SetDefaultCredentialFeeBillingState choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementSetDefaultCredentialFeeBillingStateWithPackageID(contractID string, packageID string, args CommercialAgreementSetDefaultCredentialFeeBillingState) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_SetDefaultCredentialFeeBillingState",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t CommercialAgreement) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// CommercialAgreementModifyDataPublishingConsent exercises the CommercialAgreement_ModifyDataPublishingConsent choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementModifyDataPublishingConsent(contractID string, args CommercialAgreementModifyDataPublishingConsent) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_ModifyDataPublishingConsent",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementModifyDataPublishingConsentWithPackageID exercises the CommercialAgreement_ModifyDataPublishingConsent choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementModifyDataPublishingConsentWithPackageID(contractID string, packageID string, args CommercialAgreementModifyDataPublishingConsent) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_ModifyDataPublishingConsent",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementBill exercises the CommercialAgreement_Bill choice on this CommercialAgreement contract
// This method uses the package name in the template ID
func (t CommercialAgreement) CommercialAgreementBill(contractID string, args CommercialAgreementBill) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_Bill",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementBillWithPackageID exercises the CommercialAgreement_Bill choice using the provided package ID instead of package name
func (t CommercialAgreement) CommercialAgreementBillWithPackageID(contractID string, packageID string, args CommercialAgreementBill) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.CommercialAgreement", "CommercialAgreement"),
		ContractID: contractID,
		Choice:     "CommercialAgreement_Bill",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementOffer is a Template type
type CommercialAgreementOffer struct {
	Operator               types.PARTY    `json:"operator"`
	User                   types.PARTY    `json:"user"`
	FeeReceiver            types.PARTY    `json:"feeReceiver"`
	UtilityFees            UtilityFees    `json:"utilityFees"`
	Dso                    types.PARTY    `json:"dso"`
	InitialDepositAmountCc *types.NUMERIC `json:"initialDepositAmountCc" hex:"optional"`
	RewardReceiver         *types.PARTY   `json:"rewardReceiver" hex:"optional"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t CommercialAgreementOffer) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t CommercialAgreementOffer) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t CommercialAgreementOffer) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["feeReceiver"] = t.FeeReceiver.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["utilityFees"] = model.NestedToDAMLValue(t.UtilityFees)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	if t.InitialDepositAmountCc != nil {
		args["initialDepositAmountCc"] = map[string]any{
			"_type": "optional",
			"value": *t.InitialDepositAmountCc,
		}
	} else {
		args["initialDepositAmountCc"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.RewardReceiver != nil {
		args["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": (*t.RewardReceiver).ToMap(),
		}
	} else {
		args["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t CommercialAgreementOffer) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["feeReceiver"] = t.FeeReceiver.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["utilityFees"] = model.NestedToDAMLValue(t.UtilityFees)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	if t.InitialDepositAmountCc != nil {
		args["initialDepositAmountCc"] = map[string]any{
			"_type": "optional",
			"value": *t.InitialDepositAmountCc,
		}
	} else {
		args["initialDepositAmountCc"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.RewardReceiver != nil {
		args["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": (*t.RewardReceiver).ToMap(),
		}
	} else {
		args["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t CommercialAgreementOffer) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOffer) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOffer to hex string (Canton MCMS format)
func (t CommercialAgreementOffer) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOffer from hex string (Canton MCMS format)
func (t *CommercialAgreementOffer) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for CommercialAgreementOffer

// CommercialAgreementOfferAccept exercises the CommercialAgreementOffer_Accept choice on this CommercialAgreementOffer contract
// This method uses the package name in the template ID
func (t CommercialAgreementOffer) CommercialAgreementOfferAccept(contractID string, args CommercialAgreementOfferAccept) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "CommercialAgreementOffer_Accept",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementOfferAcceptWithPackageID exercises the CommercialAgreementOffer_Accept choice using the provided package ID instead of package name
func (t CommercialAgreementOffer) CommercialAgreementOfferAcceptWithPackageID(contractID string, packageID string, args CommercialAgreementOfferAccept) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "CommercialAgreementOffer_Accept",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementOfferAcceptAndTopup exercises the CommercialAgreementOffer_AcceptAndTopup choice on this CommercialAgreementOffer contract
// This method uses the package name in the template ID
func (t CommercialAgreementOffer) CommercialAgreementOfferAcceptAndTopup(contractID string, args CommercialAgreementOfferAcceptAndTopup) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "CommercialAgreementOffer_AcceptAndTopup",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementOfferAcceptAndTopupWithPackageID exercises the CommercialAgreementOffer_AcceptAndTopup choice using the provided package ID instead of package name
func (t CommercialAgreementOffer) CommercialAgreementOfferAcceptAndTopupWithPackageID(contractID string, packageID string, args CommercialAgreementOfferAcceptAndTopup) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "CommercialAgreementOffer_AcceptAndTopup",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementOfferCancel exercises the CommercialAgreementOffer_Cancel choice on this CommercialAgreementOffer contract
// This method uses the package name in the template ID
func (t CommercialAgreementOffer) CommercialAgreementOfferCancel(contractID string, args CommercialAgreementOfferCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "CommercialAgreementOffer_Cancel",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementOfferCancelWithPackageID exercises the CommercialAgreementOffer_Cancel choice using the provided package ID instead of package name
func (t CommercialAgreementOffer) CommercialAgreementOfferCancelWithPackageID(contractID string, packageID string, args CommercialAgreementOfferCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "CommercialAgreementOffer_Cancel",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this CommercialAgreementOffer contract
// This method uses the package name in the template ID
func (t CommercialAgreementOffer) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t CommercialAgreementOffer) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// CommercialAgreementOfferReject exercises the CommercialAgreementOffer_Reject choice on this CommercialAgreementOffer contract
// This method uses the package name in the template ID
func (t CommercialAgreementOffer) CommercialAgreementOfferReject(contractID string, args CommercialAgreementOfferReject) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "CommercialAgreementOffer_Reject",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementOfferRejectWithPackageID exercises the CommercialAgreementOffer_Reject choice using the provided package ID instead of package name
func (t CommercialAgreementOffer) CommercialAgreementOfferRejectWithPackageID(contractID string, packageID string, args CommercialAgreementOfferReject) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.Offer", "CommercialAgreementOffer"),
		ContractID: contractID,
		Choice:     "CommercialAgreementOffer_Reject",
		Arguments:  argsToMap(args),
	}
}

// CommercialAgreementOfferAccept is a Record type
type CommercialAgreementOfferAccept struct {
	DataPublishingConsent *types.BOOL `json:"dataPublishingConsent" hex:"optional"`
}

// ToMap converts CommercialAgreementOfferAccept to a map for DAML arguments
func (t CommercialAgreementOfferAccept) ToMap() map[string]any {
	m := make(map[string]any)

	if t.DataPublishingConsent != nil {
		m["dataPublishingConsent"] = map[string]any{
			"_type": "optional",
			"value": bool(*t.DataPublishingConsent),
		}
	} else {
		m["dataPublishingConsent"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CommercialAgreementOfferAccept) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOfferAccept) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOfferAccept to hex string (Canton MCMS format)
func (t CommercialAgreementOfferAccept) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOfferAccept from hex string (Canton MCMS format)
func (t *CommercialAgreementOfferAccept) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementOfferAcceptAndTopup is a Record type
type CommercialAgreementOfferAcceptAndTopup struct {
	HolderInputs          []types.CONTRACT_ID `json:"holderInputs"`
	AppTransferContext    AppTransferContext  `json:"appTransferContext"`
	DataPublishingConsent *types.BOOL         `json:"dataPublishingConsent" hex:"optional"`
}

// ToMap converts CommercialAgreementOfferAcceptAndTopup to a map for DAML arguments
func (t CommercialAgreementOfferAcceptAndTopup) ToMap() map[string]any {
	m := make(map[string]any)

	m["holderInputs"] = func() []any {
		res := make([]any, 0, len(t.HolderInputs))
		for _, e := range t.HolderInputs {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	if t.DataPublishingConsent != nil {
		m["dataPublishingConsent"] = map[string]any{
			"_type": "optional",
			"value": bool(*t.DataPublishingConsent),
		}
	} else {
		m["dataPublishingConsent"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CommercialAgreementOfferAcceptAndTopup) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOfferAcceptAndTopup) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOfferAcceptAndTopup to hex string (Canton MCMS format)
func (t CommercialAgreementOfferAcceptAndTopup) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOfferAcceptAndTopup from hex string (Canton MCMS format)
func (t *CommercialAgreementOfferAcceptAndTopup) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementOfferAcceptAndTopupResult is a Record type
type CommercialAgreementOfferAcceptAndTopupResult struct {
	CommercialAgreementCid types.CONTRACT_ID `json:"commercialAgreementCid"`
}

// ToMap converts CommercialAgreementOfferAcceptAndTopupResult to a map for DAML arguments
func (t CommercialAgreementOfferAcceptAndTopupResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	return m
}

func (t CommercialAgreementOfferAcceptAndTopupResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOfferAcceptAndTopupResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOfferAcceptAndTopupResult to hex string (Canton MCMS format)
func (t CommercialAgreementOfferAcceptAndTopupResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOfferAcceptAndTopupResult from hex string (Canton MCMS format)
func (t *CommercialAgreementOfferAcceptAndTopupResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementOfferAcceptResult is a Record type
type CommercialAgreementOfferAcceptResult struct {
	CommercialAgreementCid types.CONTRACT_ID `json:"commercialAgreementCid"`
}

// ToMap converts CommercialAgreementOfferAcceptResult to a map for DAML arguments
func (t CommercialAgreementOfferAcceptResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	return m
}

func (t CommercialAgreementOfferAcceptResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOfferAcceptResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOfferAcceptResult to hex string (Canton MCMS format)
func (t CommercialAgreementOfferAcceptResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOfferAcceptResult from hex string (Canton MCMS format)
func (t *CommercialAgreementOfferAcceptResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementOfferCancel is a Record type
type CommercialAgreementOfferCancel struct {
}

// ToMap converts CommercialAgreementOfferCancel to a map for DAML arguments
func (t CommercialAgreementOfferCancel) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t CommercialAgreementOfferCancel) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOfferCancel) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOfferCancel to hex string (Canton MCMS format)
func (t CommercialAgreementOfferCancel) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOfferCancel from hex string (Canton MCMS format)
func (t *CommercialAgreementOfferCancel) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementOfferCancelResult is a Record type
type CommercialAgreementOfferCancelResult struct {
}

// ToMap converts CommercialAgreementOfferCancelResult to a map for DAML arguments
func (t CommercialAgreementOfferCancelResult) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t CommercialAgreementOfferCancelResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOfferCancelResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOfferCancelResult to hex string (Canton MCMS format)
func (t CommercialAgreementOfferCancelResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOfferCancelResult from hex string (Canton MCMS format)
func (t *CommercialAgreementOfferCancelResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementOfferReject is a Record type
type CommercialAgreementOfferReject struct {
	Reason types.TEXT `json:"reason"`
}

// ToMap converts CommercialAgreementOfferReject to a map for DAML arguments
func (t CommercialAgreementOfferReject) ToMap() map[string]any {
	m := make(map[string]any)

	m["reason"] = string(t.Reason)

	return m
}

func (t CommercialAgreementOfferReject) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOfferReject) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOfferReject to hex string (Canton MCMS format)
func (t CommercialAgreementOfferReject) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOfferReject from hex string (Canton MCMS format)
func (t *CommercialAgreementOfferReject) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementOfferRejectResult is a Record type
type CommercialAgreementOfferRejectResult struct {
	Reason types.TEXT `json:"reason"`
}

// ToMap converts CommercialAgreementOfferRejectResult to a map for DAML arguments
func (t CommercialAgreementOfferRejectResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["reason"] = string(t.Reason)

	return m
}

func (t CommercialAgreementOfferRejectResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementOfferRejectResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementOfferRejectResult to hex string (Canton MCMS format)
func (t CommercialAgreementOfferRejectResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementOfferRejectResult from hex string (Canton MCMS format)
func (t *CommercialAgreementOfferRejectResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementBill is a Record type
type CommercialAgreementBill struct {
	TransferContext        AppTransferContext     `json:"transferContext"`
	TransferPreapprovalCid types.CONTRACT_ID      `json:"transferPreapprovalCid"`
	PaymentTransferContext PaymentTransferContext `json:"paymentTransferContext"`
}

// ToMap converts CommercialAgreementBill to a map for DAML arguments
func (t CommercialAgreementBill) ToMap() map[string]any {
	m := make(map[string]any)

	m["transferContext"] = model.NestedToDAMLValue(t.TransferContext)

	m["transferPreapprovalCid"] = model.NestedToDAMLValue(t.TransferPreapprovalCid)

	m["paymentTransferContext"] = model.NestedToDAMLValue(t.PaymentTransferContext)

	return m
}

func (t CommercialAgreementBill) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementBill) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementBill to hex string (Canton MCMS format)
func (t CommercialAgreementBill) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementBill from hex string (Canton MCMS format)
func (t *CommercialAgreementBill) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementBillBaseFee is a Record type
type CommercialAgreementBillBaseFee struct {
	TransferContext        AppTransferContext     `json:"transferContext"`
	TransferPreapprovalCid types.CONTRACT_ID      `json:"transferPreapprovalCid"`
	PaymentTransferContext PaymentTransferContext `json:"paymentTransferContext"`
}

// ToMap converts CommercialAgreementBillBaseFee to a map for DAML arguments
func (t CommercialAgreementBillBaseFee) ToMap() map[string]any {
	m := make(map[string]any)

	m["transferContext"] = model.NestedToDAMLValue(t.TransferContext)

	m["transferPreapprovalCid"] = model.NestedToDAMLValue(t.TransferPreapprovalCid)

	m["paymentTransferContext"] = model.NestedToDAMLValue(t.PaymentTransferContext)

	return m
}

func (t CommercialAgreementBillBaseFee) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementBillBaseFee) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementBillBaseFee to hex string (Canton MCMS format)
func (t CommercialAgreementBillBaseFee) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementBillBaseFee from hex string (Canton MCMS format)
func (t *CommercialAgreementBillBaseFee) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementBillBaseFeeResult is a Record type
type CommercialAgreementBillBaseFeeResult struct {
	CommercialAgreementCid types.CONTRACT_ID `json:"commercialAgreementCid"`
}

// ToMap converts CommercialAgreementBillBaseFeeResult to a map for DAML arguments
func (t CommercialAgreementBillBaseFeeResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	return m
}

func (t CommercialAgreementBillBaseFeeResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementBillBaseFeeResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementBillBaseFeeResult to hex string (Canton MCMS format)
func (t CommercialAgreementBillBaseFeeResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementBillBaseFeeResult from hex string (Canton MCMS format)
func (t *CommercialAgreementBillBaseFeeResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementBillCredentialFeeMulti is a Record type
type CommercialAgreementBillCredentialFeeMulti struct {
	TransferContext        AppTransferContext     `json:"transferContext"`
	TransferPreapprovalCid types.CONTRACT_ID      `json:"transferPreapprovalCid"`
	PaymentTransferContext PaymentTransferContext `json:"paymentTransferContext"`
	NumberOfBillings       types.INT64            `json:"numberOfBillings"`
	CurrentLedgerOffset    types.INT64            `json:"currentLedgerOffset"`
	CurrentMigrationId     *types.TEXT            `json:"currentMigrationId" hex:"optional"`
}

// ToMap converts CommercialAgreementBillCredentialFeeMulti to a map for DAML arguments
func (t CommercialAgreementBillCredentialFeeMulti) ToMap() map[string]any {
	m := make(map[string]any)

	m["transferContext"] = model.NestedToDAMLValue(t.TransferContext)

	m["transferPreapprovalCid"] = model.NestedToDAMLValue(t.TransferPreapprovalCid)

	m["paymentTransferContext"] = model.NestedToDAMLValue(t.PaymentTransferContext)

	m["numberOfBillings"] = int64(t.NumberOfBillings)

	m["currentLedgerOffset"] = int64(t.CurrentLedgerOffset)

	if t.CurrentMigrationId != nil {
		m["currentMigrationId"] = map[string]any{
			"_type": "optional",
			"value": string(*t.CurrentMigrationId),
		}
	} else {
		m["currentMigrationId"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CommercialAgreementBillCredentialFeeMulti) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementBillCredentialFeeMulti) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementBillCredentialFeeMulti to hex string (Canton MCMS format)
func (t CommercialAgreementBillCredentialFeeMulti) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementBillCredentialFeeMulti from hex string (Canton MCMS format)
func (t *CommercialAgreementBillCredentialFeeMulti) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementBillCredentialFeeMultiUnfeatured is a Record type
type CommercialAgreementBillCredentialFeeMultiUnfeatured struct {
	TransferContext        AppTransferContext `json:"transferContext"`
	TransferPreapprovalCid types.CONTRACT_ID  `json:"transferPreapprovalCid"`
	NumberOfBillings       types.INT64        `json:"numberOfBillings"`
	CurrentLedgerOffset    types.INT64        `json:"currentLedgerOffset"`
	PayoutThresholdCc      *types.NUMERIC     `json:"payoutThresholdCc" hex:"optional"`
	CurrentMigrationId     *types.TEXT        `json:"currentMigrationId" hex:"optional"`
}

// ToMap converts CommercialAgreementBillCredentialFeeMultiUnfeatured to a map for DAML arguments
func (t CommercialAgreementBillCredentialFeeMultiUnfeatured) ToMap() map[string]any {
	m := make(map[string]any)

	m["transferContext"] = model.NestedToDAMLValue(t.TransferContext)

	m["transferPreapprovalCid"] = model.NestedToDAMLValue(t.TransferPreapprovalCid)

	m["numberOfBillings"] = int64(t.NumberOfBillings)

	m["currentLedgerOffset"] = int64(t.CurrentLedgerOffset)

	if t.PayoutThresholdCc != nil {
		m["payoutThresholdCc"] = map[string]any{
			"_type": "optional",
			"value": *t.PayoutThresholdCc,
		}
	} else {
		m["payoutThresholdCc"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.CurrentMigrationId != nil {
		m["currentMigrationId"] = map[string]any{
			"_type": "optional",
			"value": string(*t.CurrentMigrationId),
		}
	} else {
		m["currentMigrationId"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CommercialAgreementBillCredentialFeeMultiUnfeatured) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementBillCredentialFeeMultiUnfeatured) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementBillCredentialFeeMultiUnfeatured to hex string (Canton MCMS format)
func (t CommercialAgreementBillCredentialFeeMultiUnfeatured) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementBillCredentialFeeMultiUnfeatured from hex string (Canton MCMS format)
func (t *CommercialAgreementBillCredentialFeeMultiUnfeatured) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementBillCredentialFeeMultiUnfeaturedResult is a Record type
type CommercialAgreementBillCredentialFeeMultiUnfeaturedResult struct {
	CommercialAgreementCid    types.CONTRACT_ID `json:"commercialAgreementCid"`
	CredentialFeeBillingState EventBillingState `json:"credentialFeeBillingState"`
}

// ToMap converts CommercialAgreementBillCredentialFeeMultiUnfeaturedResult to a map for DAML arguments
func (t CommercialAgreementBillCredentialFeeMultiUnfeaturedResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	m["credentialFeeBillingState"] = model.NestedToDAMLValue(t.CredentialFeeBillingState)

	return m
}

func (t CommercialAgreementBillCredentialFeeMultiUnfeaturedResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementBillCredentialFeeMultiUnfeaturedResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementBillCredentialFeeMultiUnfeaturedResult to hex string (Canton MCMS format)
func (t CommercialAgreementBillCredentialFeeMultiUnfeaturedResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementBillCredentialFeeMultiUnfeaturedResult from hex string (Canton MCMS format)
func (t *CommercialAgreementBillCredentialFeeMultiUnfeaturedResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementBillCredentialFeeMultiResult is a Record type
type CommercialAgreementBillCredentialFeeMultiResult struct {
	CommercialAgreementCid    types.CONTRACT_ID `json:"commercialAgreementCid"`
	CredentialFeeBillingState EventBillingState `json:"credentialFeeBillingState"`
}

// ToMap converts CommercialAgreementBillCredentialFeeMultiResult to a map for DAML arguments
func (t CommercialAgreementBillCredentialFeeMultiResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	m["credentialFeeBillingState"] = model.NestedToDAMLValue(t.CredentialFeeBillingState)

	return m
}

func (t CommercialAgreementBillCredentialFeeMultiResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementBillCredentialFeeMultiResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementBillCredentialFeeMultiResult to hex string (Canton MCMS format)
func (t CommercialAgreementBillCredentialFeeMultiResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementBillCredentialFeeMultiResult from hex string (Canton MCMS format)
func (t *CommercialAgreementBillCredentialFeeMultiResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementBillResult is a Record type
type CommercialAgreementBillResult struct {
	CommercialAgreementCid types.CONTRACT_ID `json:"commercialAgreementCid"`
}

// ToMap converts CommercialAgreementBillResult to a map for DAML arguments
func (t CommercialAgreementBillResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	return m
}

func (t CommercialAgreementBillResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementBillResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementBillResult to hex string (Canton MCMS format)
func (t CommercialAgreementBillResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementBillResult from hex string (Canton MCMS format)
func (t *CommercialAgreementBillResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementFlushExpiredDeposit is a Record type
type CommercialAgreementFlushExpiredDeposit struct {
	Actor types.PARTY `json:"actor"`
}

// ToMap converts CommercialAgreementFlushExpiredDeposit to a map for DAML arguments
func (t CommercialAgreementFlushExpiredDeposit) ToMap() map[string]any {
	m := make(map[string]any)

	m["actor"] = t.Actor.ToMap()

	return m
}

func (t CommercialAgreementFlushExpiredDeposit) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementFlushExpiredDeposit) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementFlushExpiredDeposit to hex string (Canton MCMS format)
func (t CommercialAgreementFlushExpiredDeposit) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementFlushExpiredDeposit from hex string (Canton MCMS format)
func (t *CommercialAgreementFlushExpiredDeposit) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementFlushExpiredDepositResult is a Record type
type CommercialAgreementFlushExpiredDepositResult struct {
	CommercialAgreementCid types.CONTRACT_ID `json:"commercialAgreementCid"`
}

// ToMap converts CommercialAgreementFlushExpiredDepositResult to a map for DAML arguments
func (t CommercialAgreementFlushExpiredDepositResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	return m
}

func (t CommercialAgreementFlushExpiredDepositResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementFlushExpiredDepositResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementFlushExpiredDepositResult to hex string (Canton MCMS format)
func (t CommercialAgreementFlushExpiredDepositResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementFlushExpiredDepositResult from hex string (Canton MCMS format)
func (t *CommercialAgreementFlushExpiredDepositResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementLockCoin is a Record type
type CommercialAgreementLockCoin struct {
	TargetAmount    types.NUMERIC       `json:"targetAmount"`
	CoinCids        []types.CONTRACT_ID `json:"coinCids"`
	TransferContext AppTransferContext  `json:"transferContext"`
}

// ToMap converts CommercialAgreementLockCoin to a map for DAML arguments
func (t CommercialAgreementLockCoin) ToMap() map[string]any {
	m := make(map[string]any)

	m["targetAmount"] = t.TargetAmount

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["transferContext"] = model.NestedToDAMLValue(t.TransferContext)

	return m
}

func (t CommercialAgreementLockCoin) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementLockCoin) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementLockCoin to hex string (Canton MCMS format)
func (t CommercialAgreementLockCoin) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementLockCoin from hex string (Canton MCMS format)
func (t *CommercialAgreementLockCoin) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementLockCoinResult is a Record type
type CommercialAgreementLockCoinResult struct {
	CommercialAgreementCid types.CONTRACT_ID `json:"commercialAgreementCid"`
}

// ToMap converts CommercialAgreementLockCoinResult to a map for DAML arguments
func (t CommercialAgreementLockCoinResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	return m
}

func (t CommercialAgreementLockCoinResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementLockCoinResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementLockCoinResult to hex string (Canton MCMS format)
func (t CommercialAgreementLockCoinResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementLockCoinResult from hex string (Canton MCMS format)
func (t *CommercialAgreementLockCoinResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementModify is a Record type
type CommercialAgreementModify struct {
	FeeReceiver    types.PARTY  `json:"feeReceiver"`
	UtilityFees    UtilityFees  `json:"utilityFees"`
	RewardReceiver *types.PARTY `json:"rewardReceiver" hex:"optional"`
}

// ToMap converts CommercialAgreementModify to a map for DAML arguments
func (t CommercialAgreementModify) ToMap() map[string]any {
	m := make(map[string]any)

	m["feeReceiver"] = t.FeeReceiver.ToMap()

	m["utilityFees"] = model.NestedToDAMLValue(t.UtilityFees)

	if t.RewardReceiver != nil {
		m["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": (*t.RewardReceiver).ToMap(),
		}
	} else {
		m["rewardReceiver"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CommercialAgreementModify) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementModify) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementModify to hex string (Canton MCMS format)
func (t CommercialAgreementModify) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementModify from hex string (Canton MCMS format)
func (t *CommercialAgreementModify) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementModifyDataPublishingConsent is a Record type
type CommercialAgreementModifyDataPublishingConsent struct {
	DataPublishingConsent types.BOOL `json:"dataPublishingConsent"`
}

// ToMap converts CommercialAgreementModifyDataPublishingConsent to a map for DAML arguments
func (t CommercialAgreementModifyDataPublishingConsent) ToMap() map[string]any {
	m := make(map[string]any)

	m["dataPublishingConsent"] = bool(t.DataPublishingConsent)

	return m
}

func (t CommercialAgreementModifyDataPublishingConsent) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementModifyDataPublishingConsent) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementModifyDataPublishingConsent to hex string (Canton MCMS format)
func (t CommercialAgreementModifyDataPublishingConsent) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementModifyDataPublishingConsent from hex string (Canton MCMS format)
func (t *CommercialAgreementModifyDataPublishingConsent) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementModifyDataPublishingConsentResult is a Record type
type CommercialAgreementModifyDataPublishingConsentResult struct {
	CommercialAgreementCid types.CONTRACT_ID `json:"commercialAgreementCid"`
}

// ToMap converts CommercialAgreementModifyDataPublishingConsentResult to a map for DAML arguments
func (t CommercialAgreementModifyDataPublishingConsentResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	return m
}

func (t CommercialAgreementModifyDataPublishingConsentResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementModifyDataPublishingConsentResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementModifyDataPublishingConsentResult to hex string (Canton MCMS format)
func (t CommercialAgreementModifyDataPublishingConsentResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementModifyDataPublishingConsentResult from hex string (Canton MCMS format)
func (t *CommercialAgreementModifyDataPublishingConsentResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementModifyResult is a Record type
type CommercialAgreementModifyResult struct {
	CommercialAgreementCid types.CONTRACT_ID `json:"commercialAgreementCid"`
}

// ToMap converts CommercialAgreementModifyResult to a map for DAML arguments
func (t CommercialAgreementModifyResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	return m
}

func (t CommercialAgreementModifyResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementModifyResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementModifyResult to hex string (Canton MCMS format)
func (t CommercialAgreementModifyResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementModifyResult from hex string (Canton MCMS format)
func (t *CommercialAgreementModifyResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementRevoke is a Record type
type CommercialAgreementRevoke struct {
	TransferContext        AppTransferContext `json:"transferContext"`
	Actor                  types.PARTY        `json:"actor"`
	TransferPreapprovalCid *types.CONTRACT_ID `json:"transferPreapprovalCid" hex:"optional"`
}

// ToMap converts CommercialAgreementRevoke to a map for DAML arguments
func (t CommercialAgreementRevoke) ToMap() map[string]any {
	m := make(map[string]any)

	m["transferContext"] = model.NestedToDAMLValue(t.TransferContext)

	m["actor"] = t.Actor.ToMap()

	if t.TransferPreapprovalCid != nil {
		m["transferPreapprovalCid"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.TransferPreapprovalCid),
		}
	} else {
		m["transferPreapprovalCid"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CommercialAgreementRevoke) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementRevoke) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementRevoke to hex string (Canton MCMS format)
func (t CommercialAgreementRevoke) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementRevoke from hex string (Canton MCMS format)
func (t *CommercialAgreementRevoke) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementRevokeResult is a Record type
type CommercialAgreementRevokeResult struct {
	UnlockedDeposit []types.CONTRACT_ID `json:"unlockedDeposit"`
}

// ToMap converts CommercialAgreementRevokeResult to a map for DAML arguments
func (t CommercialAgreementRevokeResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["unlockedDeposit"] = func() []any {
		res := make([]any, 0, len(t.UnlockedDeposit))
		for _, e := range t.UnlockedDeposit {
			res = append(res, e)
		}
		return res
	}()

	return m
}

func (t CommercialAgreementRevokeResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementRevokeResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementRevokeResult to hex string (Canton MCMS format)
func (t CommercialAgreementRevokeResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementRevokeResult from hex string (Canton MCMS format)
func (t *CommercialAgreementRevokeResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementSetDefaultCredentialFeeBillingState is a Record type
type CommercialAgreementSetDefaultCredentialFeeBillingState struct {
	CurrentLedgerOffset types.INT64 `json:"currentLedgerOffset"`
	CurrentMigrationId  *types.TEXT `json:"currentMigrationId" hex:"optional"`
}

// ToMap converts CommercialAgreementSetDefaultCredentialFeeBillingState to a map for DAML arguments
func (t CommercialAgreementSetDefaultCredentialFeeBillingState) ToMap() map[string]any {
	m := make(map[string]any)

	m["currentLedgerOffset"] = int64(t.CurrentLedgerOffset)

	if t.CurrentMigrationId != nil {
		m["currentMigrationId"] = map[string]any{
			"_type": "optional",
			"value": string(*t.CurrentMigrationId),
		}
	} else {
		m["currentMigrationId"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CommercialAgreementSetDefaultCredentialFeeBillingState) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementSetDefaultCredentialFeeBillingState) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementSetDefaultCredentialFeeBillingState to hex string (Canton MCMS format)
func (t CommercialAgreementSetDefaultCredentialFeeBillingState) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementSetDefaultCredentialFeeBillingState from hex string (Canton MCMS format)
func (t *CommercialAgreementSetDefaultCredentialFeeBillingState) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CommercialAgreementSetDefaultCredentialFeeBillingStateResult is a Record type
type CommercialAgreementSetDefaultCredentialFeeBillingStateResult struct {
	CommercialAgreementCid    types.CONTRACT_ID `json:"commercialAgreementCid"`
	CredentialFeeBillingState EventBillingState `json:"credentialFeeBillingState"`
}

// ToMap converts CommercialAgreementSetDefaultCredentialFeeBillingStateResult to a map for DAML arguments
func (t CommercialAgreementSetDefaultCredentialFeeBillingStateResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["commercialAgreementCid"] = model.NestedToDAMLValue(t.CommercialAgreementCid)

	m["credentialFeeBillingState"] = model.NestedToDAMLValue(t.CredentialFeeBillingState)

	return m
}

func (t CommercialAgreementSetDefaultCredentialFeeBillingStateResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CommercialAgreementSetDefaultCredentialFeeBillingStateResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CommercialAgreementSetDefaultCredentialFeeBillingStateResult to hex string (Canton MCMS format)
func (t CommercialAgreementSetDefaultCredentialFeeBillingStateResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CommercialAgreementSetDefaultCredentialFeeBillingStateResult from hex string (Canton MCMS format)
func (t *CommercialAgreementSetDefaultCredentialFeeBillingStateResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// DelegatedBatchedMarkersProxy is a Template type
type DelegatedBatchedMarkersProxy struct {
	Operator types.PARTY `json:"operator"`
	Provider types.PARTY `json:"provider"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t DelegatedBatchedMarkersProxy) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.DelegatedBatchedMarkersProxy", "DelegatedBatchedMarkersProxy")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t DelegatedBatchedMarkersProxy) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Commercials.V0.Model.DelegatedBatchedMarkersProxy", "DelegatedBatchedMarkersProxy")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t DelegatedBatchedMarkersProxy) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["provider"] = t.Provider.ToMap()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t DelegatedBatchedMarkersProxy) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["provider"] = t.Provider.ToMap()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t DelegatedBatchedMarkersProxy) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *DelegatedBatchedMarkersProxy) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes DelegatedBatchedMarkersProxy to hex string (Canton MCMS format)
func (t DelegatedBatchedMarkersProxy) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes DelegatedBatchedMarkersProxy from hex string (Canton MCMS format)
func (t *DelegatedBatchedMarkersProxy) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for DelegatedBatchedMarkersProxy

// DelegatedBatchedMarkersProxyCreateMarkers exercises the DelegatedBatchedMarkersProxy_CreateMarkers choice on this DelegatedBatchedMarkersProxy contract
// This method uses the package name in the template ID
func (t DelegatedBatchedMarkersProxy) DelegatedBatchedMarkersProxyCreateMarkers(contractID string, args DelegatedBatchedMarkersProxyCreateMarkers) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.DelegatedBatchedMarkersProxy", "DelegatedBatchedMarkersProxy"),
		ContractID: contractID,
		Choice:     "DelegatedBatchedMarkersProxy_CreateMarkers",
		Arguments:  argsToMap(args),
	}
}

// DelegatedBatchedMarkersProxyCreateMarkersWithPackageID exercises the DelegatedBatchedMarkersProxy_CreateMarkers choice using the provided package ID instead of package name
func (t DelegatedBatchedMarkersProxy) DelegatedBatchedMarkersProxyCreateMarkersWithPackageID(contractID string, packageID string, args DelegatedBatchedMarkersProxyCreateMarkers) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.DelegatedBatchedMarkersProxy", "DelegatedBatchedMarkersProxy"),
		ContractID: contractID,
		Choice:     "DelegatedBatchedMarkersProxy_CreateMarkers",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this DelegatedBatchedMarkersProxy contract
// This method uses the package name in the template ID
func (t DelegatedBatchedMarkersProxy) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.DelegatedBatchedMarkersProxy", "DelegatedBatchedMarkersProxy"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t DelegatedBatchedMarkersProxy) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.DelegatedBatchedMarkersProxy", "DelegatedBatchedMarkersProxy"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// DelegatedBatchedMarkersProxyArchive exercises the DelegatedBatchedMarkersProxy_Archive choice on this DelegatedBatchedMarkersProxy contract
// This method uses the package name in the template ID
func (t DelegatedBatchedMarkersProxy) DelegatedBatchedMarkersProxyArchive(contractID string, args DelegatedBatchedMarkersProxyArchive) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Commercials.V0.Model.DelegatedBatchedMarkersProxy", "DelegatedBatchedMarkersProxy"),
		ContractID: contractID,
		Choice:     "DelegatedBatchedMarkersProxy_Archive",
		Arguments:  argsToMap(args),
	}
}

// DelegatedBatchedMarkersProxyArchiveWithPackageID exercises the DelegatedBatchedMarkersProxy_Archive choice using the provided package ID instead of package name
func (t DelegatedBatchedMarkersProxy) DelegatedBatchedMarkersProxyArchiveWithPackageID(contractID string, packageID string, args DelegatedBatchedMarkersProxyArchive) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Commercials.V0.Model.DelegatedBatchedMarkersProxy", "DelegatedBatchedMarkersProxy"),
		ContractID: contractID,
		Choice:     "DelegatedBatchedMarkersProxy_Archive",
		Arguments:  argsToMap(args),
	}
}

// DelegatedBatchedMarkersProxyArchive is a Record type
type DelegatedBatchedMarkersProxyArchive struct {
}

// ToMap converts DelegatedBatchedMarkersProxyArchive to a map for DAML arguments
func (t DelegatedBatchedMarkersProxyArchive) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t DelegatedBatchedMarkersProxyArchive) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *DelegatedBatchedMarkersProxyArchive) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes DelegatedBatchedMarkersProxyArchive to hex string (Canton MCMS format)
func (t DelegatedBatchedMarkersProxyArchive) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes DelegatedBatchedMarkersProxyArchive from hex string (Canton MCMS format)
func (t *DelegatedBatchedMarkersProxyArchive) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// DelegatedBatchedMarkersProxyCreateMarkers is a Record type
type DelegatedBatchedMarkersProxyCreateMarkers struct {
	FeaturedAppRightCid types.CONTRACT_ID `json:"featuredAppRightCid"`
	Batches             []RewardBatch     `json:"batches"`
}

// ToMap converts DelegatedBatchedMarkersProxyCreateMarkers to a map for DAML arguments
func (t DelegatedBatchedMarkersProxyCreateMarkers) ToMap() map[string]any {
	m := make(map[string]any)

	m["featuredAppRightCid"] = model.NestedToDAMLValue(t.FeaturedAppRightCid)

	m["batches"] = func() []any {
		res := make([]any, 0, len(t.Batches))
		for _, e := range t.Batches {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	return m
}

func (t DelegatedBatchedMarkersProxyCreateMarkers) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *DelegatedBatchedMarkersProxyCreateMarkers) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes DelegatedBatchedMarkersProxyCreateMarkers to hex string (Canton MCMS format)
func (t DelegatedBatchedMarkersProxyCreateMarkers) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes DelegatedBatchedMarkersProxyCreateMarkers from hex string (Canton MCMS format)
func (t *DelegatedBatchedMarkersProxyCreateMarkers) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// DelegatedBatchedMarkersProxyCreateMarkersResult is a Record type
type DelegatedBatchedMarkersProxyCreateMarkersResult struct {
	Results []FeaturedAppRightCreateActivityMarkerResult `json:"results"`
}

// ToMap converts DelegatedBatchedMarkersProxyCreateMarkersResult to a map for DAML arguments
func (t DelegatedBatchedMarkersProxyCreateMarkersResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["results"] = func() []any {
		res := make([]any, 0, len(t.Results))
		for _, e := range t.Results {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	return m
}

func (t DelegatedBatchedMarkersProxyCreateMarkersResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *DelegatedBatchedMarkersProxyCreateMarkersResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes DelegatedBatchedMarkersProxyCreateMarkersResult to hex string (Canton MCMS format)
func (t DelegatedBatchedMarkersProxyCreateMarkersResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes DelegatedBatchedMarkersProxyCreateMarkersResult from hex string (Canton MCMS format)
func (t *DelegatedBatchedMarkersProxyCreateMarkersResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// EventBillingState is a Record type
type EventBillingState struct {
	Status           BillingStatus   `json:"status"`
	LastBilledAt     types.TIMESTAMP `json:"lastBilledAt"`
	LastBilledOffset types.INT64     `json:"lastBilledOffset"`
	MigrationId      *types.TEXT     `json:"migrationId" hex:"optional"`
}

// ToMap converts EventBillingState to a map for DAML arguments
func (t EventBillingState) ToMap() map[string]any {
	m := make(map[string]any)

	m["status"] = model.NestedToDAMLValue(t.Status)

	m["lastBilledAt"] = t.LastBilledAt

	m["lastBilledOffset"] = int64(t.LastBilledOffset)

	if t.MigrationId != nil {
		m["migrationId"] = map[string]any{
			"_type": "optional",
			"value": string(*t.MigrationId),
		}
	} else {
		m["migrationId"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t EventBillingState) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *EventBillingState) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes EventBillingState to hex string (Canton MCMS format)
func (t EventBillingState) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes EventBillingState from hex string (Canton MCMS format)
func (t *EventBillingState) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Failure is a Record type
type Failure struct {
	Reason  types.TEXT `json:"reason"`
	Context types.TEXT `json:"context"`
}

// ToMap converts Failure to a map for DAML arguments
func (t Failure) ToMap() map[string]any {
	m := make(map[string]any)

	m["reason"] = string(t.Reason)

	m["context"] = string(t.Context)

	return m
}

func (t Failure) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *Failure) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes Failure to hex string (Canton MCMS format)
func (t Failure) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes Failure from hex string (Canton MCMS format)
func (t *Failure) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// FixedFee2 is a Record type
type FixedFee2 struct {
	FeePerDayUsd         RatePerDay2 `json:"feePerDayUsd"`
	BillingPeriodMinutes types.INT64 `json:"billingPeriodMinutes"`
}

// ToMap converts FixedFee2 to a map for DAML arguments
func (t FixedFee2) ToMap() map[string]any {
	m := make(map[string]any)

	m["feePerDayUsd"] = model.NestedToDAMLValue(t.FeePerDayUsd)

	m["billingPeriodMinutes"] = int64(t.BillingPeriodMinutes)

	return m
}

func (t FixedFee2) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *FixedFee2) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes FixedFee2 to hex string (Canton MCMS format)
func (t FixedFee2) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes FixedFee2 from hex string (Canton MCMS format)
func (t *FixedFee2) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// RatePerDay2 is a Record type
type RatePerDay2 struct {
	Rate types.NUMERIC `json:"rate"`
}

// ToMap converts RatePerDay2 to a map for DAML arguments
func (t RatePerDay2) ToMap() map[string]any {
	m := make(map[string]any)

	m["rate"] = t.Rate

	return m
}

func (t RatePerDay2) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *RatePerDay2) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes RatePerDay2 to hex string (Canton MCMS format)
func (t RatePerDay2) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes RatePerDay2 from hex string (Canton MCMS format)
func (t *RatePerDay2) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// RewardBatch is a Record type
type RewardBatch struct {
	Beneficiaries []AppRewardBeneficiary `json:"beneficiaries"`
	MarkerWeight  types.NUMERIC          `json:"markerWeight"`
}

// ToMap converts RewardBatch to a map for DAML arguments
func (t RewardBatch) ToMap() map[string]any {
	m := make(map[string]any)

	m["beneficiaries"] = func() []any {
		res := make([]any, 0, len(t.Beneficiaries))
		for _, e := range t.Beneficiaries {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	m["markerWeight"] = t.MarkerWeight

	return m
}

func (t RewardBatch) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *RewardBatch) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes RewardBatch to hex string (Canton MCMS format)
func (t RewardBatch) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes RewardBatch from hex string (Canton MCMS format)
func (t *RewardBatch) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UtilityFees is a Record type
type UtilityFees struct {
	CredentialBillingFeeUsd *types.NUMERIC `json:"credentialBillingFeeUsd" hex:"optional"`
	BaseFee                 *FixedFee2     `json:"baseFee" hex:"optional"`
}

// ToMap converts UtilityFees to a map for DAML arguments
func (t UtilityFees) ToMap() map[string]any {
	m := make(map[string]any)

	if t.CredentialBillingFeeUsd != nil {
		m["credentialBillingFeeUsd"] = map[string]any{
			"_type": "optional",
			"value": *t.CredentialBillingFeeUsd,
		}
	} else {
		m["credentialBillingFeeUsd"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.BaseFee != nil {
		m["baseFee"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.BaseFee),
		}
	} else {
		m["baseFee"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t UtilityFees) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UtilityFees) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UtilityFees to hex string (Canton MCMS format)
func (t UtilityFees) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UtilityFees from hex string (Canton MCMS format)
func (t *UtilityFees) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// MCMSEncoder interface for typed encoding methods.
// Implemented by Encoder for method-based encoding.
type MCMSEncoder interface {
	CommercialAgreementOfferAccept(args CommercialAgreementOfferAccept) (*bind.EncodedChoice, error)
	CommercialAgreementOfferAcceptAndTopup(args CommercialAgreementOfferAcceptAndTopup) (*bind.EncodedChoice, error)
	CommercialAgreementOfferCancel(args CommercialAgreementOfferCancel) (*bind.EncodedChoice, error)
	CommercialAgreementOfferReject(args CommercialAgreementOfferReject) (*bind.EncodedChoice, error)
	CommercialAgreementBill(args CommercialAgreementBill) (*bind.EncodedChoice, error)
	CommercialAgreementBillBaseFee(args CommercialAgreementBillBaseFee) (*bind.EncodedChoice, error)
	CommercialAgreementBillCredentialFeeMulti(args CommercialAgreementBillCredentialFeeMulti) (*bind.EncodedChoice, error)
	CommercialAgreementBillCredentialFeeMultiUnfeatured(args CommercialAgreementBillCredentialFeeMultiUnfeatured) (*bind.EncodedChoice, error)
	CommercialAgreementFlushExpiredDeposit(args CommercialAgreementFlushExpiredDeposit) (*bind.EncodedChoice, error)
	CommercialAgreementLockCoin(args CommercialAgreementLockCoin) (*bind.EncodedChoice, error)
	CommercialAgreementModify(args CommercialAgreementModify) (*bind.EncodedChoice, error)
	CommercialAgreementModifyDataPublishingConsent(args CommercialAgreementModifyDataPublishingConsent) (*bind.EncodedChoice, error)
	CommercialAgreementRevoke(args CommercialAgreementRevoke) (*bind.EncodedChoice, error)
	CommercialAgreementSetDefaultCredentialFeeBillingState(args CommercialAgreementSetDefaultCredentialFeeBillingState) (*bind.EncodedChoice, error)
	DelegatedBatchedMarkersProxyArchive(args DelegatedBatchedMarkersProxyArchive) (*bind.EncodedChoice, error)
	DelegatedBatchedMarkersProxyCreateMarkers(args DelegatedBatchedMarkersProxyCreateMarkers) (*bind.EncodedChoice, error)
}

// encoder provides typed encoding methods for choice parameters (unexported).
// It wraps bind.BoundTemplate to encode parameters to hex-encoded operation data.
type encoder struct {
	*bind.BoundTemplate
}

// Contract wraps template operations with Sui-style API access.
// Use NewContract to create instances, then call Encoder() for encoding methods.
type Contract struct {
	enc *encoder
}

// NewContract creates a Contract with encoder for the given template.
// This provides Sui-style API: contract.Encoder().Method(args)
func NewContract(packageID, moduleName, templateName string) *Contract {
	return &Contract{
		enc: &encoder{
			BoundTemplate: bind.NewBoundTemplate(packageID, moduleName, templateName),
		},
	}
}

// Encoder returns the encoder for Sui-style contract.Encoder().Method() usage.
func (c *Contract) Encoder() MCMSEncoder {
	return c.enc
}

// CommercialAgreementOfferAccept encodes parameters for the CommercialAgreementOffer_Accept choice.
func (e *encoder) CommercialAgreementOfferAccept(args CommercialAgreementOfferAccept) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreementOffer_Accept", args)
}

// CommercialAgreementOfferAcceptAndTopup encodes parameters for the CommercialAgreementOffer_AcceptAndTopup choice.
func (e *encoder) CommercialAgreementOfferAcceptAndTopup(args CommercialAgreementOfferAcceptAndTopup) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreementOffer_AcceptAndTopup", args)
}

// CommercialAgreementOfferCancel encodes parameters for the CommercialAgreementOffer_Cancel choice.
func (e *encoder) CommercialAgreementOfferCancel(args CommercialAgreementOfferCancel) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreementOffer_Cancel", args)
}

// CommercialAgreementOfferReject encodes parameters for the CommercialAgreementOffer_Reject choice.
func (e *encoder) CommercialAgreementOfferReject(args CommercialAgreementOfferReject) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreementOffer_Reject", args)
}

// CommercialAgreementBill encodes parameters for the CommercialAgreement_Bill choice.
func (e *encoder) CommercialAgreementBill(args CommercialAgreementBill) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_Bill", args)
}

// CommercialAgreementBillBaseFee encodes parameters for the CommercialAgreement_BillBaseFee choice.
func (e *encoder) CommercialAgreementBillBaseFee(args CommercialAgreementBillBaseFee) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_BillBaseFee", args)
}

// CommercialAgreementBillCredentialFeeMulti encodes parameters for the CommercialAgreement_BillCredentialFeeMulti choice.
func (e *encoder) CommercialAgreementBillCredentialFeeMulti(args CommercialAgreementBillCredentialFeeMulti) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_BillCredentialFeeMulti", args)
}

// CommercialAgreementBillCredentialFeeMultiUnfeatured encodes parameters for the CommercialAgreement_BillCredentialFeeMultiUnfeatured choice.
func (e *encoder) CommercialAgreementBillCredentialFeeMultiUnfeatured(args CommercialAgreementBillCredentialFeeMultiUnfeatured) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_BillCredentialFeeMultiUnfeatured", args)
}

// CommercialAgreementFlushExpiredDeposit encodes parameters for the CommercialAgreement_FlushExpiredDeposit choice.
func (e *encoder) CommercialAgreementFlushExpiredDeposit(args CommercialAgreementFlushExpiredDeposit) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_FlushExpiredDeposit", args)
}

// CommercialAgreementLockCoin encodes parameters for the CommercialAgreement_LockCoin choice.
func (e *encoder) CommercialAgreementLockCoin(args CommercialAgreementLockCoin) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_LockCoin", args)
}

// CommercialAgreementModify encodes parameters for the CommercialAgreement_Modify choice.
func (e *encoder) CommercialAgreementModify(args CommercialAgreementModify) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_Modify", args)
}

// CommercialAgreementModifyDataPublishingConsent encodes parameters for the CommercialAgreement_ModifyDataPublishingConsent choice.
func (e *encoder) CommercialAgreementModifyDataPublishingConsent(args CommercialAgreementModifyDataPublishingConsent) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_ModifyDataPublishingConsent", args)
}

// CommercialAgreementRevoke encodes parameters for the CommercialAgreement_Revoke choice.
func (e *encoder) CommercialAgreementRevoke(args CommercialAgreementRevoke) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_Revoke", args)
}

// CommercialAgreementSetDefaultCredentialFeeBillingState encodes parameters for the CommercialAgreement_SetDefaultCredentialFeeBillingState choice.
func (e *encoder) CommercialAgreementSetDefaultCredentialFeeBillingState(args CommercialAgreementSetDefaultCredentialFeeBillingState) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CommercialAgreement_SetDefaultCredentialFeeBillingState", args)
}

// DelegatedBatchedMarkersProxyArchive encodes parameters for the DelegatedBatchedMarkersProxy_Archive choice.
func (e *encoder) DelegatedBatchedMarkersProxyArchive(args DelegatedBatchedMarkersProxyArchive) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("DelegatedBatchedMarkersProxy_Archive", args)
}

// DelegatedBatchedMarkersProxyCreateMarkers encodes parameters for the DelegatedBatchedMarkersProxy_CreateMarkers choice.
func (e *encoder) DelegatedBatchedMarkersProxyCreateMarkers(args DelegatedBatchedMarkersProxyCreateMarkers) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("DelegatedBatchedMarkersProxy_CreateMarkers", args)
}

// Verify MCMSEncoder interface implementation
var _ MCMSEncoder = (*encoder)(nil)
