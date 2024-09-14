package data

type LogEntry struct {
	Term    int
	Command string
}

type LogStore struct {
	Logs []LogEntry
}

func NewLogStore() *LogStore {
	return &LogStore{
		Logs: make([]LogEntry, 0),
	}
}

// AddLogEntry adds a new transaction log entry to the log store
func (ls *LogStore) AddLogEntry(entry LogEntry) {
	ls.Logs = append(ls.Logs, entry)
}

// GetLogs returns the current list of logs
func (ls *LogStore) GetLogs() []LogEntry {
	return ls.Logs
}
