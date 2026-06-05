package credential_app_v0

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	credential_v0 "github.com/smartcontractkit/chainlink-canton/bindings/generated/latest/utility/credential_v0"
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
	PackageName = "utility-credential-app-v0"
	PackageID   = "e9a3b7df354dfd2f15c7d015328c34256308c90ba96f86f185dad58ffca8299b"
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

// BalanceState is a Record type
type BalanceState struct {
	CurrentDepositAmountCc    types.NUMERIC `json:"currentDepositAmountCc"`
	TotalCredentialFeesPaidCc types.NUMERIC `json:"totalCredentialFeesPaidCc"`
	TotalDistributedCc        types.NUMERIC `json:"totalDistributedCc"`
	TotalPaidOutCc            types.NUMERIC `json:"totalPaidOutCc"`
	TotalUserDepositCc        types.NUMERIC `json:"totalUserDepositCc"`
}

// ToMap converts BalanceState to a map for DAML arguments
func (t BalanceState) ToMap() map[string]any {
	m := make(map[string]any)

	m["currentDepositAmountCc"] = t.CurrentDepositAmountCc

	m["totalCredentialFeesPaidCc"] = t.TotalCredentialFeesPaidCc

	m["totalDistributedCc"] = t.TotalDistributedCc

	m["totalPaidOutCc"] = t.TotalPaidOutCc

	m["totalUserDepositCc"] = t.TotalUserDepositCc

	return m
}

func (t BalanceState) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BalanceState) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BalanceState to hex string (Canton MCMS format)
func (t BalanceState) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BalanceState from hex string (Canton MCMS format)
func (t *BalanceState) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// BillingContext is a Record type
type BillingContext struct {
	Now                       types.TIMESTAMP    `json:"now"`
	AmuletRulesCid            types.CONTRACT_ID  `json:"amuletRulesCid"`
	OpenRoundCid              types.CONTRACT_ID  `json:"openRoundCid"`
	OpenRound                 OpenMiningRound    `json:"openRound"`
	FeaturedTransferContext   AppTransferContext `json:"featuredTransferContext"`
	UnfeaturedTransferContext AppTransferContext `json:"unfeaturedTransferContext"`
	FeeComputationContext     TransferContext    `json:"feeComputationContext"`
}

// ToMap converts BillingContext to a map for DAML arguments
func (t BillingContext) ToMap() map[string]any {
	m := make(map[string]any)

	m["now"] = t.Now

	m["amuletRulesCid"] = model.NestedToDAMLValue(t.AmuletRulesCid)

	m["openRoundCid"] = model.NestedToDAMLValue(t.OpenRoundCid)

	m["openRound"] = model.NestedToDAMLValue(t.OpenRound)

	m["featuredTransferContext"] = model.NestedToDAMLValue(t.FeaturedTransferContext)

	m["unfeaturedTransferContext"] = model.NestedToDAMLValue(t.UnfeaturedTransferContext)

	m["feeComputationContext"] = model.NestedToDAMLValue(t.FeeComputationContext)

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
	CredentialFeeCc  types.NUMERIC   `json:"credentialFeeCc"`
	NewBilledUntil   types.TIMESTAMP `json:"newBilledUntil"`
	DepositExpiresAt types.TIMESTAMP `json:"depositExpiresAt"`
}

// ToMap converts BillingCycleParams to a map for DAML arguments
func (t BillingCycleParams) ToMap() map[string]any {
	m := make(map[string]any)

	m["amuletPrice"] = t.AmuletPrice

	m["credentialFeeCc"] = t.CredentialFeeCc

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

// BillingParams is a Record type
type BillingParams struct {
	FeePerDayUsd           RatePerDay2    `json:"feePerDayUsd"`
	BillingPeriodMinutes   types.INT64    `json:"billingPeriodMinutes"`
	DepositTargetAmountUsd types.NUMERIC  `json:"depositTargetAmountUsd"`
	HolderActivityWeight   *types.NUMERIC `json:"holderActivityWeight" hex:"optional"`
}

// ToMap converts BillingParams to a map for DAML arguments
func (t BillingParams) ToMap() map[string]any {
	m := make(map[string]any)

	m["feePerDayUsd"] = model.NestedToDAMLValue(t.FeePerDayUsd)

	m["billingPeriodMinutes"] = int64(t.BillingPeriodMinutes)

	m["depositTargetAmountUsd"] = t.DepositTargetAmountUsd

	if t.HolderActivityWeight != nil {
		m["holderActivityWeight"] = map[string]any{
			"_type": "optional",
			"value": *t.HolderActivityWeight,
		}
	} else {
		m["holderActivityWeight"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t BillingParams) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BillingParams) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BillingParams to hex string (Canton MCMS format)
func (t BillingParams) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BillingParams from hex string (Canton MCMS format)
func (t *BillingParams) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// BillingParamsAdjustmentRequest is a Template type
type BillingParamsAdjustmentRequest struct {
	Operator     types.PARTY   `json:"operator"`
	Issuer       types.PARTY   `json:"issuer"`
	Holder       types.PARTY   `json:"holder"`
	Params       BillingParams `json:"params"`
	CredentialId types.TEXT    `json:"credentialId"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t BillingParamsAdjustmentRequest) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequest")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t BillingParamsAdjustmentRequest) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequest")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t BillingParamsAdjustmentRequest) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["issuer"] = t.Issuer.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["holder"] = t.Holder.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["params"] = model.NestedToDAMLValue(t.Params)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["credentialId"] = string(t.CredentialId)

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t BillingParamsAdjustmentRequest) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["issuer"] = t.Issuer.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["holder"] = t.Holder.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["params"] = model.NestedToDAMLValue(t.Params)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["credentialId"] = string(t.CredentialId)

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t BillingParamsAdjustmentRequest) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BillingParamsAdjustmentRequest) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BillingParamsAdjustmentRequest to hex string (Canton MCMS format)
func (t BillingParamsAdjustmentRequest) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BillingParamsAdjustmentRequest from hex string (Canton MCMS format)
func (t *BillingParamsAdjustmentRequest) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for BillingParamsAdjustmentRequest

// BillingParamsAdjustmentRequestAccept exercises the BillingParamsAdjustmentRequest_Accept choice on this BillingParamsAdjustmentRequest contract
// This method uses the package name in the template ID
func (t BillingParamsAdjustmentRequest) BillingParamsAdjustmentRequestAccept(contractID string, args BillingParamsAdjustmentRequestAccept) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequest"),
		ContractID: contractID,
		Choice:     "BillingParamsAdjustmentRequest_Accept",
		Arguments:  argsToMap(args),
	}
}

// BillingParamsAdjustmentRequestAcceptWithPackageID exercises the BillingParamsAdjustmentRequest_Accept choice using the provided package ID instead of package name
func (t BillingParamsAdjustmentRequest) BillingParamsAdjustmentRequestAcceptWithPackageID(contractID string, packageID string, args BillingParamsAdjustmentRequestAccept) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequest"),
		ContractID: contractID,
		Choice:     "BillingParamsAdjustmentRequest_Accept",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this BillingParamsAdjustmentRequest contract
// This method uses the package name in the template ID
func (t BillingParamsAdjustmentRequest) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequest"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t BillingParamsAdjustmentRequest) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequest"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// BillingParamsAdjustmentRequestCancel exercises the BillingParamsAdjustmentRequest_Cancel choice on this BillingParamsAdjustmentRequest contract
// This method uses the package name in the template ID
func (t BillingParamsAdjustmentRequest) BillingParamsAdjustmentRequestCancel(contractID string, args BillingParamsAdjustmentRequestCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequest"),
		ContractID: contractID,
		Choice:     "BillingParamsAdjustmentRequest_Cancel",
		Arguments:  argsToMap(args),
	}
}

// BillingParamsAdjustmentRequestCancelWithPackageID exercises the BillingParamsAdjustmentRequest_Cancel choice using the provided package ID instead of package name
func (t BillingParamsAdjustmentRequest) BillingParamsAdjustmentRequestCancelWithPackageID(contractID string, packageID string, args BillingParamsAdjustmentRequestCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequest"),
		ContractID: contractID,
		Choice:     "BillingParamsAdjustmentRequest_Cancel",
		Arguments:  argsToMap(args),
	}
}

// BillingParamsAdjustmentRequestAccept is a Record type
type BillingParamsAdjustmentRequestAccept struct {
	CredentialBillingCid types.CONTRACT_ID `json:"credentialBillingCid"`
}

// ToMap converts BillingParamsAdjustmentRequestAccept to a map for DAML arguments
func (t BillingParamsAdjustmentRequestAccept) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	return m
}

func (t BillingParamsAdjustmentRequestAccept) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BillingParamsAdjustmentRequestAccept) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BillingParamsAdjustmentRequestAccept to hex string (Canton MCMS format)
func (t BillingParamsAdjustmentRequestAccept) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BillingParamsAdjustmentRequestAccept from hex string (Canton MCMS format)
func (t *BillingParamsAdjustmentRequestAccept) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// BillingParamsAdjustmentRequestAcceptResult is a Record type
type BillingParamsAdjustmentRequestAcceptResult struct {
	CredentialBillingCid types.CONTRACT_ID `json:"credentialBillingCid"`
}

// ToMap converts BillingParamsAdjustmentRequestAcceptResult to a map for DAML arguments
func (t BillingParamsAdjustmentRequestAcceptResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	return m
}

func (t BillingParamsAdjustmentRequestAcceptResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BillingParamsAdjustmentRequestAcceptResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BillingParamsAdjustmentRequestAcceptResult to hex string (Canton MCMS format)
func (t BillingParamsAdjustmentRequestAcceptResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BillingParamsAdjustmentRequestAcceptResult from hex string (Canton MCMS format)
func (t *BillingParamsAdjustmentRequestAcceptResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// BillingParamsAdjustmentRequestCancel is a Record type
type BillingParamsAdjustmentRequestCancel struct {
	Actor types.PARTY `json:"actor"`
}

// ToMap converts BillingParamsAdjustmentRequestCancel to a map for DAML arguments
func (t BillingParamsAdjustmentRequestCancel) ToMap() map[string]any {
	m := make(map[string]any)

	m["actor"] = t.Actor.ToMap()

	return m
}

func (t BillingParamsAdjustmentRequestCancel) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *BillingParamsAdjustmentRequestCancel) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes BillingParamsAdjustmentRequestCancel to hex string (Canton MCMS format)
func (t BillingParamsAdjustmentRequestCancel) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes BillingParamsAdjustmentRequestCancel from hex string (Canton MCMS format)
func (t *BillingParamsAdjustmentRequestCancel) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// BillingParamsAdjustmentRequestCancelResult is an enum type
type BillingParamsAdjustmentRequestCancelResult string

const (
	BillingParamsAdjustmentRequestCancelResultBillingParamsAdjustmentRequest_Cancel_Result BillingParamsAdjustmentRequestCancelResult = "BillingParamsAdjustmentRequest_Cancel_Result"
)

func (e BillingParamsAdjustmentRequestCancelResult) GetEnumConstructor() string { return string(e) }

func (e BillingParamsAdjustmentRequestCancelResult) GetEnumTypeID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequestCancelResult")
}

// GetEnumTypeIDWithPackageID returns the enum type ID using the provided package ID instead of package name
func (e BillingParamsAdjustmentRequestCancelResult) GetEnumTypeIDWithPackageID(packageID string) string {
	return fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "BillingParamsAdjustmentRequestCancelResult")
}

func (e BillingParamsAdjustmentRequestCancelResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(e)
}

func (e *BillingParamsAdjustmentRequestCancelResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, e)
}

// MarshalHex encodes BillingParamsAdjustmentRequestCancelResult to hex string (Canton MCMS format)
func (e BillingParamsAdjustmentRequestCancelResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(e)
}

// UnmarshalHex decodes BillingParamsAdjustmentRequestCancelResult from hex string (Canton MCMS format)
func (e *BillingParamsAdjustmentRequestCancelResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, e)
}

var _ types.ENUM = BillingParamsAdjustmentRequestCancelResult("")

// BillingState is a Record type
type BillingState struct {
	CreatedAt               types.TIMESTAMP `json:"createdAt"`
	Status                  BillingStatus   `json:"status"`
	LastBilledAt            types.TIMESTAMP `json:"lastBilledAt"`
	BilledUntil             types.TIMESTAMP `json:"billedUntil"`
	LastBilledInRound       Round           `json:"lastBilledInRound"`
	TotalCcFeesPaidIssuerCc types.NUMERIC   `json:"totalCcFeesPaidIssuerCc"`
	TotalCcFeesPaidHolderCc types.NUMERIC   `json:"totalCcFeesPaidHolderCc"`
}

// ToMap converts BillingState to a map for DAML arguments
func (t BillingState) ToMap() map[string]any {
	m := make(map[string]any)

	m["createdAt"] = t.CreatedAt

	m["status"] = model.NestedToDAMLValue(t.Status)

	m["lastBilledAt"] = t.LastBilledAt

	m["billedUntil"] = t.BilledUntil

	m["lastBilledInRound"] = model.NestedToDAMLValue(t.LastBilledInRound)

	m["totalCcFeesPaidIssuerCc"] = t.TotalCcFeesPaidIssuerCc

	m["totalCcFeesPaidHolderCc"] = t.TotalCcFeesPaidHolderCc

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

// CanceledCredentialBilling is a Template type
type CanceledCredentialBilling struct {
	Payload              CredentialBilling `json:"payload"`
	CancelledBy          types.PARTY       `json:"cancelledBy"`
	CancelledAt          types.TIMESTAMP   `json:"cancelledAt"`
	ReturnedUserAmountCc types.NUMERIC     `json:"returnedUserAmountCc"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t CanceledCredentialBilling) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CanceledCredentialBilling")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t CanceledCredentialBilling) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CanceledCredentialBilling")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t CanceledCredentialBilling) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["payload"] = model.NestedToDAMLValue(t.Payload)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["cancelledBy"] = t.CancelledBy.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["cancelledAt"] = t.CancelledAt

	if t.ReturnedUserAmountCc != "" {
		args["returnedUserAmountCc"] = t.ReturnedUserAmountCc
	}

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t CanceledCredentialBilling) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["payload"] = model.NestedToDAMLValue(t.Payload)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["cancelledBy"] = t.CancelledBy.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["cancelledAt"] = t.CancelledAt

	if t.ReturnedUserAmountCc != "" {
		args["returnedUserAmountCc"] = t.ReturnedUserAmountCc
	}

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t CanceledCredentialBilling) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CanceledCredentialBilling) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CanceledCredentialBilling to hex string (Canton MCMS format)
func (t CanceledCredentialBilling) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CanceledCredentialBilling from hex string (Canton MCMS format)
func (t *CanceledCredentialBilling) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for CanceledCredentialBilling

