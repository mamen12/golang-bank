package service

import "bank/entity"

type CustomerService interface {
	GetAll() []entity.Customer
}
