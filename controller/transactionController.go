package controller

import (
	"bank/entity"
	"bank/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewTransactionController(transactionService *service.TransactionService) TransactionController {
	return TransactionController{TransactionService: *transactionService}
}

type TransactionController struct {
	TransactionService service.TransactionService
}

func (controller *TransactionController) Route(router, auth *mux.Router) {
	// auth.HandleFunc("/transaction", controller.Create).Methods("POST")
	auth.HandleFunc("/transaction/{id}", controller.Create).Methods("POST")
}

func (controller *TransactionController) Create(w http.ResponseWriter, r *http.Request) {
	var transaction entity.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	var id string
	for _, value := range params {
		id = value
	}
	// user := context.Get(r, "user").(*dto.UserCredential)

	transaction.SetAccountId(id)

	err := controller.TransactionService.Create(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"berhasil transaksi"}`))
}
