package service

import "bank/entity"

type TransactionService interface {
	Create(transaction entity.Transaction) error
}
