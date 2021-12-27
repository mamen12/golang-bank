package service

import (
	"bank/entity"
	"bank/repo"
)

func NewTransactionService(transactionRepo *repo.TransactionRepo) TransactionService {
	return &repo.TransactionRepoImpl{}
}

type TransactionServiceImpl struct {
	transactionRepo repo.TransactionRepo
}

func (service *TransactionServiceImpl) Create(transaction entity.Transaction) error {
	transaction.SetId("transA")
	err := service.transactionRepo.Create(transaction)
	return err
}
