package repo

import "bank/entity"

type MerchantRepo interface {
	GetAll() []entity.Merchant
	GetMerchantId(id string) (entity.Merchant, error)
}
