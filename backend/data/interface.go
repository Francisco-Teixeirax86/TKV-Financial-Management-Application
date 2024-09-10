package data

type StoreInterface interface {
	CreateAccount(accountID string, balance float64) error
	GetBalance(accountID string) (float64, error)
	UpdateBalance(accountID string, balance float64) error
	Deposit(logStore *LogStore, accountID string, amount float64) error
	Withdraw(logStore *LogStore, accountID string, amount float64) error
	Transfer(logStore *LogStore, from string, to string, amount float64) error
}