// Archive exercises the Archive choice on this CanceledCredentialBilling contract
// This method uses the package name in the template ID
func (t CanceledCredentialBilling) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CanceledCredentialBilling"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t CanceledCredentialBilling) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CanceledCredentialBilling"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// CredentialBilling is a Template type
type CredentialBilling struct {
	Operator     types.PARTY         `json:"operator"`
	Issuer       types.PARTY         `json:"issuer"`
	Holder       types.PARTY         `json:"holder"`
	Dso          types.PARTY         `json:"dso"`
	CredentialId types.TEXT          `json:"credentialId"`
	Params       BillingParams       `json:"params"`
	BalanceState BalanceState        `json:"balanceState"`
	BillingState BillingState        `json:"billingState"`
	Deposits     []types.CONTRACT_ID `json:"deposits"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t CredentialBilling) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t CredentialBilling) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t CredentialBilling) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["issuer"] = t.Issuer.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["holder"] = t.Holder.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["credentialId"] = string(t.CredentialId)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["params"] = model.NestedToDAMLValue(t.Params)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["balanceState"] = model.NestedToDAMLValue(t.BalanceState)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["billingState"] = model.NestedToDAMLValue(t.BillingState)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["deposits"] = func() []any {
		res := make([]any, 0, len(t.Deposits))
		for _, e := range t.Deposits {
			res = append(res, e)
		}
		return res
	}()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t CredentialBilling) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["issuer"] = t.Issuer.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["holder"] = t.Holder.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["credentialId"] = string(t.CredentialId)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["params"] = model.NestedToDAMLValue(t.Params)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["balanceState"] = model.NestedToDAMLValue(t.BalanceState)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["billingState"] = model.NestedToDAMLValue(t.BillingState)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["deposits"] = func() []any {
		res := make([]any, 0, len(t.Deposits))
		for _, e := range t.Deposits {
			res = append(res, e)
		}
		return res
	}()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t CredentialBilling) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBilling) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBilling to hex string (Canton MCMS format)
func (t CredentialBilling) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBilling from hex string (Canton MCMS format)
func (t *CredentialBilling) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for CredentialBilling

// CredentialBillingRequestToAdjustBillingParams exercises the CredentialBilling_RequestToAdjustBillingParams choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingRequestToAdjustBillingParams(contractID string, args CredentialBillingRequestToAdjustBillingParams) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_RequestToAdjustBillingParams",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingRequestToAdjustBillingParamsWithPackageID exercises the CredentialBilling_RequestToAdjustBillingParams choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingRequestToAdjustBillingParamsWithPackageID(contractID string, packageID string, args CredentialBillingRequestToAdjustBillingParams) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_RequestToAdjustBillingParams",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingDistributeAndAdjustDeposit exercises the CredentialBilling_DistributeAndAdjustDeposit choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingDistributeAndAdjustDeposit(contractID string, args CredentialBillingDistributeAndAdjustDeposit) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_DistributeAndAdjustDeposit",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingDistributeAndAdjustDepositWithPackageID exercises the CredentialBilling_DistributeAndAdjustDeposit choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingDistributeAndAdjustDepositWithPackageID(contractID string, packageID string, args CredentialBillingDistributeAndAdjustDeposit) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_DistributeAndAdjustDeposit",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingBill exercises the CredentialBilling_Bill choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingBill(contractID string, args CredentialBillingBill) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_Bill",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingBillWithPackageID exercises the CredentialBilling_Bill choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingBillWithPackageID(contractID string, packageID string, args CredentialBillingBill) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_Bill",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingCancel exercises the CredentialBilling_Cancel choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingCancel(contractID string, args CredentialBillingCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_Cancel",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingCancelWithPackageID exercises the CredentialBilling_Cancel choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingCancelWithPackageID(contractID string, packageID string, args CredentialBillingCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_Cancel",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingTopUp exercises the CredentialBilling_TopUp choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingTopUp(contractID string, args CredentialBillingTopUp) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_TopUp",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingTopUpWithPackageID exercises the CredentialBilling_TopUp choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingTopUpWithPackageID(contractID string, packageID string, args CredentialBillingTopUp) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_TopUp",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingDistribute exercises the CredentialBilling_Distribute choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingDistribute(contractID string, args CredentialBillingDistribute) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_Distribute",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingDistributeWithPackageID exercises the CredentialBilling_Distribute choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingDistributeWithPackageID(contractID string, packageID string, args CredentialBillingDistribute) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_Distribute",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingCancelExpired exercises the CredentialBilling_CancelExpired choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingCancelExpired(contractID string, args CredentialBillingCancelExpired) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_CancelExpired",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingCancelExpiredWithPackageID exercises the CredentialBilling_CancelExpired choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingCancelExpiredWithPackageID(contractID string, packageID string, args CredentialBillingCancelExpired) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_CancelExpired",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingFlushExpiredDeposit exercises the CredentialBilling_FlushExpiredDeposit choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingFlushExpiredDeposit(contractID string, args CredentialBillingFlushExpiredDeposit) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_FlushExpiredDeposit",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingFlushExpiredDepositWithPackageID exercises the CredentialBilling_FlushExpiredDeposit choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingFlushExpiredDepositWithPackageID(contractID string, packageID string, args CredentialBillingFlushExpiredDeposit) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_FlushExpiredDeposit",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t CredentialBilling) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// CredentialBillingAdjustBillingParams exercises the CredentialBilling_AdjustBillingParams choice on this CredentialBilling contract
// This method uses the package name in the template ID
func (t CredentialBilling) CredentialBillingAdjustBillingParams(contractID string, args CredentialBillingAdjustBillingParams) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_AdjustBillingParams",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingAdjustBillingParamsWithPackageID exercises the CredentialBilling_AdjustBillingParams choice using the provided package ID instead of package name
func (t CredentialBilling) CredentialBillingAdjustBillingParamsWithPackageID(contractID string, packageID string, args CredentialBillingAdjustBillingParams) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Billing", "CredentialBilling"),
		ContractID: contractID,
		Choice:     "CredentialBilling_AdjustBillingParams",
		Arguments:  argsToMap(args),
	}
}

// CredentialBillingAdjustBillingParams is a Record type
type CredentialBillingAdjustBillingParams struct {
	NewParams BillingParams `json:"newParams"`
}

// ToMap converts CredentialBillingAdjustBillingParams to a map for DAML arguments
func (t CredentialBillingAdjustBillingParams) ToMap() map[string]any {
	m := make(map[string]any)

	m["newParams"] = model.NestedToDAMLValue(t.NewParams)

	return m
}

func (t CredentialBillingAdjustBillingParams) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingAdjustBillingParams) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingAdjustBillingParams to hex string (Canton MCMS format)
func (t CredentialBillingAdjustBillingParams) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingAdjustBillingParams from hex string (Canton MCMS format)
func (t *CredentialBillingAdjustBillingParams) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingAdjustBillingParamsResult is a Record type
type CredentialBillingAdjustBillingParamsResult struct {
	NewCredentialBillingCid types.CONTRACT_ID `json:"newCredentialBillingCid"`
}

// ToMap converts CredentialBillingAdjustBillingParamsResult to a map for DAML arguments
func (t CredentialBillingAdjustBillingParamsResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["newCredentialBillingCid"] = model.NestedToDAMLValue(t.NewCredentialBillingCid)

	return m
}

func (t CredentialBillingAdjustBillingParamsResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingAdjustBillingParamsResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingAdjustBillingParamsResult to hex string (Canton MCMS format)
func (t CredentialBillingAdjustBillingParamsResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingAdjustBillingParamsResult from hex string (Canton MCMS format)
func (t *CredentialBillingAdjustBillingParamsResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingBill is a Record type
type CredentialBillingBill struct {
	AppTransferContext  AppTransferContext `json:"appTransferContext"`
	EnableFeeRecord     *types.BOOL        `json:"enableFeeRecord" hex:"optional"`
	RewardReceiver      *types.PARTY       `json:"rewardReceiver" hex:"optional"`
	FeaturedAppRightCid *types.CONTRACT_ID `json:"featuredAppRightCid" hex:"optional"`
}

// ToMap converts CredentialBillingBill to a map for DAML arguments
func (t CredentialBillingBill) ToMap() map[string]any {
	m := make(map[string]any)

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	if t.EnableFeeRecord != nil {
		m["enableFeeRecord"] = map[string]any{
			"_type": "optional",
			"value": bool(*t.EnableFeeRecord),
		}
	} else {
		m["enableFeeRecord"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

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

	if t.FeaturedAppRightCid != nil {
		m["featuredAppRightCid"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.FeaturedAppRightCid),
		}
	} else {
		m["featuredAppRightCid"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CredentialBillingBill) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingBill) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingBill to hex string (Canton MCMS format)
func (t CredentialBillingBill) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingBill from hex string (Canton MCMS format)
func (t *CredentialBillingBill) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingBillResult is a Record type
type CredentialBillingBillResult struct {
	BillingCycleParams      BillingCycleParams `json:"billingCycleParams"`
	TransferResult          *TransferResult    `json:"transferResult" hex:"optional"`
	NewCredentialBillingCid types.CONTRACT_ID  `json:"newCredentialBillingCid"`
	FeeRecordCid            *types.CONTRACT_ID `json:"feeRecordCid" hex:"optional"`
}

// ToMap converts CredentialBillingBillResult to a map for DAML arguments
func (t CredentialBillingBillResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["billingCycleParams"] = model.NestedToDAMLValue(t.BillingCycleParams)

	if t.TransferResult != nil {
		m["transferResult"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.TransferResult),
		}
	} else {
		m["transferResult"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	m["newCredentialBillingCid"] = model.NestedToDAMLValue(t.NewCredentialBillingCid)

	if t.FeeRecordCid != nil {
		m["feeRecordCid"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.FeeRecordCid),
		}
	} else {
		m["feeRecordCid"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t CredentialBillingBillResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingBillResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingBillResult to hex string (Canton MCMS format)
func (t CredentialBillingBillResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingBillResult from hex string (Canton MCMS format)
func (t *CredentialBillingBillResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingCancel is a Record type
type CredentialBillingCancel struct {
	Actor              types.PARTY        `json:"actor"`
	AppTransferContext AppTransferContext `json:"appTransferContext"`
}

// ToMap converts CredentialBillingCancel to a map for DAML arguments
func (t CredentialBillingCancel) ToMap() map[string]any {
	m := make(map[string]any)

	m["actor"] = t.Actor.ToMap()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t CredentialBillingCancel) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingCancel) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingCancel to hex string (Canton MCMS format)
func (t CredentialBillingCancel) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingCancel from hex string (Canton MCMS format)
func (t *CredentialBillingCancel) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingCancelExpired is a Record type
type CredentialBillingCancelExpired struct {
	Actor types.PARTY `json:"actor"`
}

// ToMap converts CredentialBillingCancelExpired to a map for DAML arguments
func (t CredentialBillingCancelExpired) ToMap() map[string]any {
	m := make(map[string]any)

	m["actor"] = t.Actor.ToMap()

	return m
}

func (t CredentialBillingCancelExpired) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingCancelExpired) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingCancelExpired to hex string (Canton MCMS format)
func (t CredentialBillingCancelExpired) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingCancelExpired from hex string (Canton MCMS format)
func (t *CredentialBillingCancelExpired) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingCancelResult is a Record type
type CredentialBillingCancelResult struct {
	CanceledCredentialBillingCid types.CONTRACT_ID `json:"canceledCredentialBillingCid"`
}

// ToMap converts CredentialBillingCancelResult to a map for DAML arguments
func (t CredentialBillingCancelResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["canceledCredentialBillingCid"] = model.NestedToDAMLValue(t.CanceledCredentialBillingCid)

	return m
}

func (t CredentialBillingCancelResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingCancelResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingCancelResult to hex string (Canton MCMS format)
func (t CredentialBillingCancelResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingCancelResult from hex string (Canton MCMS format)
func (t *CredentialBillingCancelResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingDistribute is a Record type
type CredentialBillingDistribute struct {
	AmountUsd          types.NUMERIC       `json:"amountUsd"`
	CoinCids           []types.CONTRACT_ID `json:"coinCids"`
	AppTransferContext AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts CredentialBillingDistribute to a map for DAML arguments
func (t CredentialBillingDistribute) ToMap() map[string]any {
	m := make(map[string]any)

	m["amountUsd"] = t.AmountUsd

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t CredentialBillingDistribute) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingDistribute) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingDistribute to hex string (Canton MCMS format)
func (t CredentialBillingDistribute) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingDistribute from hex string (Canton MCMS format)
func (t *CredentialBillingDistribute) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingDistributeAndAdjustDeposit is a Record type
type CredentialBillingDistributeAndAdjustDeposit struct {
	AmountUsd          types.NUMERIC       `json:"amountUsd"`
	CoinCids           []types.CONTRACT_ID `json:"coinCids"`
	AppTransferContext AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts CredentialBillingDistributeAndAdjustDeposit to a map for DAML arguments
func (t CredentialBillingDistributeAndAdjustDeposit) ToMap() map[string]any {
	m := make(map[string]any)

	m["amountUsd"] = t.AmountUsd

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t CredentialBillingDistributeAndAdjustDeposit) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingDistributeAndAdjustDeposit) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingDistributeAndAdjustDeposit to hex string (Canton MCMS format)
func (t CredentialBillingDistributeAndAdjustDeposit) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingDistributeAndAdjustDeposit from hex string (Canton MCMS format)
func (t *CredentialBillingDistributeAndAdjustDeposit) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingDistributeAndAdjustDepositResult is a Record type
type CredentialBillingDistributeAndAdjustDepositResult struct {
	NewCredentialBillingCid types.CONTRACT_ID `json:"newCredentialBillingCid"`
	TransferResult          TransferResult    `json:"transferResult"`
}

// ToMap converts CredentialBillingDistributeAndAdjustDepositResult to a map for DAML arguments
func (t CredentialBillingDistributeAndAdjustDepositResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["newCredentialBillingCid"] = model.NestedToDAMLValue(t.NewCredentialBillingCid)

	m["transferResult"] = model.NestedToDAMLValue(t.TransferResult)

	return m
}

func (t CredentialBillingDistributeAndAdjustDepositResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingDistributeAndAdjustDepositResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingDistributeAndAdjustDepositResult to hex string (Canton MCMS format)
func (t CredentialBillingDistributeAndAdjustDepositResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingDistributeAndAdjustDepositResult from hex string (Canton MCMS format)
func (t *CredentialBillingDistributeAndAdjustDepositResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingDistributeResult is a Record type
type CredentialBillingDistributeResult struct {
	TransferResult TransferResult `json:"transferResult"`
}

// ToMap converts CredentialBillingDistributeResult to a map for DAML arguments
func (t CredentialBillingDistributeResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["transferResult"] = model.NestedToDAMLValue(t.TransferResult)

	return m
}

func (t CredentialBillingDistributeResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingDistributeResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingDistributeResult to hex string (Canton MCMS format)
func (t CredentialBillingDistributeResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingDistributeResult from hex string (Canton MCMS format)
func (t *CredentialBillingDistributeResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingFlushExpiredDeposit is a Record type
type CredentialBillingFlushExpiredDeposit struct {
	Actor types.PARTY `json:"actor"`
}

// ToMap converts CredentialBillingFlushExpiredDeposit to a map for DAML arguments
func (t CredentialBillingFlushExpiredDeposit) ToMap() map[string]any {
	m := make(map[string]any)

	m["actor"] = t.Actor.ToMap()

	return m
}

func (t CredentialBillingFlushExpiredDeposit) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingFlushExpiredDeposit) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingFlushExpiredDeposit to hex string (Canton MCMS format)
func (t CredentialBillingFlushExpiredDeposit) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingFlushExpiredDeposit from hex string (Canton MCMS format)
func (t *CredentialBillingFlushExpiredDeposit) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingFlushExpiredDepositResult is a Record type
type CredentialBillingFlushExpiredDepositResult struct {
	NewCredentialBillingCid types.CONTRACT_ID `json:"newCredentialBillingCid"`
}

// ToMap converts CredentialBillingFlushExpiredDepositResult to a map for DAML arguments
func (t CredentialBillingFlushExpiredDepositResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["newCredentialBillingCid"] = model.NestedToDAMLValue(t.NewCredentialBillingCid)

	return m
}

func (t CredentialBillingFlushExpiredDepositResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingFlushExpiredDepositResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingFlushExpiredDepositResult to hex string (Canton MCMS format)
func (t CredentialBillingFlushExpiredDepositResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingFlushExpiredDepositResult from hex string (Canton MCMS format)
func (t *CredentialBillingFlushExpiredDepositResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingRequestToAdjustBillingParams is a Record type
type CredentialBillingRequestToAdjustBillingParams struct {
	NewParams BillingParams `json:"newParams"`
}

// ToMap converts CredentialBillingRequestToAdjustBillingParams to a map for DAML arguments
func (t CredentialBillingRequestToAdjustBillingParams) ToMap() map[string]any {
	m := make(map[string]any)

	m["newParams"] = model.NestedToDAMLValue(t.NewParams)

	return m
}

func (t CredentialBillingRequestToAdjustBillingParams) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingRequestToAdjustBillingParams) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingRequestToAdjustBillingParams to hex string (Canton MCMS format)
func (t CredentialBillingRequestToAdjustBillingParams) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingRequestToAdjustBillingParams from hex string (Canton MCMS format)
func (t *CredentialBillingRequestToAdjustBillingParams) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingRequestToAdjustBillingParamsResult is a Record type
type CredentialBillingRequestToAdjustBillingParamsResult struct {
	RequestCid types.CONTRACT_ID `json:"requestCid"`
}

// ToMap converts CredentialBillingRequestToAdjustBillingParamsResult to a map for DAML arguments
func (t CredentialBillingRequestToAdjustBillingParamsResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["requestCid"] = model.NestedToDAMLValue(t.RequestCid)

	return m
}

func (t CredentialBillingRequestToAdjustBillingParamsResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingRequestToAdjustBillingParamsResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingRequestToAdjustBillingParamsResult to hex string (Canton MCMS format)
func (t CredentialBillingRequestToAdjustBillingParamsResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingRequestToAdjustBillingParamsResult from hex string (Canton MCMS format)
func (t *CredentialBillingRequestToAdjustBillingParamsResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingTopUp is a Record type
type CredentialBillingTopUp struct {
	AmountUsd          types.NUMERIC       `json:"amountUsd"`
	CoinCids           []types.CONTRACT_ID `json:"coinCids"`
	AppTransferContext AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts CredentialBillingTopUp to a map for DAML arguments
func (t CredentialBillingTopUp) ToMap() map[string]any {
	m := make(map[string]any)

	m["amountUsd"] = t.AmountUsd

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t CredentialBillingTopUp) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingTopUp) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingTopUp to hex string (Canton MCMS format)
func (t CredentialBillingTopUp) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingTopUp from hex string (Canton MCMS format)
func (t *CredentialBillingTopUp) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialBillingTopUpResult is a Record type
type CredentialBillingTopUpResult struct {
	NewCredentialBillingCid types.CONTRACT_ID `json:"newCredentialBillingCid"`
}

// ToMap converts CredentialBillingTopUpResult to a map for DAML arguments
func (t CredentialBillingTopUpResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["newCredentialBillingCid"] = model.NestedToDAMLValue(t.NewCredentialBillingCid)

	return m
}

func (t CredentialBillingTopUpResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialBillingTopUpResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialBillingTopUpResult to hex string (Canton MCMS format)
func (t CredentialBillingTopUpResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialBillingTopUpResult from hex string (Canton MCMS format)
func (t *CredentialBillingTopUpResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialOffer is a Template type
type CredentialOffer struct {
	Operator                types.PARTY           `json:"operator"`
	Issuer                  types.PARTY           `json:"issuer"`
	Holder                  types.PARTY           `json:"holder"`
	Dso                     types.PARTY           `json:"dso"`
	Id                      types.TEXT            `json:"id"`
	Description             types.TEXT            `json:"description"`
	Claims                  []credential_v0.Claim `json:"claims"`
	BillingParams           *BillingParams        `json:"billingParams" hex:"optional"`
	DepositInitialAmountUsd *types.NUMERIC        `json:"depositInitialAmountUsd" hex:"optional"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t CredentialOffer) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t CredentialOffer) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t CredentialOffer) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["issuer"] = t.Issuer.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["holder"] = t.Holder.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["id"] = string(t.Id)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["description"] = string(t.Description)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["claims"] = func() []any {
		res := make([]any, 0, len(t.Claims))
		for _, e := range t.Claims {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	if t.BillingParams != nil {
		args["billingParams"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.BillingParams),
		}
	} else {
		args["billingParams"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.DepositInitialAmountUsd != nil {
		args["depositInitialAmountUsd"] = map[string]any{
			"_type": "optional",
			"value": *t.DepositInitialAmountUsd,
		}
	} else {
		args["depositInitialAmountUsd"] = map[string]any{
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
func (t CredentialOffer) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["issuer"] = t.Issuer.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["holder"] = t.Holder.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["id"] = string(t.Id)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["description"] = string(t.Description)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["claims"] = func() []any {
		res := make([]any, 0, len(t.Claims))
		for _, e := range t.Claims {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	if t.BillingParams != nil {
		args["billingParams"] = map[string]any{
			"_type": "optional",
			"value": model.NestedToDAMLValue(*t.BillingParams),
		}
	} else {
		args["billingParams"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	if t.DepositInitialAmountUsd != nil {
		args["depositInitialAmountUsd"] = map[string]any{
			"_type": "optional",
			"value": *t.DepositInitialAmountUsd,
		}
	} else {
		args["depositInitialAmountUsd"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t CredentialOffer) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOffer) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOffer to hex string (Canton MCMS format)
func (t CredentialOffer) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOffer from hex string (Canton MCMS format)
func (t *CredentialOffer) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for CredentialOffer

// Archive exercises the Archive choice on this CredentialOffer contract
// This method uses the package name in the template ID
func (t CredentialOffer) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t CredentialOffer) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// CredentialOfferAcceptFree exercises the CredentialOffer_AcceptFree choice on this CredentialOffer contract
// This method uses the package name in the template ID
func (t CredentialOffer) CredentialOfferAcceptFree(contractID string, args CredentialOfferAcceptFree) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "CredentialOffer_AcceptFree",
		Arguments:  argsToMap(args),
	}
}

// CredentialOfferAcceptFreeWithPackageID exercises the CredentialOffer_AcceptFree choice using the provided package ID instead of package name
func (t CredentialOffer) CredentialOfferAcceptFreeWithPackageID(contractID string, packageID string, args CredentialOfferAcceptFree) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "CredentialOffer_AcceptFree",
		Arguments:  argsToMap(args),
	}
}

// CredentialOfferAcceptPaid exercises the CredentialOffer_AcceptPaid choice on this CredentialOffer contract
// This method uses the package name in the template ID
func (t CredentialOffer) CredentialOfferAcceptPaid(contractID string, args CredentialOfferAcceptPaid) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "CredentialOffer_AcceptPaid",
		Arguments:  argsToMap(args),
	}
}

// CredentialOfferAcceptPaidWithPackageID exercises the CredentialOffer_AcceptPaid choice using the provided package ID instead of package name
func (t CredentialOffer) CredentialOfferAcceptPaidWithPackageID(contractID string, packageID string, args CredentialOfferAcceptPaid) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "CredentialOffer_AcceptPaid",
		Arguments:  argsToMap(args),
	}
}

