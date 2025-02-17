package wallet

import (
	"github.com/iotaledger/hive.go/bitmask"

	"github.com/iotaledger/goshimmer/client"
	"github.com/iotaledger/goshimmer/client/wallet/packages/seed"
)

// Option represents an optional parameter .
type Option func(*Wallet)

// WebAPI connects the wallet with the remote API of a node.
func WebAPI(baseURL string, setters ...client.Option) Option {
	return func(wallet *Wallet) {
		wallet.connector = NewWebConnector(baseURL, setters...)
	}
}

// Import restores a wallet that has previously been created.
func Import(seed *seed.Seed, lastAddressIndex uint64, spentAddresses []bitmask.BitMask, assetRegistry *AssetRegistry) Option {
	return func(wallet *Wallet) {
		wallet.addressManager = NewAddressManager(seed, lastAddressIndex, spentAddresses)
		wallet.assetRegistry = assetRegistry
	}
}

// ReusableAddress configures the wallet to run in "single address" mode where all the funds are always managed on a
// single reusable address.
func ReusableAddress(enabled bool) Option {
	return func(wallet *Wallet) {
		wallet.reusableAddress = enabled
	}
}

// ConfirmationPollingInterval defines how often the wallet polls the node for confirmation info.
func ConfirmationPollingInterval(interval int) Option {
	return func(wallet *Wallet) {
		if interval < 0 {
			wallet.ConfirmationPollInterval = DefaultPollingInterval
		} else {
			wallet.ConfirmationPollInterval = interval
		}
	}
}

// ConfirmationTimeout defines the timeout for waiting for tx confirmation.
func ConfirmationTimeout(timeout int) Option {
	return func(wallet *Wallet) {
		if timeout < 0 {
			wallet.ConfirmationPollInterval = DefaultConfirmationTimeout
		} else {
			wallet.ConfirmationPollInterval = timeout
		}
	}
}

// GenericConnector allows us to provide a generic connector to the wallet. It can be used to mock the behavior of a
// real connector in tests or to provide new connection methods for nodes.
func GenericConnector(connector Connector) Option {
	return func(wallet *Wallet) {
		wallet.connector = connector
	}
}
