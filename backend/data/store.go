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

// Deposit updated the balance of accountID by depositing a new amount
func (s *Store) Deposit(accountID string, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.accounts[accountID]; !exists {
		return errors.New("account does not exist")
	}

	s.accounts[accountID] += amount
	return nil
}

// Withdraw update the balance of accountID by withdrawing funds
func (s *Store) Withdraw(accountID string, amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	balance, exists := s.accounts[accountID]
	if !exists {
		return errors.New("account does not exist")
	}

	if balance < amount {
		return errors.New("insufficient funds")
	}

	s.accounts[accountID] -= amount
	return nil
}

func (s *Store) Transfer(fromAccountID, toAccountID string, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	fromBalance, fromExists := s.accounts[fromAccountID]
	toBalance, toExists := s.accounts[toAccountID]

	if !fromExists {
		return errors.New("source account does not exist")
	}

	if !toExists {
		return errors.New("destination account does not exist")
	}

	if fromBalance < amount {
		return errors.New("insufficient funds in the source account")
	}

	s.accounts[toAccountID] = toBalance + amount
	s.accounts[fromAccountID] = fromBalance - amount

	return nil
}
