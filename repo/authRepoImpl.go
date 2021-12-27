package repo

import (
	"bank/dto"
	"bank/entity"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

func NewAuthRepo() AuthRepo {
	return &AuthRepoImpl{}
}

type AuthRepoImpl struct {
}

func (repo *AuthRepoImpl) Login(username, password string) (entity.Customer, error) {
	byteValue, err := ioutil.ReadFile("./jsonDb/customer.json")
	if err != nil {
		return entity.Customer{}, err
	}
	data := gjson.Get(string(byteValue), `#(username=="`+username+`")`).String()
	passMatch := gjson.Get(data, `password`).String()

	if password != passMatch {
		return entity.Customer{}, errors.New("username or password incorrect")
	}

	var customer entity.Customer

	err = json.Unmarshal([]byte(data), &customer)
	if err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}
func (repo *AuthRepoImpl) SaveToken(token string) error {
	// read token list
	var tokens []dto.Authorize
	byteValue, _ := ioutil.ReadFile("./jsonDb/authorize.json")
	err := json.Unmarshal(byteValue, &tokens)
	if err != nil {
		return err
	}

	tokens = append(tokens, dto.Authorize{Token: token})
	tokenByte, err := json.Marshal(tokens)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./jsonDb/authorize.json", tokenByte, 0644)
	return err
}

func (repo *AuthRepoImpl) Logout(token string) error {
	// read token list
	var tokens []dto.Authorize
	byteValue, _ := ioutil.ReadFile("./jsonDb/authorize.json")
	err := json.Unmarshal(byteValue, &tokens)
	if err != nil {
		return err
	}

	for idx, val := range tokens {
		if val.Token == token {
			tokens = append(tokens[:idx], tokens[idx+1:]...)
			tokensByte, _ := json.Marshal(tokens)
			err = ioutil.WriteFile("./jsonDb/authorize.json", tokensByte, 0644)
			return err
		}
	}

	return errors.New("token gk ada")
}
