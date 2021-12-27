package repo

import "bank/entity"

type TransactionRepo interface {
	Create(transaction entity.Transaction) error
}