// CredentialOfferCancel exercises the CredentialOffer_Cancel choice on this CredentialOffer contract
// This method uses the package name in the template ID
func (t CredentialOffer) CredentialOfferCancel(contractID string, args CredentialOfferCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "CredentialOffer_Cancel",
		Arguments:  argsToMap(args),
	}
}

// CredentialOfferCancelWithPackageID exercises the CredentialOffer_Cancel choice using the provided package ID instead of package name
func (t CredentialOffer) CredentialOfferCancelWithPackageID(contractID string, packageID string, args CredentialOfferCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "CredentialOffer_Cancel",
		Arguments:  argsToMap(args),
	}
}

// CredentialOfferReject exercises the CredentialOffer_Reject choice on this CredentialOffer contract
// This method uses the package name in the template ID
func (t CredentialOffer) CredentialOfferReject(contractID string, args CredentialOfferReject) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "CredentialOffer_Reject",
		Arguments:  argsToMap(args),
	}
}

// CredentialOfferRejectWithPackageID exercises the CredentialOffer_Reject choice using the provided package ID instead of package name
func (t CredentialOffer) CredentialOfferRejectWithPackageID(contractID string, packageID string, args CredentialOfferReject) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Offer", "CredentialOffer"),
		ContractID: contractID,
		Choice:     "CredentialOffer_Reject",
		Arguments:  argsToMap(args),
	}
}

// CredentialOfferAcceptFree is a Record type
type CredentialOfferAcceptFree struct {
}

// ToMap converts CredentialOfferAcceptFree to a map for DAML arguments
func (t CredentialOfferAcceptFree) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t CredentialOfferAcceptFree) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOfferAcceptFree) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOfferAcceptFree to hex string (Canton MCMS format)
func (t CredentialOfferAcceptFree) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOfferAcceptFree from hex string (Canton MCMS format)
func (t *CredentialOfferAcceptFree) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialOfferAcceptFreeResult is a Record type
type CredentialOfferAcceptFreeResult struct {
	CredentialCid types.CONTRACT_ID `json:"credentialCid"`
}

// ToMap converts CredentialOfferAcceptFreeResult to a map for DAML arguments
func (t CredentialOfferAcceptFreeResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialCid"] = model.NestedToDAMLValue(t.CredentialCid)

	return m
}

func (t CredentialOfferAcceptFreeResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOfferAcceptFreeResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOfferAcceptFreeResult to hex string (Canton MCMS format)
func (t CredentialOfferAcceptFreeResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOfferAcceptFreeResult from hex string (Canton MCMS format)
func (t *CredentialOfferAcceptFreeResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialOfferAcceptPaid is a Record type
type CredentialOfferAcceptPaid struct {
	HolderInputs       []types.CONTRACT_ID `json:"holderInputs"`
	AppTransferContext AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts CredentialOfferAcceptPaid to a map for DAML arguments
func (t CredentialOfferAcceptPaid) ToMap() map[string]any {
	m := make(map[string]any)

	m["holderInputs"] = func() []any {
		res := make([]any, 0, len(t.HolderInputs))
		for _, e := range t.HolderInputs {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t CredentialOfferAcceptPaid) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOfferAcceptPaid) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOfferAcceptPaid to hex string (Canton MCMS format)
func (t CredentialOfferAcceptPaid) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOfferAcceptPaid from hex string (Canton MCMS format)
func (t *CredentialOfferAcceptPaid) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialOfferAcceptPaidResult is a Record type
type CredentialOfferAcceptPaidResult struct {
	CredentialBillingCid types.CONTRACT_ID `json:"credentialBillingCid"`
	CredentialCid        types.CONTRACT_ID `json:"credentialCid"`
}

// ToMap converts CredentialOfferAcceptPaidResult to a map for DAML arguments
func (t CredentialOfferAcceptPaidResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["credentialCid"] = model.NestedToDAMLValue(t.CredentialCid)

	return m
}

func (t CredentialOfferAcceptPaidResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOfferAcceptPaidResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOfferAcceptPaidResult to hex string (Canton MCMS format)
func (t CredentialOfferAcceptPaidResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOfferAcceptPaidResult from hex string (Canton MCMS format)
func (t *CredentialOfferAcceptPaidResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialOfferCancel is a Record type
type CredentialOfferCancel struct {
}

// ToMap converts CredentialOfferCancel to a map for DAML arguments
func (t CredentialOfferCancel) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t CredentialOfferCancel) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOfferCancel) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOfferCancel to hex string (Canton MCMS format)
func (t CredentialOfferCancel) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOfferCancel from hex string (Canton MCMS format)
func (t *CredentialOfferCancel) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialOfferCancelResult is a Record type
type CredentialOfferCancelResult struct {
}

// ToMap converts CredentialOfferCancelResult to a map for DAML arguments
func (t CredentialOfferCancelResult) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t CredentialOfferCancelResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOfferCancelResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOfferCancelResult to hex string (Canton MCMS format)
func (t CredentialOfferCancelResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOfferCancelResult from hex string (Canton MCMS format)
func (t *CredentialOfferCancelResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialOfferReject is a Record type
type CredentialOfferReject struct {
	Reason types.TEXT `json:"reason"`
}

// ToMap converts CredentialOfferReject to a map for DAML arguments
func (t CredentialOfferReject) ToMap() map[string]any {
	m := make(map[string]any)

	m["reason"] = string(t.Reason)

	return m
}

func (t CredentialOfferReject) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOfferReject) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOfferReject to hex string (Canton MCMS format)
func (t CredentialOfferReject) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOfferReject from hex string (Canton MCMS format)
func (t *CredentialOfferReject) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// CredentialOfferRejectResult is a Record type
type CredentialOfferRejectResult struct {
	RejectedCredentialOfferCid types.CONTRACT_ID `json:"rejectedCredentialOfferCid"`
}

// ToMap converts CredentialOfferRejectResult to a map for DAML arguments
func (t CredentialOfferRejectResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["rejectedCredentialOfferCid"] = model.NestedToDAMLValue(t.RejectedCredentialOfferCid)

	return m
}

func (t CredentialOfferRejectResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *CredentialOfferRejectResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes CredentialOfferRejectResult to hex string (Canton MCMS format)
func (t CredentialOfferRejectResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes CredentialOfferRejectResult from hex string (Canton MCMS format)
func (t *CredentialOfferRejectResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// DistributionSlice is a Record type
type DistributionSlice struct {
	CredentialBillingCid types.CONTRACT_ID `json:"credentialBillingCid"`
	Percentage           types.NUMERIC     `json:"percentage"`
}

// ToMap converts DistributionSlice to a map for DAML arguments
func (t DistributionSlice) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["percentage"] = t.Percentage

	return m
}

func (t DistributionSlice) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *DistributionSlice) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes DistributionSlice to hex string (Canton MCMS format)
func (t DistributionSlice) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes DistributionSlice from hex string (Canton MCMS format)
func (t *DistributionSlice) UnmarshalHex(data string) error {
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

// FeeRecord is a Template type
type FeeRecord struct {
	Operator                     types.PARTY   `json:"operator"`
	Provider                     types.PARTY   `json:"provider"`
	User                         types.PARTY   `json:"user"`
	Dso                          types.PARTY   `json:"dso"`
	CcFeesBurned                 types.NUMERIC `json:"ccFeesBurned"`
	ExtraFeaturedAppCcFeesBurned types.NUMERIC `json:"extraFeaturedAppCcFeesBurned"`
	IsFeatured                   types.BOOL    `json:"isFeatured"`
	Round                        Round         `json:"round"`
	Reference                    types.TEXT    `json:"reference"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t FeeRecord) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Accounting", "FeeRecord")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t FeeRecord) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Accounting", "FeeRecord")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t FeeRecord) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["provider"] = t.Provider.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	if t.CcFeesBurned != "" {
		args["ccFeesBurned"] = t.CcFeesBurned
	}

	if t.ExtraFeaturedAppCcFeesBurned != "" {
		args["extraFeaturedAppCcFeesBurned"] = t.ExtraFeaturedAppCcFeesBurned
	}

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["isFeatured"] = bool(t.IsFeatured)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["round"] = model.NestedToDAMLValue(t.Round)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["reference"] = string(t.Reference)

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t FeeRecord) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["provider"] = t.Provider.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	if t.CcFeesBurned != "" {
		args["ccFeesBurned"] = t.CcFeesBurned
	}

	if t.ExtraFeaturedAppCcFeesBurned != "" {
		args["extraFeaturedAppCcFeesBurned"] = t.ExtraFeaturedAppCcFeesBurned
	}

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["isFeatured"] = bool(t.IsFeatured)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["round"] = model.NestedToDAMLValue(t.Round)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["reference"] = string(t.Reference)

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t FeeRecord) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *FeeRecord) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes FeeRecord to hex string (Canton MCMS format)
func (t FeeRecord) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes FeeRecord from hex string (Canton MCMS format)
func (t *FeeRecord) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for FeeRecord

// FeeRecordCalculateReward exercises the FeeRecord_CalculateReward choice on this FeeRecord contract
// This method uses the package name in the template ID
func (t FeeRecord) FeeRecordCalculateReward(contractID string, args FeeRecordCalculateReward) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Accounting", "FeeRecord"),
		ContractID: contractID,
		Choice:     "FeeRecord_CalculateReward",
		Arguments:  argsToMap(args),
	}
}

// FeeRecordCalculateRewardWithPackageID exercises the FeeRecord_CalculateReward choice using the provided package ID instead of package name
func (t FeeRecord) FeeRecordCalculateRewardWithPackageID(contractID string, packageID string, args FeeRecordCalculateReward) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Accounting", "FeeRecord"),
		ContractID: contractID,
		Choice:     "FeeRecord_CalculateReward",
		Arguments:  argsToMap(args),
	}
}

// FeeRecordArchive exercises the FeeRecord_Archive choice on this FeeRecord contract
// This method uses the package name in the template ID
func (t FeeRecord) FeeRecordArchive(contractID string, args FeeRecordArchive) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Accounting", "FeeRecord"),
		ContractID: contractID,
		Choice:     "FeeRecord_Archive",
		Arguments:  argsToMap(args),
	}
}

// FeeRecordArchiveWithPackageID exercises the FeeRecord_Archive choice using the provided package ID instead of package name
func (t FeeRecord) FeeRecordArchiveWithPackageID(contractID string, packageID string, args FeeRecordArchive) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Accounting", "FeeRecord"),
		ContractID: contractID,
		Choice:     "FeeRecord_Archive",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this FeeRecord contract
// This method uses the package name in the template ID
func (t FeeRecord) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Accounting", "FeeRecord"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t FeeRecord) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Accounting", "FeeRecord"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// FeeRecordArchive is a Record type
type FeeRecordArchive struct {
}

// ToMap converts FeeRecordArchive to a map for DAML arguments
func (t FeeRecordArchive) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t FeeRecordArchive) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *FeeRecordArchive) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes FeeRecordArchive to hex string (Canton MCMS format)
func (t FeeRecordArchive) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes FeeRecordArchive from hex string (Canton MCMS format)
func (t *FeeRecordArchive) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// FeeRecordArchiveResult is a Record type
type FeeRecordArchiveResult struct {
}

// ToMap converts FeeRecordArchiveResult to a map for DAML arguments
func (t FeeRecordArchiveResult) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t FeeRecordArchiveResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *FeeRecordArchiveResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes FeeRecordArchiveResult to hex string (Canton MCMS format)
func (t FeeRecordArchiveResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes FeeRecordArchiveResult from hex string (Canton MCMS format)
func (t *FeeRecordArchiveResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// FeeRecordCalculateReward is a Record type
type FeeRecordCalculateReward struct {
	IssuingMiningRound IssuingMiningRound `json:"issuingMiningRound"`
}

// ToMap converts FeeRecordCalculateReward to a map for DAML arguments
func (t FeeRecordCalculateReward) ToMap() map[string]any {
	m := make(map[string]any)

	m["issuingMiningRound"] = model.NestedToDAMLValue(t.IssuingMiningRound)

	return m
}

func (t FeeRecordCalculateReward) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *FeeRecordCalculateReward) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes FeeRecordCalculateReward to hex string (Canton MCMS format)
func (t FeeRecordCalculateReward) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes FeeRecordCalculateReward from hex string (Canton MCMS format)
func (t *FeeRecordCalculateReward) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// OperatorService is a Template type
type OperatorService struct {
	Operator types.PARTY `json:"operator"`
	Dso      types.PARTY `json:"dso"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t OperatorService) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.Operator", "OperatorService")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t OperatorService) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.Operator", "OperatorService")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t OperatorService) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t OperatorService) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t OperatorService) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *OperatorService) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes OperatorService to hex string (Canton MCMS format)
func (t OperatorService) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes OperatorService from hex string (Canton MCMS format)
func (t *OperatorService) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for OperatorService

// OperatorServiceAcceptUserServiceRequest exercises the OperatorService_AcceptUserServiceRequest choice on this OperatorService contract
// This method uses the package name in the template ID
func (t OperatorService) OperatorServiceAcceptUserServiceRequest(contractID string, args OperatorServiceAcceptUserServiceRequest) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.Operator", "OperatorService"),
		ContractID: contractID,
		Choice:     "OperatorService_AcceptUserServiceRequest",
		Arguments:  argsToMap(args),
	}
}

// OperatorServiceAcceptUserServiceRequestWithPackageID exercises the OperatorService_AcceptUserServiceRequest choice using the provided package ID instead of package name
func (t OperatorService) OperatorServiceAcceptUserServiceRequestWithPackageID(contractID string, packageID string, args OperatorServiceAcceptUserServiceRequest) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.Operator", "OperatorService"),
		ContractID: contractID,
		Choice:     "OperatorService_AcceptUserServiceRequest",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this OperatorService contract
// This method uses the package name in the template ID
func (t OperatorService) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.Operator", "OperatorService"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t OperatorService) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.Operator", "OperatorService"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// OperatorServiceRejectUserServiceRequest exercises the OperatorService_RejectUserServiceRequest choice on this OperatorService contract
// This method uses the package name in the template ID
func (t OperatorService) OperatorServiceRejectUserServiceRequest(contractID string, args OperatorServiceRejectUserServiceRequest) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.Operator", "OperatorService"),
		ContractID: contractID,
		Choice:     "OperatorService_RejectUserServiceRequest",
		Arguments:  argsToMap(args),
	}
}

// OperatorServiceRejectUserServiceRequestWithPackageID exercises the OperatorService_RejectUserServiceRequest choice using the provided package ID instead of package name
func (t OperatorService) OperatorServiceRejectUserServiceRequestWithPackageID(contractID string, packageID string, args OperatorServiceRejectUserServiceRequest) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.Operator", "OperatorService"),
		ContractID: contractID,
		Choice:     "OperatorService_RejectUserServiceRequest",
		Arguments:  argsToMap(args),
	}
}

// OperatorServiceAcceptUserServiceRequest is a Record type
type OperatorServiceAcceptUserServiceRequest struct {
	UserServiceRequestCid types.CONTRACT_ID `json:"userServiceRequestCid"`
}

// ToMap converts OperatorServiceAcceptUserServiceRequest to a map for DAML arguments
func (t OperatorServiceAcceptUserServiceRequest) ToMap() map[string]any {
	m := make(map[string]any)

	m["userServiceRequestCid"] = model.NestedToDAMLValue(t.UserServiceRequestCid)

	return m
}

func (t OperatorServiceAcceptUserServiceRequest) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *OperatorServiceAcceptUserServiceRequest) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes OperatorServiceAcceptUserServiceRequest to hex string (Canton MCMS format)
func (t OperatorServiceAcceptUserServiceRequest) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes OperatorServiceAcceptUserServiceRequest from hex string (Canton MCMS format)
func (t *OperatorServiceAcceptUserServiceRequest) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// OperatorServiceRejectUserServiceRequest is a Record type
type OperatorServiceRejectUserServiceRequest struct {
	UserServiceRequestCid types.CONTRACT_ID `json:"userServiceRequestCid"`
}

// ToMap converts OperatorServiceRejectUserServiceRequest to a map for DAML arguments
func (t OperatorServiceRejectUserServiceRequest) ToMap() map[string]any {
	m := make(map[string]any)

	m["userServiceRequestCid"] = model.NestedToDAMLValue(t.UserServiceRequestCid)

	return m
}

func (t OperatorServiceRejectUserServiceRequest) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *OperatorServiceRejectUserServiceRequest) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes OperatorServiceRejectUserServiceRequest to hex string (Canton MCMS format)
func (t OperatorServiceRejectUserServiceRequest) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes OperatorServiceRejectUserServiceRequest from hex string (Canton MCMS format)
func (t *OperatorServiceRejectUserServiceRequest) UnmarshalHex(data string) error {
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

// RejectedCredentialOffer is a Template type
type RejectedCredentialOffer struct {
	Offer  CredentialOffer `json:"offer"`
	Reason types.TEXT      `json:"reason"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t RejectedCredentialOffer) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Offer", "RejectedCredentialOffer")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t RejectedCredentialOffer) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Offer", "RejectedCredentialOffer")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t RejectedCredentialOffer) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["offer"] = model.NestedToDAMLValue(t.Offer)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["reason"] = string(t.Reason)

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t RejectedCredentialOffer) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["offer"] = model.NestedToDAMLValue(t.Offer)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["reason"] = string(t.Reason)

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t RejectedCredentialOffer) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *RejectedCredentialOffer) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes RejectedCredentialOffer to hex string (Canton MCMS format)
func (t RejectedCredentialOffer) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes RejectedCredentialOffer from hex string (Canton MCMS format)
func (t *RejectedCredentialOffer) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for RejectedCredentialOffer

// Archive exercises the Archive choice on this RejectedCredentialOffer contract
// This method uses the package name in the template ID
func (t RejectedCredentialOffer) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Offer", "RejectedCredentialOffer"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t RejectedCredentialOffer) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Offer", "RejectedCredentialOffer"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// RewardRecord is a Template type
type RewardRecord struct {
	Operator        types.PARTY   `json:"operator"`
	Provider        types.PARTY   `json:"provider"`
	User            types.PARTY   `json:"user"`
	CcRewardsEarned types.NUMERIC `json:"ccRewardsEarned"`
	Round           Round         `json:"round"`
	Reference       types.TEXT    `json:"reference"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t RewardRecord) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Accounting", "RewardRecord")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t RewardRecord) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Accounting", "RewardRecord")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t RewardRecord) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["provider"] = t.Provider.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	if t.CcRewardsEarned != "" {
		args["ccRewardsEarned"] = t.CcRewardsEarned
	}

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["round"] = model.NestedToDAMLValue(t.Round)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["reference"] = string(t.Reference)

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t RewardRecord) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["provider"] = t.Provider.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	if t.CcRewardsEarned != "" {
		args["ccRewardsEarned"] = t.CcRewardsEarned
	}

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["round"] = model.NestedToDAMLValue(t.Round)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["reference"] = string(t.Reference)

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t RewardRecord) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *RewardRecord) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes RewardRecord to hex string (Canton MCMS format)
func (t RewardRecord) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes RewardRecord from hex string (Canton MCMS format)
func (t *RewardRecord) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for RewardRecord

// Archive exercises the Archive choice on this RewardRecord contract
// This method uses the package name in the template ID
func (t RewardRecord) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Accounting", "RewardRecord"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t RewardRecord) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Accounting", "RewardRecord"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// RewardRecordArchive exercises the RewardRecord_Archive choice on this RewardRecord contract
// This method uses the package name in the template ID
func (t RewardRecord) RewardRecordArchive(contractID string, args RewardRecordArchive) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Model.Accounting", "RewardRecord"),
		ContractID: contractID,
		Choice:     "RewardRecord_Archive",
		Arguments:  argsToMap(args),
	}
}

// RewardRecordArchiveWithPackageID exercises the RewardRecord_Archive choice using the provided package ID instead of package name
func (t RewardRecord) RewardRecordArchiveWithPackageID(contractID string, packageID string, args RewardRecordArchive) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Model.Accounting", "RewardRecord"),
		ContractID: contractID,
		Choice:     "RewardRecord_Archive",
		Arguments:  argsToMap(args),
	}
}

// RewardRecordArchive is a Record type
type RewardRecordArchive struct {
}

// ToMap converts RewardRecordArchive to a map for DAML arguments
func (t RewardRecordArchive) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t RewardRecordArchive) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *RewardRecordArchive) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes RewardRecordArchive to hex string (Canton MCMS format)
func (t RewardRecordArchive) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes RewardRecordArchive from hex string (Canton MCMS format)
func (t *RewardRecordArchive) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// RewardRecordArchiveResult is a Record type
type RewardRecordArchiveResult struct {
}

