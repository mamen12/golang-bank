package repo

import (
	"bank/entity"
	"encoding/json"
	"io/ioutil"
)

func NewCustomerRepo() CustomerRepo {
	return &CustomerRepoImpl{}
}

type CustomerRepoImpl struct {
}

func (repo *CustomerRepoImpl) GetAll() []entity.Customer {
	var customers []entity.Customer
	byteValue, _ := ioutil.ReadFile("./jsonDb/customer.json")
	json.Unmarshal(byteValue, &customers)
	return customers
}
