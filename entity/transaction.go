package entity

import "encoding/json"

type Transaction struct {
	id         string
	accountId  string
	merchantId string
	amount     int
}

func (tx *Transaction) GetId() string {
	return tx.id
}

func (tx *Transaction) GetAccountId() string {
	return tx.accountId
}

func (tx *Transaction) GetMerchantId() string {
	return tx.merchantId
}

func (tx *Transaction) GetAmount() int {
	return tx.amount
}

func (tx *Transaction) SetAccountId(accountId string) {
	tx.accountId = accountId
}

func (tx *Transaction) SetId(id string) {
	tx.id = id
}

func (tx *Transaction) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id         string `json:"id"`
		AccountId  string `json:"account_id"`
		MerchantId string `json:"merchant_id"`
		Amount     int    `json:"amount"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	tx.id = alias.Id
	tx.accountId = alias.AccountId
	tx.merchantId = alias.MerchantId
	tx.amount = alias.Amount

	return nil
}

func (tx *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         string `json:"id"`
		AccountId  string `json:"account_id"`
		MerchantId string `json:"merchant_id"`
		Amount     int    `json:"amount"`
	}{
		Id:         tx.id,
		AccountId:  tx.accountId,
		MerchantId: tx.merchantId,
		Amount:     tx.amount,
	})
}
