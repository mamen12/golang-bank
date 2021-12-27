package service

import (
	"bank/entity"
	"bank/repo"
)

func NewMerchantService(merchantRepo *repo.MerchantRepo) MerchantService {
	return &MerchantServiceImpl{
		merchantRepo: *merchantRepo,
	}
}

type MerchantServiceImpl struct {
	merchantRepo repo.MerchantRepo
}

func (service *MerchantServiceImpl) GetAll() []entity.Merchant {
	merchants := service.merchantRepo.GetAll()
	return merchants
}

func (service *MerchantServiceImpl) GetMerchantId(id string) entity.Merchant {
	merchant, _ := service.merchantRepo.GetMerchantId(id)
	return merchant
}
