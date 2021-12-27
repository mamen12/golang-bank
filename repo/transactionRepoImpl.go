package repo

import (
	"bank/entity"
	"encoding/json"
	"io/ioutil"
)

func NewTransactionRepo() TransactionRepo {
	return &TransactionRepoImpl{}
}

type TransactionRepoImpl struct {
}

func (repo *TransactionRepoImpl) Create(transaction entity.Transaction) error {
	var transactions []entity.Transaction
	byteValue, _ := ioutil.ReadFile("./jsonDb/transaction.json")
	json.Unmarshal(byteValue, &transactions)

	transactions = append(transactions, transaction)
	tokenByte, err := json.Marshal(transactions)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./jsonDb/transaction.json", tokenByte, 0644)
	return err
}
