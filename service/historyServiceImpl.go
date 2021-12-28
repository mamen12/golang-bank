package service

import (
	"bank/entity"
	"bank/repo"
)

func NewHistoryService(historyRepo *repo.HistoryRepo) HistoryService {
	return &HistoryServiceImpl{
		HistoryRepo: *historyRepo,
	}
}

type HistoryServiceImpl struct {
	HistoryRepo repo.HistoryRepo
}

func (service *HistoryServiceImpl) Create(history entity.History) error {
	err := service.HistoryRepo.Create(history)
	return err
}
