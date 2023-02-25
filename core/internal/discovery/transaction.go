package discovery

type TransactionState uint16

const (
	TransactionStateRequested TransactionState = iota
	TransactionStateResponded
	TransactionStateFinished
)
