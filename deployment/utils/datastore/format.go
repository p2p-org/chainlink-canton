package datastore

import (
	"fmt"

	datastore2 "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-canton/contracts"
)

// ToInstanceAddress derives the hashed instance address from the raw label stored on the ref.
func ToInstanceAddress(ref datastore.AddressRef) (contracts.InstanceAddress, error) {
	raw, err := GetRawInstanceAddressFromAddressRef(ref)
	if err != nil {
		return contracts.InstanceAddress{}, fmt.Errorf("resolve raw instance address from ref %s: %w", datastore2.SprintRef(ref), err)
	}

	return raw.InstanceAddress(), nil
}

func ToInstanceAddressBytes(ref datastore.AddressRef) ([]byte, error) {
	addr, err := ToInstanceAddress(ref)
	if err != nil {
		return nil, err
	}

	return addr.Bytes(), nil
}
