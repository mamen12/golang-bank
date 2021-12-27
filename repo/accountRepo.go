package repo

import "bank/entity"

type AccountRepo interface {
	GetAll() []entity.Account
	GetAccountId(id string) (entity.Account, error)
}
