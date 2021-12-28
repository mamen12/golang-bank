package service

import (
	"bank/entity"
	"bank/repo"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func NewTransactionService(transactionRepo *repo.TransactionRepo) TransactionService {
	accountRepo := repo.NewAccountRepo()
	accountService := NewAccountService(&accountRepo)

	merchantRepo := repo.NewMerchantRepo()
	merchantService := NewMerchantService(&merchantRepo)

	return &TransactionServiceImpl{
		AccountService:  accountService,
		MerchantService: merchantService,
		TransactionRepo: *transactionRepo,
	}
}

type TransactionServiceImpl struct {
	AccountService  AccountService
	MerchantService MerchantService
	TransactionRepo repo.TransactionRepo
}

func (service *TransactionServiceImpl) Create(transaction entity.Transaction) error {

	var account entity.Account
	var id_account string

	var merchant entity.Merchant
	var id_merchant string

	// getAccountById
	id_account = transaction.GetAccountId()
	account = service.AccountService.GetAccountId(id_account)

	// getMerchantById
	id_merchant = transaction.GetMerchantId()
	merchant = service.MerchantService.GetMerchantId(id_merchant)

	// payment merchant
	bill := account.GetBill() - transaction.GetAmount()
	err := service.TransactionRepo.Create(transaction)
	fmt.Println(err)
	fmt.Println(bill)
	if account.GetId() == "" {
		err := errors.New("access ilegal, data tidak ditemukan")
		return err

	} else if merchant.GetId() == "" {
		err := errors.New("access ilegal, data tidak ditemukan")
		return err

	} else {
		transaction.SetId("transA")
		account.SetBill(bill)

		//get list account
		var accounts []entity.Account
		byteValue, _ := ioutil.ReadFile("./jsonDb/account.json")
		err := json.Unmarshal(byteValue, &accounts)
		if err != nil {
			return err
		}
		//update account
		for idx, val := range accounts {
			if val.GetId() == id_account {
				accounts = append(accounts[:idx], accounts[idx+1:]...)
				fmt.Println(accounts)
				accounts = append(accounts, account)
				fmt.Println(accounts)
				tokensByte, _ := json.Marshal(accounts)
				s := string(tokensByte)
				fmt.Println(s)
				err = ioutil.WriteFile("./jsonDb/account.json", tokensByte, 0644)
				return err
			}
		}
	}

	return err

}
