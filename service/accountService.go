package service

import "bank/entity"

type AccountService interface {
	GetAll() []entity.Account
	GetAccountId(id string) entity.Account
}
