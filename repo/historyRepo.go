package repo

import "bank/entity"

type HistoryRepo interface {
	Create(history entity.History) error
}
