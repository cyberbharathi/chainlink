package cmd

import "github.com/smartcontractkit/chainlink/core/services/keystore"

func (auth TerminalKeyStoreAuthenticator) ExportedValidatePasswordStrength(ethKeyStore *keystore.EthKeyStore, password string) error {
	return auth.validatePasswordStrength(ethKeyStore, password)
}
