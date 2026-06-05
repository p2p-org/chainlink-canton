package ccip

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-canton/bindings"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/latest/ccip/tokenadminregistry"
	"github.com/smartcontractkit/chainlink-canton/contracts"
	"github.com/smartcontractkit/chainlink-canton/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-canton/registry-kit/ledger"
	"github.com/smartcontractkit/chainlink-canton/testhelpers"
	"github.com/smartcontractkit/chainlink-canton/bindings/generated/v1_0_0/splice/splice_api_token_holding_v1"
	"github.com/smartcontractkit/go-daml/pkg/types"
)

// Validate performs read-only checks that TAR maps the instrument to the expected pool.
func Validate(
	ctx context.Context,
	client ledger.Client,
	instrumentID splice_api_token_holding_v1.InstrumentId,
	tarInstanceAddr contracts.InstanceAddress,
	ccipParty, poolInstanceID string,
) error {
	tokenConfigAddr := contracts.InstanceID(hex.EncodeToString(contracts.EncodeInstrumentID(instrumentID).Bytes())).
		RawInstanceAddress(types.PARTY(ccipParty)).
		InstanceAddress()

	participant := client.Participant()
	active, err := contract.FindActiveContractByInstanceAddress(
		ctx,
		participant.LedgerServices.State,
		contract.LedgerQueryParties(participant),
		tokenadminregistry.TokenConfig{}.GetTemplateID(),
		tokenConfigAddr,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no active contract found") {
			return fmt.Errorf("token config not found for instrument %s/%s", instrumentID.Admin, instrumentID.Id)
		}
		return fmt.Errorf("fetch token config: %w", err)
	}
	cfg, err := bindings.UnmarshalCreatedEvent[tokenadminregistry.TokenConfig](active.GetCreatedEvent())
	if err != nil {
		return fmt.Errorf("unmarshal token config: %w", err)
	}
	if cfg.Admin == nil || string(*cfg.Admin) == "" {
		return fmt.Errorf("token admin not set for instrument %s/%s", instrumentID.Admin, instrumentID.Id)
	}
	if cfg.TokenPool == nil {
		return fmt.Errorf("token pool not registered for instrument %s/%s", instrumentID.Admin, instrumentID.Id)
	}
	if string(cfg.TokenPool.PoolInstanceId) != poolInstanceID {
		return fmt.Errorf("pool instance: expected %s got %s", poolInstanceID, cfg.TokenPool.PoolInstanceId)
	}

	// Holdings visibility: query succeeds (may be zero balance).
	_, err = testhelpers.ListHoldingsForInstrument(ctx, client.ForParty(string(instrumentID.Admin)), &instrumentID)
	if err != nil {
		return fmt.Errorf("holdings not visible for instrument: %w", err)
	}

	_ = tarInstanceAddr

	return nil
}
