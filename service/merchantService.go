package service

import "bank/entity"

type MerchantService interface {
	GetAll() []entity.Merchant
	GetMerchantId(id string) entity.Merchant
}
