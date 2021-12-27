package entity

import "encoding/json"

type Transaction struct {
	id         string
	customerId string
	merchantId string
	amount     int
}

func (t *Transaction) SetCustomerId(customerId string) {
	t.customerId = customerId
}

func (t *Transaction) SetId(id string) {
	t.id = id
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id         string `json:"id"`
		CustomerId string `json:"customer_id"`
		MerchantId string `json:"merchant_id"`
		Amount     int    `json:"amount"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	t.id = alias.Id
	t.customerId = alias.CustomerId
	t.merchantId = alias.MerchantId
	t.amount = alias.Amount

	return nil
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         string `json:"id"`
		CustomerId string `json:"customer_id"`
		MerchantId string `json:"merchant_id"`
		Amount     int    `json:"amount"`
	}{
		Id:         t.id,
		CustomerId: t.customerId,
		MerchantId: t.merchantId,
		Amount:     t.amount,
	})
}
