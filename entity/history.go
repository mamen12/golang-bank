package entity

import "encoding/json"

type History struct {
	id   string
	name string
	when string
}

func (h *History) GetId() string {
	return h.id
}

func (h *History) GetName() string {
	return h.name
}

func (h *History) GetWhen() string {
	return h.when
}
func (h *History) SetWhen(when string) {
	h.when = when
}
func (h *History) SetId(id string) {
	h.id = id
}
func (h *History) SetName(name string) {
	h.name = name
}

func (h *History) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		When string `json:"when"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	h.id = alias.Id
	h.name = alias.Name
	h.when = alias.When

	return nil
}

func (h *History) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		When string `json:"when"`
	}{
		Id:   h.id,
		Name: h.name,
		When: h.when,
	})
}
