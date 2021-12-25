package controller

import (
	"bank/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewCustomerController(customerService *service.CustomerService) CustomerController {
	return CustomerController{
		CustomerService: *customerService,
	}
}

type CustomerController struct {
	CustomerService service.CustomerService
}

func (controller *CustomerController) Route(router, auth *mux.Router) {
	auth.HandleFunc("/customer", controller.GetAll).Methods("GET")
}

func (controller *CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	customers := controller.CustomerService.GetAll()

	message, err := json.Marshal(&customers)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)

}
