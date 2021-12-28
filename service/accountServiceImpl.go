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
	merchants := service.Account.GetAll()
	return merchants
}

func (service *AccountServiceImpl) GetAccountId(id string) entity.Account {
	merchant, _ := service.Account.GetAccountId(id)
	return merchant
}
