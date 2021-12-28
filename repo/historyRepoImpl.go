package repo

import (
	"bank/entity"
	"encoding/json"
	"io/ioutil"
)

func NewHistoryRepo() HistoryRepo {
	return &HistoryRepoImpl{}
}

type HistoryRepoImpl struct {
}

func (repo *HistoryRepoImpl) Create(history entity.History) error {
	var historys []entity.History
	byteValue, _ := ioutil.ReadFile("./jsonDb/history.json")
	json.Unmarshal(byteValue, &historys)

	historys = append(historys, history)
	tokenByte, err := json.Marshal(historys)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./jsonDb/history.json", tokenByte, 0644)
	return err
}
