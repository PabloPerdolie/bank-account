package services

import (
	"PenzaTestTask/internal/models"
	"PenzaTestTask/internal/repository"
)

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) CreateAccount() models.BankAccount {
	return s.repo.CreateAccount()
}

func (s *accountService) Deposit(id int, amount float64) error {
	account, err := s.repo.GetAccount(id)
	if err != nil {
		return err
	}
	return account.Deposit(amount)
}

func (s *accountService) Withdraw(id int, amount float64) error {
	account, err := s.repo.GetAccount(id)
	if err != nil {
		return err
	}
	return account.Withdraw(amount)
}

func (s *accountService) GetBalance(id int) (float64, error) {
	account, err := s.repo.GetAccount(id)
	if err != nil {
		return 0, err
	}
	return account.GetBalance(), nil
}
