package controller

import (
	"bank/service"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAccountController(accountService *service.AccountService) AccountController {
	return AccountController{
		AccountService: *accountService,
	}
}

type AccountController struct {
	AccountService service.AccountService
}

func (controller *AccountController) Route(router, auth *mux.Router) {
	auth.HandleFunc("/account", controller.GetAll).Methods("GET")
	auth.HandleFunc("/account/{id}", controller.GetAccountId).Methods("GET")
}

func (controller *AccountController) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts := controller.AccountService.GetAll()

	message, err := json.Marshal(&accounts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)

}

func (controller *AccountController) GetAccountId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id string
	for _, value := range params {
		id = value
	}
	accounts := controller.AccountService.GetAccountId(id)

	if accounts.GetId() == "" {
		errMessage := errors.New("access ilegal, data tidak ditemukan")
		http.Error(w, errMessage.Error(), http.StatusNotFound)
		return
	}

	message, err := json.Marshal(&accounts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)

}
