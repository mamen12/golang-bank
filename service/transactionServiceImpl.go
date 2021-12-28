package service

import (
	"bank/entity"
	"bank/repo"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

func NewTransactionService(transactionRepo *repo.TransactionRepo) TransactionService {
	accountRepo := repo.NewAccountRepo()
	accountService := NewAccountService(&accountRepo)

	merchantRepo := repo.NewMerchantRepo()
	merchantService := NewMerchantService(&merchantRepo)

	historyRepo := repo.NewHistoryRepo()
	historyService := NewHistoryService(&historyRepo)

	return &TransactionServiceImpl{
		AccountService:  accountService,
		MerchantService: merchantService,
		HistoryService:  historyService,
		TransactionRepo: *transactionRepo,
	}
}

type TransactionServiceImpl struct {
	AccountService  AccountService
	HistoryService  HistoryService
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

	if account.GetId() == "" {
		err := errors.New("access ilegal, data not found")
		return err

	} else if merchant.GetId() == "" {
		err := errors.New("access ilegal, data not found")
		return err

	} else {
		transaction.SetId("transA")
		account.SetBill(bill)

		//save To history
		var history entity.History
		currentTime := time.Now()
		history.SetId("payment merchant - " + id_account)
		history.SetWhen(currentTime.Format("2006-01-02 15:04:05"))
		history.SetName("Payment merchant by id account")
		service.HistoryService.Create(history)

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
				accounts = append(accounts, account)
				tokensByte, _ := json.Marshal(accounts)
				err = ioutil.WriteFile("./jsonDb/account.json", tokensByte, 0644)
				return err
			}
		}
	}

	return err

}
