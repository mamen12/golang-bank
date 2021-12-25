package controller

import (
	"bank/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}

type AuthController struct {
	AuthService service.AuthService
}

func (controller *AuthController) Route(router *mux.Router) {
	router.HandleFunc("/login", controller.Login).Methods("POST")
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var login = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := controller.AuthService.Login(login.Username, login.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//save to authorize lis

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token":"` + token + `"}`))
}
func (controller *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
}
