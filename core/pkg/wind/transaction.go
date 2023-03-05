package wind

type TransactionState uint16

const (
	TransactionStateRequested TransactionState = iota
	TransactionStateResponded
	TransactionStateFinished
)
