package services

import "PenzaTestTask/internal/models"

type AccountService interface {
	CreateAccount() models.BankAccount
	Deposit(id int, amount float64) error
	Withdraw(id int, amount float64) error
	GetBalance(id int) (float64, error)
}
