package repo

import (
	"bank/entity"
	"encoding/json"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

func NewMerchantRepo() MerchantRepo {
	return &MerchantRepoImpl{}
}

type MerchantRepoImpl struct {
}

func (repo *MerchantRepoImpl) GetAll() []entity.Merchant {
	var merchant []entity.Merchant
	byteValue, _ := ioutil.ReadFile("./jsonDb/merchant.json")
	json.Unmarshal(byteValue, &merchant)
	return merchant
}

func (repo *MerchantRepoImpl) GetMerchantId(id string) (entity.Merchant, error) {
	byteValue, err := ioutil.ReadFile("./jsonDb/merchant.json")
	if err != nil {
		return entity.Merchant{}, err
	}
	data := gjson.Get(string(byteValue), `#(id=="`+id+`")`).String()

	var merchant entity.Merchant

	err = json.Unmarshal([]byte(data), &merchant)
	return merchant, err
}
