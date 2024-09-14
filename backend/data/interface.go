package data

type StoreInterface interface {
	CreateAccount(accountID string, balance float64) error
	GetBalance(accountID string) (float64, error)
	UpdateBalance(accountID string, balance float64) error
	Deposit(transactionLogStore *TransactionLogStore, accountID string, amount float64) error
	Withdraw(transactionLogStore *TransactionLogStore, accountID string, amount float64) error
	Transfer(transactionLogStore *TransactionLogStore, from string, to string, amount float64) error
}
