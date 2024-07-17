package repository

import "PenzaTestTask/internal/models"

type AccountRepository interface {
	CreateAccount() models.BankAccount
	GetAccount(id int) (models.BankAccount, error)
}
