package repo

import (
	"bank/entity"
	"encoding/json"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

func NewAccountRepo() AccountRepo {
	return &AccountRepoImpl{}
}

type AccountRepoImpl struct {
}

func (repo *AccountRepoImpl) GetAll() []entity.Account {
	var accounts []entity.Account
	byteValue, _ := ioutil.ReadFile("./jsonDb/account.json")
	json.Unmarshal(byteValue, &accounts)
	return accounts
}

func (repo *AccountRepoImpl) GetAccountId(id string) (entity.Account, error) {
	byteValue, err := ioutil.ReadFile("./jsonDb/account.json")
	if err != nil {
		return entity.Account{}, err
	}
	data := gjson.Get(string(byteValue), `#(id=="`+id+`")`).String()

	var account entity.Account

	err = json.Unmarshal([]byte(data), &account)
	return account, err
}
