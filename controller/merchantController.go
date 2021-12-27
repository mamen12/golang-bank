package controller

import (
	"bank/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMerchantController(merchantService *service.MerchantService) MerchantController {
	return MerchantController{
		MerchantService: *merchantService,
	}
}

type MerchantController struct {
	MerchantService service.MerchantService
}

func (controller *MerchantController) Route(router, auth *mux.Router) {
	auth.HandleFunc("/merchant", controller.GetAll).Methods("GET")
	auth.HandleFunc("/merchant/{id}", controller.GetMerchantId).Methods("GET")
}

func (controller *MerchantController) GetAll(w http.ResponseWriter, r *http.Request) {
	merchants := controller.MerchantService.GetAll()

	message, err := json.Marshal(&merchants)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)

}

func (controller *MerchantController) GetMerchantId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id string
	for _, value := range params {
		id = value
	}
	merchants := controller.MerchantService.GetMerchantId(id)

	message, err := json.Marshal(&merchants)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)

}
