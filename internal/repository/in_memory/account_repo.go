package in_memory

import (
	"PenzaTestTask/internal/models"
	"PenzaTestTask/internal/repository"
	"errors"
	"sync"
)

type inMemoryAccountRepository struct {
	accounts map[int]*models.BankAccount
	mu       sync.Mutex
	nextID   int
}

func NewInMemoryAccountRepository() repository.AccountRepository {
	return &inMemoryAccountRepository{
		accounts: make(map[int]*models.BankAccount),
		nextID:   1,
	}
}

func (r *inMemoryAccountRepository) CreateAccount() models.BankAccount {
	r.mu.Lock()
	defer r.mu.Unlock()

	account := models.NewAccount(r.nextID)
	r.accounts[r.nextID] = &account
	r.nextID++
	return account
}

func (r *inMemoryAccountRepository) GetAccount(id int) (models.BankAccount, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	account, exists := r.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}
	return *account, nil
}
