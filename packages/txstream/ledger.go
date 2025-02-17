package txstream

import (
	"github.com/iotaledger/hive.go/events"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
)

// Ledger is the interface between txstream and the underlying value tangle
type Ledger interface {
	GetUnspentOutputs(addr ledgerstate.Address, f func(ledgerstate.Output))
	GetOutput(outID ledgerstate.OutputID, f func(ledgerstate.Output)) bool
	GetOutputMetadata(outID ledgerstate.OutputID, f func(*ledgerstate.OutputMetadata)) bool
	GetConfirmedTransaction(txid ledgerstate.TransactionID, f func(*ledgerstate.Transaction)) bool
	GetTxInclusionState(txid ledgerstate.TransactionID) (ledgerstate.InclusionState, error)
	EventTransactionConfirmed() *events.Event
	EventTransactionBooked() *events.Event
	PostTransaction(tx *ledgerstate.Transaction) error
	RequestFunds(target ledgerstate.Address) error
	Detach()
}
