package domain

type TransactionRepository interface {
	Store(transaction Transaction) (Transaction, error)
}

type Transaction struct {
	Id       string
	Type     string
	Customer Customer
	Amount   int
	Currency string
}

func (transaction *Transaction) IsValidCharge() bool {
	return true
}
