package service

import (
	"bank/entity"
	"bank/repo"
)

func NewAccountService(merchantRepo *repo.AccountRepo) AccountService {
	return &AccountServiceImpl{
		merchantRepo: *merchantRepo,
	}
}

type AccountServiceImpl struct {
	merchantRepo repo.AccountRepo
}

func (service *AccountServiceImpl) GetAll() []entity.Account {
	merchants := service.merchantRepo.GetAll()
	return merchants
}

func (service *AccountServiceImpl) GetAccountId(id string) entity.Account {
	merchant, _ := service.merchantRepo.GetAccountId(id)
	return merchant
}
