package entity

import "encoding/json"

type Account struct {
	id        string
	accountId string
	bill      int
}

func (acc *Account) GetId() string {
	return acc.id
}

func (acc *Account) GetAccountId() string {
	return acc.accountId
}

func (acc *Account) GetBill() int {
	return acc.bill
}
func (acc *Account) SetBill(bill int) {
	acc.bill = bill
}

func (acc *Account) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id        string `json:"id"`
		AccountId string `json:"accountId"`
		Bill      int    `json:"bill"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	acc.id = alias.Id
	acc.accountId = alias.AccountId
	acc.bill = alias.Bill

	return nil
}

func (acc *Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id        string `json:"id"`
		AccountId string `json:"accountId"`
		Bill      int    `json:"bill"`
	}{
		Id:        acc.id,
		AccountId: acc.accountId,
		Bill:      acc.bill,
	})
}
