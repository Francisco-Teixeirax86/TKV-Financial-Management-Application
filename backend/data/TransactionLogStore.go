package data

import "time"

// TransactionLogEntry for storing transaction-specific information
type TransactionLogEntry struct {
	Timestamp     time.Time
	TransactionID string
	Type          string
	AccountID     string
	Amount        float64
	ResultBalance float64
}

// TransactionLogStore for storing all transaction logs
type TransactionLogStore struct {
	Transactions []TransactionLogEntry // Transaction-specific logs
}

// NewTransactionLogStore Initializes a new TransactionLogStore
func NewTransactionLogStore() *TransactionLogStore {
	return &TransactionLogStore{
		Transactions: make([]TransactionLogEntry, 0),
	}
}

// AddTransactionLog adds a new transaction log entry
func (tls *TransactionLogStore) AddTransactionLog(entry TransactionLogEntry) {
	tls.Transactions = append(tls.Transactions, entry)
}

// GetTransactionLogs returns the transaction logs
func (tls *TransactionLogStore) GetTransactionLogs() []TransactionLogEntry {
	return tls.Transactions
}
