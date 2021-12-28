package service

import "bank/entity"

type HistoryService interface {
	Create(history entity.History) error
}
