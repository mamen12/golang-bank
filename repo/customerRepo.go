package repo

import "bank/entity"

type CustomerRepo interface {
	GetAll() []entity.Customer
}
