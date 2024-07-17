package app

import (
	"PenzaTestTask/internal/handlers"
	"PenzaTestTask/internal/repository"
	"PenzaTestTask/internal/repository/in_memory"
	"PenzaTestTask/internal/services"
)

type serviceProvider struct {
	accountRepo    repository.AccountRepository
	accountService services.AccountService
	accountHandler handlers.AccountHandler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) initServices() error {
	sp.accountRepo = in_memory.NewInMemoryAccountRepository()
	sp.accountService = services.NewAccountService(sp.accountRepo)
	sp.accountHandler = handlers.NewAccountHandler(sp.accountService)
	return nil
}
