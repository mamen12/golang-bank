package entity

import "encoding/json"

type Merchant struct {
	id      string
	name    string
	balance int
}

func (m *Merchant) GetId() string {
	return m.id
}

func (m *Merchant) GetName() string {
	return m.name
}

func (m *Merchant) GetBallance() int {
	return m.balance
}
func (m *Merchant) SetBalance(balance int) {
	m.balance = balance
}

func (m *Merchant) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Ballance int    `json:"balance"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	m.id = alias.Id
	m.name = alias.Name
	m.balance = alias.Ballance

	return nil
}

func (m *Merchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Ballance int    `json:"balance"`
	}{
		Id:       m.id,
		Name:     m.name,
		Ballance: m.balance,
	})
}
