package fluxmonitorv2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/keystore/ethkey"
)

//go:generate mockery --name KeyStoreInterface --output ./mocks/ --case=underscore

// KeyStoreInterface defines an interface to interact with the keystore
type KeyStoreInterface interface {
	SendingKeys() ([]ethkey.Key, error)
	GetRoundRobinAddress(...common.Address) (common.Address, error)
}

// KeyStore implements KeyStoreInterface
type KeyStore struct {
	keystore.KeyStoreInterface
}

// NewKeyStore initializes a new keystore
func NewKeyStore(ks keystore.KeyStoreInterface) *KeyStore {
	return &KeyStore{ks}
}
