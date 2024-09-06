package data

import "time"

type TransactionLogEntry struct {
	Timestamp     time.Time
	TransactionID string
	Type          string
	AccountID     string
	Amount        float64
	ResultBalance float64
}

type LogStore struct {
	Logs []TransactionLogEntry
}

func NewLogStore() *LogStore {
	return &LogStore{
		Logs: make([]TransactionLogEntry, 0),
	}
}

// AddLogEntry adds a new transaction log entry to the log store
func (ls *LogStore) AddLogEntry(entry TransactionLogEntry) {
	ls.Logs = append(ls.Logs, entry)
}
