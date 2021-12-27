package entity

import "encoding/json"

type Account struct {
	id         string
	customerId string
	balance    int
}

func (acc *Account) GetId() string {
	return acc.id
}

func (acc *Account) GetName() string {
	return acc.customerId
}

func (acc *Account) GetBallance() int {
	return acc.balance
}
func (acc *Account) SetBalance(balance int) {
	acc.balance = balance
}

func (acc *Account) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id         string `json:"id"`
		CustomerId string `json:"customerId"`
		Balance    int    `json:"balance"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	acc.id = alias.Id
	acc.customerId = alias.CustomerId
	acc.balance = alias.Balance

	return nil
}

func (acc *Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         string `json:"id"`
		CustomerId string `json:"customerId"`
		Balance    int    `json:"balance"`
	}{
		Id:         acc.id,
		CustomerId: acc.customerId,
		Balance:    acc.balance,
	})
}
