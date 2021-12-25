package service

import (
	"bank/entity"
	"bank/repo"
)

func NewCustomerService(customerRepo *repo.CustomerRepo) CustomerService {
	return &CustomerServiceImpl{
		customerRepo: *customerRepo,
	}
}

type CustomerServiceImpl struct {
	customerRepo repo.CustomerRepo
}

func (service *CustomerServiceImpl) GetAll() []entity.Customer {
	customers := service.customerRepo.GetAll()
	return customers
}
