package service

import (
	"bank/entity"
	"bank/repo"
)

func NewAccountService(Account *repo.AccountRepo) AccountService {
	return &AccountServiceImpl{
		Account: *Account,
	}
}

type AccountServiceImpl struct {
	Account repo.AccountRepo
}

func (service *AccountServiceImpl) GetAll() []entity.Account {
	accounts := service.Account.GetAll()
	return accounts
}

func (service *AccountServiceImpl) GetAccountId(id string) entity.Account {
	account, _ := service.Account.GetAccountId(id)
	return account
}
