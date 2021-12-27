package service

import "bank/entity"

type AuthService interface {
	Login(username, password string) (string, error)
	GenerateToken(customer entity.Customer) string
	Logout(token string) error
}
