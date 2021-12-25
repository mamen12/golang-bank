package repo

import "bank/entity"

type AuthRepo interface {
	Login(username, password string) (entity.Customer, error)
	SaveToken(token string) error
}
