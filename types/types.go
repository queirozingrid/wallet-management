package types

type Transaction struct {
	Id          int
	Amount      float64
	Description string
	Type        TransactionType
}

type TransactionType int

const (
	INFLOW  TransactionType = 1
	OUTFLOW TransactionType = 0
)