// ToMap converts RewardRecordArchiveResult to a map for DAML arguments
func (t RewardRecordArchiveResult) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t RewardRecordArchiveResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *RewardRecordArchiveResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes RewardRecordArchiveResult to hex string (Canton MCMS format)
func (t RewardRecordArchiveResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes RewardRecordArchiveResult from hex string (Canton MCMS format)
func (t *RewardRecordArchiveResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserService is a Template type
type UserService struct {
	Operator types.PARTY `json:"operator"`
	User     types.PARTY `json:"user"`
	Dso      types.PARTY `json:"dso"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t UserService) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t UserService) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t UserService) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t UserService) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["dso"] = t.Dso.ToMap()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t UserService) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserService) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserService to hex string (Canton MCMS format)
func (t UserService) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserService from hex string (Canton MCMS format)
func (t *UserService) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for UserService

// UserServiceOfferPaidCredential exercises the UserService_OfferPaidCredential choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceOfferPaidCredential(contractID string, args UserServiceOfferPaidCredential) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_OfferPaidCredential",
		Arguments:  argsToMap(args),
	}
}

// UserServiceOfferPaidCredentialWithPackageID exercises the UserService_OfferPaidCredential choice using the provided package ID instead of package name
func (t UserService) UserServiceOfferPaidCredentialWithPackageID(contractID string, packageID string, args UserServiceOfferPaidCredential) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_OfferPaidCredential",
		Arguments:  argsToMap(args),
	}
}

// UserServiceOfferFreeCredential exercises the UserService_OfferFreeCredential choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceOfferFreeCredential(contractID string, args UserServiceOfferFreeCredential) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_OfferFreeCredential",
		Arguments:  argsToMap(args),
	}
}

// UserServiceOfferFreeCredentialWithPackageID exercises the UserService_OfferFreeCredential choice using the provided package ID instead of package name
func (t UserService) UserServiceOfferFreeCredentialWithPackageID(contractID string, packageID string, args UserServiceOfferFreeCredential) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_OfferFreeCredential",
		Arguments:  argsToMap(args),
	}
}

// UserServiceAcceptPaidCredentialOffer exercises the UserService_AcceptPaidCredentialOffer choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceAcceptPaidCredentialOffer(contractID string, args UserServiceAcceptPaidCredentialOffer) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_AcceptPaidCredentialOffer",
		Arguments:  argsToMap(args),
	}
}

// UserServiceAcceptPaidCredentialOfferWithPackageID exercises the UserService_AcceptPaidCredentialOffer choice using the provided package ID instead of package name
func (t UserService) UserServiceAcceptPaidCredentialOfferWithPackageID(contractID string, packageID string, args UserServiceAcceptPaidCredentialOffer) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_AcceptPaidCredentialOffer",
		Arguments:  argsToMap(args),
	}
}

// UserServiceAcceptFreeCredentialOffer exercises the UserService_AcceptFreeCredentialOffer choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceAcceptFreeCredentialOffer(contractID string, args UserServiceAcceptFreeCredentialOffer) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_AcceptFreeCredentialOffer",
		Arguments:  argsToMap(args),
	}
}

// UserServiceAcceptFreeCredentialOfferWithPackageID exercises the UserService_AcceptFreeCredentialOffer choice using the provided package ID instead of package name
func (t UserService) UserServiceAcceptFreeCredentialOfferWithPackageID(contractID string, packageID string, args UserServiceAcceptFreeCredentialOffer) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_AcceptFreeCredentialOffer",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRejectCredentialOffer exercises the UserService_RejectCredentialOffer choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceRejectCredentialOffer(contractID string, args UserServiceRejectCredentialOffer) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_RejectCredentialOffer",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRejectCredentialOfferWithPackageID exercises the UserService_RejectCredentialOffer choice using the provided package ID instead of package name
func (t UserService) UserServiceRejectCredentialOfferWithPackageID(contractID string, packageID string, args UserServiceRejectCredentialOffer) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_RejectCredentialOffer",
		Arguments:  argsToMap(args),
	}
}

// UserServiceCancelCredentialOffer exercises the UserService_CancelCredentialOffer choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceCancelCredentialOffer(contractID string, args UserServiceCancelCredentialOffer) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_CancelCredentialOffer",
		Arguments:  argsToMap(args),
	}
}

// UserServiceCancelCredentialOfferWithPackageID exercises the UserService_CancelCredentialOffer choice using the provided package ID instead of package name
func (t UserService) UserServiceCancelCredentialOfferWithPackageID(contractID string, packageID string, args UserServiceCancelCredentialOffer) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_CancelCredentialOffer",
		Arguments:  argsToMap(args),
	}
}

// UserServiceDistribute exercises the UserService_Distribute choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceDistribute(contractID string, args UserServiceDistribute) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_Distribute",
		Arguments:  argsToMap(args),
	}
}

// UserServiceDistributeWithPackageID exercises the UserService_Distribute choice using the provided package ID instead of package name
func (t UserService) UserServiceDistributeWithPackageID(contractID string, packageID string, args UserServiceDistribute) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_Distribute",
		Arguments:  argsToMap(args),
	}
}

// UserServiceDistributeMulti exercises the UserService_DistributeMulti choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceDistributeMulti(contractID string, args UserServiceDistributeMulti) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_DistributeMulti",
		Arguments:  argsToMap(args),
	}
}

// UserServiceDistributeMultiWithPackageID exercises the UserService_DistributeMulti choice using the provided package ID instead of package name
func (t UserService) UserServiceDistributeMultiWithPackageID(contractID string, packageID string, args UserServiceDistributeMulti) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_DistributeMulti",
		Arguments:  argsToMap(args),
	}
}

// UserServiceDistributeAndAdjustDeposit exercises the UserService_DistributeAndAdjustDeposit choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceDistributeAndAdjustDeposit(contractID string, args UserServiceDistributeAndAdjustDeposit) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_DistributeAndAdjustDeposit",
		Arguments:  argsToMap(args),
	}
}

// UserServiceDistributeAndAdjustDepositWithPackageID exercises the UserService_DistributeAndAdjustDeposit choice using the provided package ID instead of package name
func (t UserService) UserServiceDistributeAndAdjustDepositWithPackageID(contractID string, packageID string, args UserServiceDistributeAndAdjustDeposit) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_DistributeAndAdjustDeposit",
		Arguments:  argsToMap(args),
	}
}

// UserServiceDistributeAndAdjustDepositMulti exercises the UserService_DistributeAndAdjustDepositMulti choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceDistributeAndAdjustDepositMulti(contractID string, args UserServiceDistributeAndAdjustDepositMulti) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_DistributeAndAdjustDepositMulti",
		Arguments:  argsToMap(args),
	}
}

// UserServiceDistributeAndAdjustDepositMultiWithPackageID exercises the UserService_DistributeAndAdjustDepositMulti choice using the provided package ID instead of package name
func (t UserService) UserServiceDistributeAndAdjustDepositMultiWithPackageID(contractID string, packageID string, args UserServiceDistributeAndAdjustDepositMulti) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_DistributeAndAdjustDepositMulti",
		Arguments:  argsToMap(args),
	}
}

// UserServiceTopUp exercises the UserService_TopUp choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceTopUp(contractID string, args UserServiceTopUp) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_TopUp",
		Arguments:  argsToMap(args),
	}
}

// UserServiceTopUpWithPackageID exercises the UserService_TopUp choice using the provided package ID instead of package name
func (t UserService) UserServiceTopUpWithPackageID(contractID string, packageID string, args UserServiceTopUp) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_TopUp",
		Arguments:  argsToMap(args),
	}
}

// UserServiceAdjustBillingParams exercises the UserService_AdjustBillingParams choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceAdjustBillingParams(contractID string, args UserServiceAdjustBillingParams) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_AdjustBillingParams",
		Arguments:  argsToMap(args),
	}
}

// UserServiceAdjustBillingParamsWithPackageID exercises the UserService_AdjustBillingParams choice using the provided package ID instead of package name
func (t UserService) UserServiceAdjustBillingParamsWithPackageID(contractID string, packageID string, args UserServiceAdjustBillingParams) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_AdjustBillingParams",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRequestToAdjustBillingParams exercises the UserService_RequestToAdjustBillingParams choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceRequestToAdjustBillingParams(contractID string, args UserServiceRequestToAdjustBillingParams) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_RequestToAdjustBillingParams",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRequestToAdjustBillingParamsWithPackageID exercises the UserService_RequestToAdjustBillingParams choice using the provided package ID instead of package name
func (t UserService) UserServiceRequestToAdjustBillingParamsWithPackageID(contractID string, packageID string, args UserServiceRequestToAdjustBillingParams) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_RequestToAdjustBillingParams",
		Arguments:  argsToMap(args),
	}
}

// UserServiceBillingParamsAdjustmentRequestAccept exercises the UserService_BillingParamsAdjustmentRequest_Accept choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceBillingParamsAdjustmentRequestAccept(contractID string, args UserServiceBillingParamsAdjustmentRequestAccept) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_BillingParamsAdjustmentRequest_Accept",
		Arguments:  argsToMap(args),
	}
}

// UserServiceBillingParamsAdjustmentRequestAcceptWithPackageID exercises the UserService_BillingParamsAdjustmentRequest_Accept choice using the provided package ID instead of package name
func (t UserService) UserServiceBillingParamsAdjustmentRequestAcceptWithPackageID(contractID string, packageID string, args UserServiceBillingParamsAdjustmentRequestAccept) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_BillingParamsAdjustmentRequest_Accept",
		Arguments:  argsToMap(args),
	}
}

// UserServiceBillingParamsAdjustmentRequestCancel exercises the UserService_BillingParamsAdjustmentRequest_Cancel choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceBillingParamsAdjustmentRequestCancel(contractID string, args UserServiceBillingParamsAdjustmentRequestCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_BillingParamsAdjustmentRequest_Cancel",
		Arguments:  argsToMap(args),
	}
}

// UserServiceBillingParamsAdjustmentRequestCancelWithPackageID exercises the UserService_BillingParamsAdjustmentRequest_Cancel choice using the provided package ID instead of package name
func (t UserService) UserServiceBillingParamsAdjustmentRequestCancelWithPackageID(contractID string, packageID string, args UserServiceBillingParamsAdjustmentRequestCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_BillingParamsAdjustmentRequest_Cancel",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRevokeCredential exercises the UserService_RevokeCredential choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceRevokeCredential(contractID string, args UserServiceRevokeCredential) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_RevokeCredential",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRevokeCredentialWithPackageID exercises the UserService_RevokeCredential choice using the provided package ID instead of package name
func (t UserService) UserServiceRevokeCredentialWithPackageID(contractID string, packageID string, args UserServiceRevokeCredential) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_RevokeCredential",
		Arguments:  argsToMap(args),
	}
}

// UserServiceCancelCredentialBilling exercises the UserService_CancelCredentialBilling choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceCancelCredentialBilling(contractID string, args UserServiceCancelCredentialBilling) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_CancelCredentialBilling",
		Arguments:  argsToMap(args),
	}
}

// UserServiceCancelCredentialBillingWithPackageID exercises the UserService_CancelCredentialBilling choice using the provided package ID instead of package name
func (t UserService) UserServiceCancelCredentialBillingWithPackageID(contractID string, packageID string, args UserServiceCancelCredentialBilling) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_CancelCredentialBilling",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRevokeCredentialAndCancelBilling exercises the UserService_RevokeCredentialAndCancelBilling choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceRevokeCredentialAndCancelBilling(contractID string, args UserServiceRevokeCredentialAndCancelBilling) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_RevokeCredentialAndCancelBilling",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRevokeCredentialAndCancelBillingWithPackageID exercises the UserService_RevokeCredentialAndCancelBilling choice using the provided package ID instead of package name
func (t UserService) UserServiceRevokeCredentialAndCancelBillingWithPackageID(contractID string, packageID string, args UserServiceRevokeCredentialAndCancelBilling) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_RevokeCredentialAndCancelBilling",
		Arguments:  argsToMap(args),
	}
}

// UserServiceTerminate exercises the UserService_Terminate choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) UserServiceTerminate(contractID string, args UserServiceTerminate) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_Terminate",
		Arguments:  argsToMap(args),
	}
}

// UserServiceTerminateWithPackageID exercises the UserService_Terminate choice using the provided package ID instead of package name
func (t UserService) UserServiceTerminateWithPackageID(contractID string, packageID string, args UserServiceTerminate) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "UserService_Terminate",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this UserService contract
// This method uses the package name in the template ID
func (t UserService) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t UserService) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserService"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// UserServiceRequest is a Template type
type UserServiceRequest struct {
	Operator types.PARTY `json:"operator"`
	User     types.PARTY `json:"user"`
}

// GetTemplateID returns the template ID for this template using the package name
func (t UserServiceRequest) GetTemplateID() string {
	return fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserServiceRequest")
}

// GetTemplateIDWithPackageID returns the template ID using the provided package ID instead of package name
func (t UserServiceRequest) GetTemplateIDWithPackageID(packageID string) string {
	return fmt.Sprintf("%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserServiceRequest")
}

// CreateCommand returns a CreateCommand for this template using the package name
func (t UserServiceRequest) CreateCommand() *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateID(),
		Arguments:  args,
	}
}

// CreateCommandWithPackageID returns a CreateCommand using the provided package ID instead of package name
func (t UserServiceRequest) CreateCommandWithPackageID(packageID string) *model.CreateCommand {
	args := make(map[string]any)

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["operator"] = t.Operator.ToMap()

	// IMPORTANT: always include non-optional fields (GENMAP/MAP/LIST/[] etc), even if empty
	args["user"] = t.User.ToMap()

	return &model.CreateCommand{
		TemplateID: t.GetTemplateIDWithPackageID(packageID),
		Arguments:  args,
	}
}

func (t UserServiceRequest) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRequest) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRequest to hex string (Canton MCMS format)
func (t UserServiceRequest) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRequest from hex string (Canton MCMS format)
func (t *UserServiceRequest) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// Choice methods for UserServiceRequest

// UserServiceRequestAccept exercises the UserServiceRequest_Accept choice on this UserServiceRequest contract
// This method uses the package name in the template ID
func (t UserServiceRequest) UserServiceRequestAccept(contractID string, args UserServiceRequestAccept) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserServiceRequest"),
		ContractID: contractID,
		Choice:     "UserServiceRequest_Accept",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRequestAcceptWithPackageID exercises the UserServiceRequest_Accept choice using the provided package ID instead of package name
func (t UserServiceRequest) UserServiceRequestAcceptWithPackageID(contractID string, packageID string, args UserServiceRequestAccept) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserServiceRequest"),
		ContractID: contractID,
		Choice:     "UserServiceRequest_Accept",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRequestCancel exercises the UserServiceRequest_Cancel choice on this UserServiceRequest contract
// This method uses the package name in the template ID
func (t UserServiceRequest) UserServiceRequestCancel(contractID string, args UserServiceRequestCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserServiceRequest"),
		ContractID: contractID,
		Choice:     "UserServiceRequest_Cancel",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRequestCancelWithPackageID exercises the UserServiceRequest_Cancel choice using the provided package ID instead of package name
func (t UserServiceRequest) UserServiceRequestCancelWithPackageID(contractID string, packageID string, args UserServiceRequestCancel) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserServiceRequest"),
		ContractID: contractID,
		Choice:     "UserServiceRequest_Cancel",
		Arguments:  argsToMap(args),
	}
}

// Archive exercises the Archive choice on this UserServiceRequest contract
// This method uses the package name in the template ID
func (t UserServiceRequest) Archive(contractID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserServiceRequest"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// ArchiveWithPackageID exercises the Archive choice using the provided package ID instead of package name
func (t UserServiceRequest) ArchiveWithPackageID(contractID string, packageID string) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserServiceRequest"),
		ContractID: contractID,
		Choice:     "Archive",
		Arguments:  map[string]any{},
	}
}

// UserServiceRequestReject exercises the UserServiceRequest_Reject choice on this UserServiceRequest contract
// This method uses the package name in the template ID
func (t UserServiceRequest) UserServiceRequestReject(contractID string, args UserServiceRequestReject) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", PackageName, "Utility.Credential.App.V0.Service.User", "UserServiceRequest"),
		ContractID: contractID,
		Choice:     "UserServiceRequest_Reject",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRequestRejectWithPackageID exercises the UserServiceRequest_Reject choice using the provided package ID instead of package name
func (t UserServiceRequest) UserServiceRequestRejectWithPackageID(contractID string, packageID string, args UserServiceRequestReject) *model.ExerciseCommand {
	return &model.ExerciseCommand{
		TemplateID: fmt.Sprintf("#%s:%s:%s", packageID, "Utility.Credential.App.V0.Service.User", "UserServiceRequest"),
		ContractID: contractID,
		Choice:     "UserServiceRequest_Reject",
		Arguments:  argsToMap(args),
	}
}

// UserServiceRequestAccept is a Record type
type UserServiceRequestAccept struct {
	Dso types.PARTY `json:"dso"`
}

// ToMap converts UserServiceRequestAccept to a map for DAML arguments
func (t UserServiceRequestAccept) ToMap() map[string]any {
	m := make(map[string]any)

	m["dso"] = t.Dso.ToMap()

	return m
}

func (t UserServiceRequestAccept) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRequestAccept) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRequestAccept to hex string (Canton MCMS format)
func (t UserServiceRequestAccept) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRequestAccept from hex string (Canton MCMS format)
func (t *UserServiceRequestAccept) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRequestAcceptResult is a Record type
type UserServiceRequestAcceptResult struct {
	UserServiceCid types.CONTRACT_ID `json:"userServiceCid"`
}

// ToMap converts UserServiceRequestAcceptResult to a map for DAML arguments
func (t UserServiceRequestAcceptResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["userServiceCid"] = model.NestedToDAMLValue(t.UserServiceCid)

	return m
}

func (t UserServiceRequestAcceptResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRequestAcceptResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRequestAcceptResult to hex string (Canton MCMS format)
func (t UserServiceRequestAcceptResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRequestAcceptResult from hex string (Canton MCMS format)
func (t *UserServiceRequestAcceptResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRequestCancel is a Record type
type UserServiceRequestCancel struct {
}

// ToMap converts UserServiceRequestCancel to a map for DAML arguments
func (t UserServiceRequestCancel) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t UserServiceRequestCancel) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRequestCancel) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRequestCancel to hex string (Canton MCMS format)
func (t UserServiceRequestCancel) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRequestCancel from hex string (Canton MCMS format)
func (t *UserServiceRequestCancel) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRequestCancelResult is a Record type
type UserServiceRequestCancelResult struct {
}

// ToMap converts UserServiceRequestCancelResult to a map for DAML arguments
func (t UserServiceRequestCancelResult) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t UserServiceRequestCancelResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRequestCancelResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRequestCancelResult to hex string (Canton MCMS format)
func (t UserServiceRequestCancelResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRequestCancelResult from hex string (Canton MCMS format)
func (t *UserServiceRequestCancelResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRequestReject is a Record type
type UserServiceRequestReject struct {
}

// ToMap converts UserServiceRequestReject to a map for DAML arguments
func (t UserServiceRequestReject) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t UserServiceRequestReject) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRequestReject) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRequestReject to hex string (Canton MCMS format)
func (t UserServiceRequestReject) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRequestReject from hex string (Canton MCMS format)
func (t *UserServiceRequestReject) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRequestRejectResult is a Record type
type UserServiceRequestRejectResult struct {
}

// ToMap converts UserServiceRequestRejectResult to a map for DAML arguments
func (t UserServiceRequestRejectResult) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t UserServiceRequestRejectResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRequestRejectResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRequestRejectResult to hex string (Canton MCMS format)
func (t UserServiceRequestRejectResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRequestRejectResult from hex string (Canton MCMS format)
func (t *UserServiceRequestRejectResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceAcceptFreeCredentialOffer is a Record type
type UserServiceAcceptFreeCredentialOffer struct {
	CredentialOfferCid types.CONTRACT_ID `json:"credentialOfferCid"`
}

// ToMap converts UserServiceAcceptFreeCredentialOffer to a map for DAML arguments
func (t UserServiceAcceptFreeCredentialOffer) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialOfferCid"] = model.NestedToDAMLValue(t.CredentialOfferCid)

	return m
}

func (t UserServiceAcceptFreeCredentialOffer) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceAcceptFreeCredentialOffer) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceAcceptFreeCredentialOffer to hex string (Canton MCMS format)
func (t UserServiceAcceptFreeCredentialOffer) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceAcceptFreeCredentialOffer from hex string (Canton MCMS format)
func (t *UserServiceAcceptFreeCredentialOffer) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceAcceptPaidCredentialOffer is a Record type
type UserServiceAcceptPaidCredentialOffer struct {
	CredentialOfferCid types.CONTRACT_ID   `json:"credentialOfferCid"`
	DepositAmulets     []types.CONTRACT_ID `json:"depositAmulets"`
	AppTransferContext AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts UserServiceAcceptPaidCredentialOffer to a map for DAML arguments
func (t UserServiceAcceptPaidCredentialOffer) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialOfferCid"] = model.NestedToDAMLValue(t.CredentialOfferCid)

	m["depositAmulets"] = func() []any {
		res := make([]any, 0, len(t.DepositAmulets))
		for _, e := range t.DepositAmulets {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t UserServiceAcceptPaidCredentialOffer) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceAcceptPaidCredentialOffer) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceAcceptPaidCredentialOffer to hex string (Canton MCMS format)
func (t UserServiceAcceptPaidCredentialOffer) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceAcceptPaidCredentialOffer from hex string (Canton MCMS format)
func (t *UserServiceAcceptPaidCredentialOffer) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceAdjustBillingParams is a Record type
type UserServiceAdjustBillingParams struct {
	CredentialBillingCid types.CONTRACT_ID `json:"credentialBillingCid"`
	BillingParams        BillingParams     `json:"billingParams"`
}

// ToMap converts UserServiceAdjustBillingParams to a map for DAML arguments
func (t UserServiceAdjustBillingParams) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["billingParams"] = model.NestedToDAMLValue(t.BillingParams)

	return m
}

func (t UserServiceAdjustBillingParams) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceAdjustBillingParams) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceAdjustBillingParams to hex string (Canton MCMS format)
func (t UserServiceAdjustBillingParams) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceAdjustBillingParams from hex string (Canton MCMS format)
func (t *UserServiceAdjustBillingParams) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceBillingParamsAdjustmentRequestAccept is a Record type
type UserServiceBillingParamsAdjustmentRequestAccept struct {
	RequestCid           types.CONTRACT_ID `json:"requestCid"`
	CredentialBillingCid types.CONTRACT_ID `json:"credentialBillingCid"`
}

// ToMap converts UserServiceBillingParamsAdjustmentRequestAccept to a map for DAML arguments
func (t UserServiceBillingParamsAdjustmentRequestAccept) ToMap() map[string]any {
	m := make(map[string]any)

	m["requestCid"] = model.NestedToDAMLValue(t.RequestCid)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	return m
}

func (t UserServiceBillingParamsAdjustmentRequestAccept) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceBillingParamsAdjustmentRequestAccept) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceBillingParamsAdjustmentRequestAccept to hex string (Canton MCMS format)
func (t UserServiceBillingParamsAdjustmentRequestAccept) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceBillingParamsAdjustmentRequestAccept from hex string (Canton MCMS format)
func (t *UserServiceBillingParamsAdjustmentRequestAccept) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceBillingParamsAdjustmentRequestCancel is a Record type
type UserServiceBillingParamsAdjustmentRequestCancel struct {
	RequestCid types.CONTRACT_ID `json:"requestCid"`
}

// ToMap converts UserServiceBillingParamsAdjustmentRequestCancel to a map for DAML arguments
func (t UserServiceBillingParamsAdjustmentRequestCancel) ToMap() map[string]any {
	m := make(map[string]any)

	m["requestCid"] = model.NestedToDAMLValue(t.RequestCid)

	return m
}

func (t UserServiceBillingParamsAdjustmentRequestCancel) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceBillingParamsAdjustmentRequestCancel) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceBillingParamsAdjustmentRequestCancel to hex string (Canton MCMS format)
func (t UserServiceBillingParamsAdjustmentRequestCancel) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceBillingParamsAdjustmentRequestCancel from hex string (Canton MCMS format)
func (t *UserServiceBillingParamsAdjustmentRequestCancel) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceCancelCredentialBilling is a Record type
type UserServiceCancelCredentialBilling struct {
	CredentialBillingCid types.CONTRACT_ID  `json:"credentialBillingCid"`
	AppTransferContext   AppTransferContext `json:"appTransferContext"`
}

// ToMap converts UserServiceCancelCredentialBilling to a map for DAML arguments
func (t UserServiceCancelCredentialBilling) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t UserServiceCancelCredentialBilling) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceCancelCredentialBilling) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceCancelCredentialBilling to hex string (Canton MCMS format)
func (t UserServiceCancelCredentialBilling) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceCancelCredentialBilling from hex string (Canton MCMS format)
func (t *UserServiceCancelCredentialBilling) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceCancelCredentialOffer is a Record type
type UserServiceCancelCredentialOffer struct {
	CredentialOfferCid types.CONTRACT_ID `json:"credentialOfferCid"`
}

// ToMap converts UserServiceCancelCredentialOffer to a map for DAML arguments
func (t UserServiceCancelCredentialOffer) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialOfferCid"] = model.NestedToDAMLValue(t.CredentialOfferCid)

	return m
}

func (t UserServiceCancelCredentialOffer) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceCancelCredentialOffer) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceCancelCredentialOffer to hex string (Canton MCMS format)
func (t UserServiceCancelCredentialOffer) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceCancelCredentialOffer from hex string (Canton MCMS format)
func (t *UserServiceCancelCredentialOffer) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceDistribute is a Record type
type UserServiceDistribute struct {
	CredentialBillingCid types.CONTRACT_ID   `json:"credentialBillingCid"`
	AmountUsd            types.NUMERIC       `json:"amountUsd"`
	CoinCids             []types.CONTRACT_ID `json:"coinCids"`
	AppTransferContext   AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts UserServiceDistribute to a map for DAML arguments
func (t UserServiceDistribute) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["amountUsd"] = t.AmountUsd

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t UserServiceDistribute) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceDistribute) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceDistribute to hex string (Canton MCMS format)
func (t UserServiceDistribute) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceDistribute from hex string (Canton MCMS format)
func (t *UserServiceDistribute) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceDistributeAndAdjustDeposit is a Record type
type UserServiceDistributeAndAdjustDeposit struct {
	CredentialBillingCid types.CONTRACT_ID   `json:"credentialBillingCid"`
	AmountUsd            types.NUMERIC       `json:"amountUsd"`
	CoinCids             []types.CONTRACT_ID `json:"coinCids"`
	AppTransferContext   AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts UserServiceDistributeAndAdjustDeposit to a map for DAML arguments
func (t UserServiceDistributeAndAdjustDeposit) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["amountUsd"] = t.AmountUsd

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t UserServiceDistributeAndAdjustDeposit) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceDistributeAndAdjustDeposit) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceDistributeAndAdjustDeposit to hex string (Canton MCMS format)
func (t UserServiceDistributeAndAdjustDeposit) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceDistributeAndAdjustDeposit from hex string (Canton MCMS format)
func (t *UserServiceDistributeAndAdjustDeposit) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceDistributeAndAdjustDepositMulti is a Record type
type UserServiceDistributeAndAdjustDepositMulti struct {
	DistributionSlices []DistributionSlice `json:"distributionSlices"`
	AmountUsd          types.NUMERIC       `json:"amountUsd"`
	CoinCids           []types.CONTRACT_ID `json:"coinCids"`
	AppTransferContext AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts UserServiceDistributeAndAdjustDepositMulti to a map for DAML arguments
func (t UserServiceDistributeAndAdjustDepositMulti) ToMap() map[string]any {
	m := make(map[string]any)

	m["distributionSlices"] = func() []any {
		res := make([]any, 0, len(t.DistributionSlices))
		for _, e := range t.DistributionSlices {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	m["amountUsd"] = t.AmountUsd

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t UserServiceDistributeAndAdjustDepositMulti) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceDistributeAndAdjustDepositMulti) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceDistributeAndAdjustDepositMulti to hex string (Canton MCMS format)
func (t UserServiceDistributeAndAdjustDepositMulti) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceDistributeAndAdjustDepositMulti from hex string (Canton MCMS format)
func (t *UserServiceDistributeAndAdjustDepositMulti) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceDistributeAndAdjustDepositMultiResult is a Record type
type UserServiceDistributeAndAdjustDepositMultiResult struct {
	TransferResults []TransferResult `json:"transferResults"`
}

// ToMap converts UserServiceDistributeAndAdjustDepositMultiResult to a map for DAML arguments
func (t UserServiceDistributeAndAdjustDepositMultiResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["transferResults"] = func() []any {
		res := make([]any, 0, len(t.TransferResults))
		for _, e := range t.TransferResults {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	return m
}

func (t UserServiceDistributeAndAdjustDepositMultiResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceDistributeAndAdjustDepositMultiResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceDistributeAndAdjustDepositMultiResult to hex string (Canton MCMS format)
func (t UserServiceDistributeAndAdjustDepositMultiResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceDistributeAndAdjustDepositMultiResult from hex string (Canton MCMS format)
func (t *UserServiceDistributeAndAdjustDepositMultiResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceDistributeMulti is a Record type
type UserServiceDistributeMulti struct {
	DistributionSlices []DistributionSlice `json:"distributionSlices"`
	AmountUsd          types.NUMERIC       `json:"amountUsd"`
	CoinCids           []types.CONTRACT_ID `json:"coinCids"`
	AppTransferContext AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts UserServiceDistributeMulti to a map for DAML arguments
func (t UserServiceDistributeMulti) ToMap() map[string]any {
	m := make(map[string]any)

	m["distributionSlices"] = func() []any {
		res := make([]any, 0, len(t.DistributionSlices))
		for _, e := range t.DistributionSlices {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	m["amountUsd"] = t.AmountUsd

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t UserServiceDistributeMulti) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceDistributeMulti) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceDistributeMulti to hex string (Canton MCMS format)
func (t UserServiceDistributeMulti) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceDistributeMulti from hex string (Canton MCMS format)
func (t *UserServiceDistributeMulti) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceDistributeMultiResult is a Record type
type UserServiceDistributeMultiResult struct {
	TransferResults []TransferResult `json:"transferResults"`
}

// ToMap converts UserServiceDistributeMultiResult to a map for DAML arguments
func (t UserServiceDistributeMultiResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["transferResults"] = func() []any {
		res := make([]any, 0, len(t.TransferResults))
		for _, e := range t.TransferResults {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	return m
}

func (t UserServiceDistributeMultiResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceDistributeMultiResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceDistributeMultiResult to hex string (Canton MCMS format)
func (t UserServiceDistributeMultiResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceDistributeMultiResult from hex string (Canton MCMS format)
func (t *UserServiceDistributeMultiResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceOfferFreeCredential is a Record type
type UserServiceOfferFreeCredential struct {
	Holder      types.PARTY           `json:"holder"`
	Id          types.TEXT            `json:"id"`
	Description types.TEXT            `json:"description"`
	Claims      []credential_v0.Claim `json:"claims"`
}

// ToMap converts UserServiceOfferFreeCredential to a map for DAML arguments
func (t UserServiceOfferFreeCredential) ToMap() map[string]any {
	m := make(map[string]any)

	m["holder"] = t.Holder.ToMap()

	m["id"] = string(t.Id)

	m["description"] = string(t.Description)

	m["claims"] = func() []any {
		res := make([]any, 0, len(t.Claims))
		for _, e := range t.Claims {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	return m
}

func (t UserServiceOfferFreeCredential) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceOfferFreeCredential) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceOfferFreeCredential to hex string (Canton MCMS format)
func (t UserServiceOfferFreeCredential) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceOfferFreeCredential from hex string (Canton MCMS format)
func (t *UserServiceOfferFreeCredential) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceOfferFreeCredentialResult is a Record type
type UserServiceOfferFreeCredentialResult struct {
	CredentialOfferCid types.CONTRACT_ID `json:"credentialOfferCid"`
}

// ToMap converts UserServiceOfferFreeCredentialResult to a map for DAML arguments
func (t UserServiceOfferFreeCredentialResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialOfferCid"] = model.NestedToDAMLValue(t.CredentialOfferCid)

	return m
}

func (t UserServiceOfferFreeCredentialResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceOfferFreeCredentialResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceOfferFreeCredentialResult to hex string (Canton MCMS format)
func (t UserServiceOfferFreeCredentialResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceOfferFreeCredentialResult from hex string (Canton MCMS format)
func (t *UserServiceOfferFreeCredentialResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceOfferPaidCredential is a Record type
type UserServiceOfferPaidCredential struct {
	Holder                  types.PARTY           `json:"holder"`
	Id                      types.TEXT            `json:"id"`
	Description             types.TEXT            `json:"description"`
	Claims                  []credential_v0.Claim `json:"claims"`
	BillingParams           BillingParams         `json:"billingParams"`
	DepositInitialAmountUsd *types.NUMERIC        `json:"depositInitialAmountUsd" hex:"optional"`
}

// ToMap converts UserServiceOfferPaidCredential to a map for DAML arguments
func (t UserServiceOfferPaidCredential) ToMap() map[string]any {
	m := make(map[string]any)

	m["holder"] = t.Holder.ToMap()

	m["id"] = string(t.Id)

	m["description"] = string(t.Description)

	m["claims"] = func() []any {
		res := make([]any, 0, len(t.Claims))
		for _, e := range t.Claims {
			res = append(res, model.NestedToDAMLValue(e))
		}
		return res
	}()

	m["billingParams"] = model.NestedToDAMLValue(t.BillingParams)

	if t.DepositInitialAmountUsd != nil {
		m["depositInitialAmountUsd"] = map[string]any{
			"_type": "optional",
			"value": *t.DepositInitialAmountUsd,
		}
	} else {
		m["depositInitialAmountUsd"] = map[string]any{
			"_type": "optional",
			"value": nil,
		}
	}

	return m
}

func (t UserServiceOfferPaidCredential) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceOfferPaidCredential) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceOfferPaidCredential to hex string (Canton MCMS format)
func (t UserServiceOfferPaidCredential) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceOfferPaidCredential from hex string (Canton MCMS format)
func (t *UserServiceOfferPaidCredential) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceOfferPaidCredentialResult is a Record type
type UserServiceOfferPaidCredentialResult struct {
	CredentialOfferCid types.CONTRACT_ID `json:"credentialOfferCid"`
}

// ToMap converts UserServiceOfferPaidCredentialResult to a map for DAML arguments
func (t UserServiceOfferPaidCredentialResult) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialOfferCid"] = model.NestedToDAMLValue(t.CredentialOfferCid)

	return m
}

func (t UserServiceOfferPaidCredentialResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceOfferPaidCredentialResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceOfferPaidCredentialResult to hex string (Canton MCMS format)
func (t UserServiceOfferPaidCredentialResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceOfferPaidCredentialResult from hex string (Canton MCMS format)
func (t *UserServiceOfferPaidCredentialResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRejectCredentialOffer is a Record type
type UserServiceRejectCredentialOffer struct {
	CredentialOfferCid types.CONTRACT_ID `json:"credentialOfferCid"`
	Reason             types.TEXT        `json:"reason"`
}

// ToMap converts UserServiceRejectCredentialOffer to a map for DAML arguments
func (t UserServiceRejectCredentialOffer) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialOfferCid"] = model.NestedToDAMLValue(t.CredentialOfferCid)

	m["reason"] = string(t.Reason)

	return m
}

func (t UserServiceRejectCredentialOffer) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRejectCredentialOffer) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRejectCredentialOffer to hex string (Canton MCMS format)
func (t UserServiceRejectCredentialOffer) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRejectCredentialOffer from hex string (Canton MCMS format)
func (t *UserServiceRejectCredentialOffer) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRequestToAdjustBillingParams is a Record type
type UserServiceRequestToAdjustBillingParams struct {
	CredentialBillingCid types.CONTRACT_ID `json:"credentialBillingCid"`
	BillingParams        BillingParams     `json:"billingParams"`
}

// ToMap converts UserServiceRequestToAdjustBillingParams to a map for DAML arguments
func (t UserServiceRequestToAdjustBillingParams) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["billingParams"] = model.NestedToDAMLValue(t.BillingParams)

	return m
}

func (t UserServiceRequestToAdjustBillingParams) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRequestToAdjustBillingParams) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRequestToAdjustBillingParams to hex string (Canton MCMS format)
func (t UserServiceRequestToAdjustBillingParams) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRequestToAdjustBillingParams from hex string (Canton MCMS format)
func (t *UserServiceRequestToAdjustBillingParams) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRevokeCredential is a Record type
type UserServiceRevokeCredential struct {
	CredentialCid types.CONTRACT_ID `json:"credentialCid"`
}

// ToMap converts UserServiceRevokeCredential to a map for DAML arguments
func (t UserServiceRevokeCredential) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialCid"] = model.NestedToDAMLValue(t.CredentialCid)

	return m
}

func (t UserServiceRevokeCredential) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRevokeCredential) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRevokeCredential to hex string (Canton MCMS format)
func (t UserServiceRevokeCredential) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRevokeCredential from hex string (Canton MCMS format)
func (t *UserServiceRevokeCredential) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceRevokeCredentialAndCancelBilling is a Record type
type UserServiceRevokeCredentialAndCancelBilling struct {
	CredentialCid        types.CONTRACT_ID  `json:"credentialCid"`
	CredentialBillingCid types.CONTRACT_ID  `json:"credentialBillingCid"`
	AppTransferContext   AppTransferContext `json:"appTransferContext"`
}

// ToMap converts UserServiceRevokeCredentialAndCancelBilling to a map for DAML arguments
func (t UserServiceRevokeCredentialAndCancelBilling) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialCid"] = model.NestedToDAMLValue(t.CredentialCid)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t UserServiceRevokeCredentialAndCancelBilling) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceRevokeCredentialAndCancelBilling) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceRevokeCredentialAndCancelBilling to hex string (Canton MCMS format)
func (t UserServiceRevokeCredentialAndCancelBilling) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceRevokeCredentialAndCancelBilling from hex string (Canton MCMS format)
func (t *UserServiceRevokeCredentialAndCancelBilling) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceTerminate is a Record type
type UserServiceTerminate struct {
	Actor types.PARTY `json:"actor"`
}

// ToMap converts UserServiceTerminate to a map for DAML arguments
func (t UserServiceTerminate) ToMap() map[string]any {
	m := make(map[string]any)

	m["actor"] = t.Actor.ToMap()

	return m
}

func (t UserServiceTerminate) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceTerminate) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceTerminate to hex string (Canton MCMS format)
func (t UserServiceTerminate) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceTerminate from hex string (Canton MCMS format)
func (t *UserServiceTerminate) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceTerminateResult is a Record type
type UserServiceTerminateResult struct {
}

// ToMap converts UserServiceTerminateResult to a map for DAML arguments
func (t UserServiceTerminateResult) ToMap() map[string]any {
	m := make(map[string]any)
	return m
}

func (t UserServiceTerminateResult) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceTerminateResult) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceTerminateResult to hex string (Canton MCMS format)
func (t UserServiceTerminateResult) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceTerminateResult from hex string (Canton MCMS format)
func (t *UserServiceTerminateResult) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// UserServiceTopUp is a Record type
type UserServiceTopUp struct {
	CredentialBillingCid types.CONTRACT_ID   `json:"credentialBillingCid"`
	AmountUsd            types.NUMERIC       `json:"amountUsd"`
	CoinCids             []types.CONTRACT_ID `json:"coinCids"`
	AppTransferContext   AppTransferContext  `json:"appTransferContext"`
}

// ToMap converts UserServiceTopUp to a map for DAML arguments
func (t UserServiceTopUp) ToMap() map[string]any {
	m := make(map[string]any)

	m["credentialBillingCid"] = model.NestedToDAMLValue(t.CredentialBillingCid)

	m["amountUsd"] = t.AmountUsd

	m["coinCids"] = func() []any {
		res := make([]any, 0, len(t.CoinCids))
		for _, e := range t.CoinCids {
			res = append(res, e)
		}
		return res
	}()

	m["appTransferContext"] = model.NestedToDAMLValue(t.AppTransferContext)

	return m
}

func (t UserServiceTopUp) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *UserServiceTopUp) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes UserServiceTopUp to hex string (Canton MCMS format)
func (t UserServiceTopUp) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes UserServiceTopUp from hex string (Canton MCMS format)
func (t *UserServiceTopUp) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// WithDsoOperator is a Record type
type WithDsoOperator struct {
	Dso      types.PARTY `json:"dso"`
	Operator types.PARTY `json:"operator"`
}

// ToMap converts WithDsoOperator to a map for DAML arguments
func (t WithDsoOperator) ToMap() map[string]any {
	m := make(map[string]any)

	m["dso"] = t.Dso.ToMap()

	m["operator"] = t.Operator.ToMap()

	return m
}

func (t WithDsoOperator) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *WithDsoOperator) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes WithDsoOperator to hex string (Canton MCMS format)
func (t WithDsoOperator) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes WithDsoOperator from hex string (Canton MCMS format)
func (t *WithDsoOperator) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// WithDsoOperatorHolder is a Record type
type WithDsoOperatorHolder struct {
	Dso      types.PARTY `json:"dso"`
	Operator types.PARTY `json:"operator"`
	Holder   types.PARTY `json:"holder"`
}

// ToMap converts WithDsoOperatorHolder to a map for DAML arguments
func (t WithDsoOperatorHolder) ToMap() map[string]any {
	m := make(map[string]any)

	m["dso"] = t.Dso.ToMap()

	m["operator"] = t.Operator.ToMap()

	m["holder"] = t.Holder.ToMap()

	return m
}

func (t WithDsoOperatorHolder) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *WithDsoOperatorHolder) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes WithDsoOperatorHolder to hex string (Canton MCMS format)
func (t WithDsoOperatorHolder) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes WithDsoOperatorHolder from hex string (Canton MCMS format)
func (t *WithDsoOperatorHolder) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// WithDsoOperatorIssuer is a Record type
type WithDsoOperatorIssuer struct {
	Dso      types.PARTY `json:"dso"`
	Operator types.PARTY `json:"operator"`
	Issuer   types.PARTY `json:"issuer"`
}

// ToMap converts WithDsoOperatorIssuer to a map for DAML arguments
func (t WithDsoOperatorIssuer) ToMap() map[string]any {
	m := make(map[string]any)

	m["dso"] = t.Dso.ToMap()

	m["operator"] = t.Operator.ToMap()

	m["issuer"] = t.Issuer.ToMap()

	return m
}

func (t WithDsoOperatorIssuer) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *WithDsoOperatorIssuer) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes WithDsoOperatorIssuer to hex string (Canton MCMS format)
func (t WithDsoOperatorIssuer) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes WithDsoOperatorIssuer from hex string (Canton MCMS format)
func (t *WithDsoOperatorIssuer) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// WithOperator is a Record type
type WithOperator struct {
	Operator types.PARTY `json:"operator"`
}

// ToMap converts WithOperator to a map for DAML arguments
func (t WithOperator) ToMap() map[string]any {
	m := make(map[string]any)

	m["operator"] = t.Operator.ToMap()

	return m
}

func (t WithOperator) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *WithOperator) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes WithOperator to hex string (Canton MCMS format)
func (t WithOperator) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes WithOperator from hex string (Canton MCMS format)
func (t *WithOperator) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// WithOperatorHolder is a Record type
type WithOperatorHolder struct {
	Operator types.PARTY `json:"operator"`
	Holder   types.PARTY `json:"holder"`
}

// ToMap converts WithOperatorHolder to a map for DAML arguments
func (t WithOperatorHolder) ToMap() map[string]any {
	m := make(map[string]any)

	m["operator"] = t.Operator.ToMap()

	m["holder"] = t.Holder.ToMap()

	return m
}

func (t WithOperatorHolder) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *WithOperatorHolder) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes WithOperatorHolder to hex string (Canton MCMS format)
func (t WithOperatorHolder) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes WithOperatorHolder from hex string (Canton MCMS format)
func (t *WithOperatorHolder) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// WithOperatorIssuerHolder is a Record type
type WithOperatorIssuerHolder struct {
	Operator types.PARTY `json:"operator"`
	Issuer   types.PARTY `json:"issuer"`
	Holder   types.PARTY `json:"holder"`
}

// ToMap converts WithOperatorIssuerHolder to a map for DAML arguments
func (t WithOperatorIssuerHolder) ToMap() map[string]any {
	m := make(map[string]any)

	m["operator"] = t.Operator.ToMap()

	m["issuer"] = t.Issuer.ToMap()

	m["holder"] = t.Holder.ToMap()

	return m
}

func (t WithOperatorIssuerHolder) MarshalJSON() ([]byte, error) {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Marshal(t)
}

func (t *WithOperatorIssuerHolder) UnmarshalJSON(data []byte) error {
	jsonCodec := codec.NewJsonCodec()
	return jsonCodec.Unmarshal(data, t)
}

// MarshalHex encodes WithOperatorIssuerHolder to hex string (Canton MCMS format)
func (t WithOperatorIssuerHolder) MarshalHex() (string, error) {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Marshal(t)
}

// UnmarshalHex decodes WithOperatorIssuerHolder from hex string (Canton MCMS format)
func (t *WithOperatorIssuerHolder) UnmarshalHex(data string) error {
	hexCodec := codec.NewHexCodec()
	return hexCodec.Unmarshal(data, t)
}

// MCMSEncoder interface for typed encoding methods.
// Implemented by Encoder for method-based encoding.
type MCMSEncoder interface {
	BillingParamsAdjustmentRequestAccept(args BillingParamsAdjustmentRequestAccept) (*bind.EncodedChoice, error)
	BillingParamsAdjustmentRequestCancel(args BillingParamsAdjustmentRequestCancel) (*bind.EncodedChoice, error)
	CredentialBillingAdjustBillingParams(args CredentialBillingAdjustBillingParams) (*bind.EncodedChoice, error)
	CredentialBillingBill(args CredentialBillingBill) (*bind.EncodedChoice, error)
	CredentialBillingCancel(args CredentialBillingCancel) (*bind.EncodedChoice, error)
	CredentialBillingCancelExpired(args CredentialBillingCancelExpired) (*bind.EncodedChoice, error)
	CredentialBillingDistribute(args CredentialBillingDistribute) (*bind.EncodedChoice, error)
	CredentialBillingDistributeAndAdjustDeposit(args CredentialBillingDistributeAndAdjustDeposit) (*bind.EncodedChoice, error)
	CredentialBillingFlushExpiredDeposit(args CredentialBillingFlushExpiredDeposit) (*bind.EncodedChoice, error)
	CredentialBillingRequestToAdjustBillingParams(args CredentialBillingRequestToAdjustBillingParams) (*bind.EncodedChoice, error)
	CredentialBillingTopUp(args CredentialBillingTopUp) (*bind.EncodedChoice, error)
	CredentialOfferAcceptFree(args CredentialOfferAcceptFree) (*bind.EncodedChoice, error)
	CredentialOfferAcceptPaid(args CredentialOfferAcceptPaid) (*bind.EncodedChoice, error)
	CredentialOfferCancel(args CredentialOfferCancel) (*bind.EncodedChoice, error)
	CredentialOfferReject(args CredentialOfferReject) (*bind.EncodedChoice, error)
	FeeRecordArchive(args FeeRecordArchive) (*bind.EncodedChoice, error)
	FeeRecordCalculateReward(args FeeRecordCalculateReward) (*bind.EncodedChoice, error)
	OperatorServiceAcceptUserServiceRequest(args OperatorServiceAcceptUserServiceRequest) (*bind.EncodedChoice, error)
	OperatorServiceRejectUserServiceRequest(args OperatorServiceRejectUserServiceRequest) (*bind.EncodedChoice, error)
	RewardRecordArchive(args RewardRecordArchive) (*bind.EncodedChoice, error)
	UserServiceRequestAccept(args UserServiceRequestAccept) (*bind.EncodedChoice, error)
	UserServiceRequestCancel(args UserServiceRequestCancel) (*bind.EncodedChoice, error)
	UserServiceRequestReject(args UserServiceRequestReject) (*bind.EncodedChoice, error)
	UserServiceAcceptFreeCredentialOffer(args UserServiceAcceptFreeCredentialOffer) (*bind.EncodedChoice, error)
	UserServiceAcceptPaidCredentialOffer(args UserServiceAcceptPaidCredentialOffer) (*bind.EncodedChoice, error)
	UserServiceAdjustBillingParams(args UserServiceAdjustBillingParams) (*bind.EncodedChoice, error)
	UserServiceBillingParamsAdjustmentRequestAccept(args UserServiceBillingParamsAdjustmentRequestAccept) (*bind.EncodedChoice, error)
	UserServiceBillingParamsAdjustmentRequestCancel(args UserServiceBillingParamsAdjustmentRequestCancel) (*bind.EncodedChoice, error)
	UserServiceCancelCredentialBilling(args UserServiceCancelCredentialBilling) (*bind.EncodedChoice, error)
	UserServiceCancelCredentialOffer(args UserServiceCancelCredentialOffer) (*bind.EncodedChoice, error)
	UserServiceDistribute(args UserServiceDistribute) (*bind.EncodedChoice, error)
	UserServiceDistributeAndAdjustDeposit(args UserServiceDistributeAndAdjustDeposit) (*bind.EncodedChoice, error)
	UserServiceDistributeAndAdjustDepositMulti(args UserServiceDistributeAndAdjustDepositMulti) (*bind.EncodedChoice, error)
	UserServiceDistributeMulti(args UserServiceDistributeMulti) (*bind.EncodedChoice, error)
	UserServiceOfferFreeCredential(args UserServiceOfferFreeCredential) (*bind.EncodedChoice, error)
	UserServiceOfferPaidCredential(args UserServiceOfferPaidCredential) (*bind.EncodedChoice, error)
	UserServiceRejectCredentialOffer(args UserServiceRejectCredentialOffer) (*bind.EncodedChoice, error)
	UserServiceRequestToAdjustBillingParams(args UserServiceRequestToAdjustBillingParams) (*bind.EncodedChoice, error)
	UserServiceRevokeCredential(args UserServiceRevokeCredential) (*bind.EncodedChoice, error)
	UserServiceRevokeCredentialAndCancelBilling(args UserServiceRevokeCredentialAndCancelBilling) (*bind.EncodedChoice, error)
	UserServiceTerminate(args UserServiceTerminate) (*bind.EncodedChoice, error)
	UserServiceTopUp(args UserServiceTopUp) (*bind.EncodedChoice, error)
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

// BillingParamsAdjustmentRequestAccept encodes parameters for the BillingParamsAdjustmentRequest_Accept choice.
func (e *encoder) BillingParamsAdjustmentRequestAccept(args BillingParamsAdjustmentRequestAccept) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("BillingParamsAdjustmentRequest_Accept", args)
}

// BillingParamsAdjustmentRequestCancel encodes parameters for the BillingParamsAdjustmentRequest_Cancel choice.
func (e *encoder) BillingParamsAdjustmentRequestCancel(args BillingParamsAdjustmentRequestCancel) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("BillingParamsAdjustmentRequest_Cancel", args)
}

// CredentialBillingAdjustBillingParams encodes parameters for the CredentialBilling_AdjustBillingParams choice.
func (e *encoder) CredentialBillingAdjustBillingParams(args CredentialBillingAdjustBillingParams) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_AdjustBillingParams", args)
}

// CredentialBillingBill encodes parameters for the CredentialBilling_Bill choice.
func (e *encoder) CredentialBillingBill(args CredentialBillingBill) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_Bill", args)
}

// CredentialBillingCancel encodes parameters for the CredentialBilling_Cancel choice.
func (e *encoder) CredentialBillingCancel(args CredentialBillingCancel) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_Cancel", args)
}

// CredentialBillingCancelExpired encodes parameters for the CredentialBilling_CancelExpired choice.
func (e *encoder) CredentialBillingCancelExpired(args CredentialBillingCancelExpired) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_CancelExpired", args)
}

// CredentialBillingDistribute encodes parameters for the CredentialBilling_Distribute choice.
func (e *encoder) CredentialBillingDistribute(args CredentialBillingDistribute) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_Distribute", args)
}

// CredentialBillingDistributeAndAdjustDeposit encodes parameters for the CredentialBilling_DistributeAndAdjustDeposit choice.
func (e *encoder) CredentialBillingDistributeAndAdjustDeposit(args CredentialBillingDistributeAndAdjustDeposit) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_DistributeAndAdjustDeposit", args)
}

// CredentialBillingFlushExpiredDeposit encodes parameters for the CredentialBilling_FlushExpiredDeposit choice.
func (e *encoder) CredentialBillingFlushExpiredDeposit(args CredentialBillingFlushExpiredDeposit) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_FlushExpiredDeposit", args)
}

// CredentialBillingRequestToAdjustBillingParams encodes parameters for the CredentialBilling_RequestToAdjustBillingParams choice.
func (e *encoder) CredentialBillingRequestToAdjustBillingParams(args CredentialBillingRequestToAdjustBillingParams) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_RequestToAdjustBillingParams", args)
}

// CredentialBillingTopUp encodes parameters for the CredentialBilling_TopUp choice.
func (e *encoder) CredentialBillingTopUp(args CredentialBillingTopUp) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialBilling_TopUp", args)
}

// CredentialOfferAcceptFree encodes parameters for the CredentialOffer_AcceptFree choice.
func (e *encoder) CredentialOfferAcceptFree(args CredentialOfferAcceptFree) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialOffer_AcceptFree", args)
}

// CredentialOfferAcceptPaid encodes parameters for the CredentialOffer_AcceptPaid choice.
func (e *encoder) CredentialOfferAcceptPaid(args CredentialOfferAcceptPaid) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialOffer_AcceptPaid", args)
}

// CredentialOfferCancel encodes parameters for the CredentialOffer_Cancel choice.
func (e *encoder) CredentialOfferCancel(args CredentialOfferCancel) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialOffer_Cancel", args)
}

// CredentialOfferReject encodes parameters for the CredentialOffer_Reject choice.
func (e *encoder) CredentialOfferReject(args CredentialOfferReject) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("CredentialOffer_Reject", args)
}

// FeeRecordArchive encodes parameters for the FeeRecord_Archive choice.
func (e *encoder) FeeRecordArchive(args FeeRecordArchive) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("FeeRecord_Archive", args)
}

// FeeRecordCalculateReward encodes parameters for the FeeRecord_CalculateReward choice.
func (e *encoder) FeeRecordCalculateReward(args FeeRecordCalculateReward) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("FeeRecord_CalculateReward", args)
}

// OperatorServiceAcceptUserServiceRequest encodes parameters for the OperatorService_AcceptUserServiceRequest choice.
func (e *encoder) OperatorServiceAcceptUserServiceRequest(args OperatorServiceAcceptUserServiceRequest) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("OperatorService_AcceptUserServiceRequest", args)
}

// OperatorServiceRejectUserServiceRequest encodes parameters for the OperatorService_RejectUserServiceRequest choice.
func (e *encoder) OperatorServiceRejectUserServiceRequest(args OperatorServiceRejectUserServiceRequest) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("OperatorService_RejectUserServiceRequest", args)
}

// RewardRecordArchive encodes parameters for the RewardRecord_Archive choice.
func (e *encoder) RewardRecordArchive(args RewardRecordArchive) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("RewardRecord_Archive", args)
}

// UserServiceRequestAccept encodes parameters for the UserServiceRequest_Accept choice.
func (e *encoder) UserServiceRequestAccept(args UserServiceRequestAccept) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserServiceRequest_Accept", args)
}

// UserServiceRequestCancel encodes parameters for the UserServiceRequest_Cancel choice.
func (e *encoder) UserServiceRequestCancel(args UserServiceRequestCancel) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserServiceRequest_Cancel", args)
}

// UserServiceRequestReject encodes parameters for the UserServiceRequest_Reject choice.
func (e *encoder) UserServiceRequestReject(args UserServiceRequestReject) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserServiceRequest_Reject", args)
}

// UserServiceAcceptFreeCredentialOffer encodes parameters for the UserService_AcceptFreeCredentialOffer choice.
func (e *encoder) UserServiceAcceptFreeCredentialOffer(args UserServiceAcceptFreeCredentialOffer) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_AcceptFreeCredentialOffer", args)
}

// UserServiceAcceptPaidCredentialOffer encodes parameters for the UserService_AcceptPaidCredentialOffer choice.
func (e *encoder) UserServiceAcceptPaidCredentialOffer(args UserServiceAcceptPaidCredentialOffer) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_AcceptPaidCredentialOffer", args)
}

// UserServiceAdjustBillingParams encodes parameters for the UserService_AdjustBillingParams choice.
func (e *encoder) UserServiceAdjustBillingParams(args UserServiceAdjustBillingParams) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_AdjustBillingParams", args)
}

// UserServiceBillingParamsAdjustmentRequestAccept encodes parameters for the UserService_BillingParamsAdjustmentRequest_Accept choice.
func (e *encoder) UserServiceBillingParamsAdjustmentRequestAccept(args UserServiceBillingParamsAdjustmentRequestAccept) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_BillingParamsAdjustmentRequest_Accept", args)
}

// UserServiceBillingParamsAdjustmentRequestCancel encodes parameters for the UserService_BillingParamsAdjustmentRequest_Cancel choice.
func (e *encoder) UserServiceBillingParamsAdjustmentRequestCancel(args UserServiceBillingParamsAdjustmentRequestCancel) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_BillingParamsAdjustmentRequest_Cancel", args)
}

// UserServiceCancelCredentialBilling encodes parameters for the UserService_CancelCredentialBilling choice.
func (e *encoder) UserServiceCancelCredentialBilling(args UserServiceCancelCredentialBilling) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_CancelCredentialBilling", args)
}

// UserServiceCancelCredentialOffer encodes parameters for the UserService_CancelCredentialOffer choice.
func (e *encoder) UserServiceCancelCredentialOffer(args UserServiceCancelCredentialOffer) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_CancelCredentialOffer", args)
}

// UserServiceDistribute encodes parameters for the UserService_Distribute choice.
func (e *encoder) UserServiceDistribute(args UserServiceDistribute) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_Distribute", args)
}

// UserServiceDistributeAndAdjustDeposit encodes parameters for the UserService_DistributeAndAdjustDeposit choice.
func (e *encoder) UserServiceDistributeAndAdjustDeposit(args UserServiceDistributeAndAdjustDeposit) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_DistributeAndAdjustDeposit", args)
}

// UserServiceDistributeAndAdjustDepositMulti encodes parameters for the UserService_DistributeAndAdjustDepositMulti choice.
func (e *encoder) UserServiceDistributeAndAdjustDepositMulti(args UserServiceDistributeAndAdjustDepositMulti) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_DistributeAndAdjustDepositMulti", args)
}

// UserServiceDistributeMulti encodes parameters for the UserService_DistributeMulti choice.
func (e *encoder) UserServiceDistributeMulti(args UserServiceDistributeMulti) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_DistributeMulti", args)
}

// UserServiceOfferFreeCredential encodes parameters for the UserService_OfferFreeCredential choice.
func (e *encoder) UserServiceOfferFreeCredential(args UserServiceOfferFreeCredential) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_OfferFreeCredential", args)
}

// UserServiceOfferPaidCredential encodes parameters for the UserService_OfferPaidCredential choice.
func (e *encoder) UserServiceOfferPaidCredential(args UserServiceOfferPaidCredential) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_OfferPaidCredential", args)
}

// UserServiceRejectCredentialOffer encodes parameters for the UserService_RejectCredentialOffer choice.
func (e *encoder) UserServiceRejectCredentialOffer(args UserServiceRejectCredentialOffer) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_RejectCredentialOffer", args)
}

// UserServiceRequestToAdjustBillingParams encodes parameters for the UserService_RequestToAdjustBillingParams choice.
func (e *encoder) UserServiceRequestToAdjustBillingParams(args UserServiceRequestToAdjustBillingParams) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_RequestToAdjustBillingParams", args)
}

// UserServiceRevokeCredential encodes parameters for the UserService_RevokeCredential choice.
func (e *encoder) UserServiceRevokeCredential(args UserServiceRevokeCredential) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_RevokeCredential", args)
}

// UserServiceRevokeCredentialAndCancelBilling encodes parameters for the UserService_RevokeCredentialAndCancelBilling choice.
func (e *encoder) UserServiceRevokeCredentialAndCancelBilling(args UserServiceRevokeCredentialAndCancelBilling) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_RevokeCredentialAndCancelBilling", args)
}

// UserServiceTerminate encodes parameters for the UserService_Terminate choice.
func (e *encoder) UserServiceTerminate(args UserServiceTerminate) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_Terminate", args)
}

// UserServiceTopUp encodes parameters for the UserService_TopUp choice.
func (e *encoder) UserServiceTopUp(args UserServiceTopUp) (*bind.EncodedChoice, error) {
	return e.EncodeChoiceArgs("UserService_TopUp", args)
}

// Verify MCMSEncoder interface implementation
var _ MCMSEncoder = (*encoder)(nil)
