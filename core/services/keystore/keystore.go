package keystore

import (
	"github.com/smartcontractkit/chainlink/core/utils"
	"gorm.io/gorm"
)

func NewKeyStore(db *gorm.DB, scryptParams utils.ScryptParams) *KeyStore {
	return &KeyStore{
		Eth: NewEthKeyStore(db, scryptParams),
		OCR: NewOCRKeyStore(db, scryptParams),
	}
}

type KeyStore struct {
	Eth *EthKeyStore
	OCR *OCRKeyStore
}
