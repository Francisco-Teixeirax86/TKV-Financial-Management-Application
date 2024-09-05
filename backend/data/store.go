package data

import (
	"errors"
	"fmt"
	"sync"
)

type Store struct {
	mu       sync.RWMutex
	accounts map[string]float64
}

// NewStore Initializes a new in-mem store
func NewStore() *Store {
	return &Store{
		accounts: make(map[string]float64),
	}
}

// CreateAccount creates a new account on the store struct
func (s *Store) CreateAccount(accountID string, balance float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[accountID]; exists {
		return errors.New("account already exists")
	}
	s.accounts[accountID] = balance
	return nil
}

// GetBalance Get the balance of a specific account
func (s *Store) GetBalance(accountID string) (float64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	balance, exists := s.accounts[accountID]
	if !exists {
		return 0, fmt.Errorf("account %s does not exist", accountID)
	}

	return balance, nil
}

// UpdateBalance updates the balance of accountID with new balance value
func (s *Store) UpdateBalance(accountID string, balance float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[accountID]; !exists {
		return errors.New("account does not exist")
	}
	s.accounts[accountID] = balance
	return nil
}
