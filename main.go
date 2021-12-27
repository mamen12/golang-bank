package main

import (
	"bank/controller"
	"bank/middleware"
	"bank/repo"
	"bank/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"Origin", "Accept", "Keep-Alive", "User-Agent", "If-Modified-Since", "Cache-Control", "Referer", "Authorization", "Content-Type", "X-Requested-With"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "HEAD"})

	auth := router.PathPrefix("").Subrouter()
	authorize := middleware.NewAuthorize()
	auth.Use(authorize.Authentication)

	//
	customerRepo := repo.NewCustomerRepo()
	customerService := service.NewCustomerService(&customerRepo)
	customerController := controller.NewCustomerController(&customerService)
	customerController.Route(router, auth)
	// var customerController controller.CustomerController
	// customerController.Route(router)

	//
	authRepo := repo.NewAuthRepo()
	authService := service.NewAuthService(&authRepo)
	authController := controller.NewAuthController(&authService)
	authController.Route(router, auth)

	//
	transactionRepo := repo.NewTransactionRepo()
	transactionService := service.NewTransactionService(&transactionRepo)
	transactionController := controller.NewTransactionController(&transactionService)
	transactionController.Route(router, auth)

	//
	merchantRepo := repo.NewMerchantRepo()
	merchantService := service.NewMerchantService(&merchantRepo)
	merchantController := controller.NewMerchantController(&merchantService)
	merchantController.Route(router, auth)

	//
	accountRepo := repo.NewAccountRepo()
	accountService := service.NewAccountService(&accountRepo)
	accountController := controller.NewAccountController(&accountService)
	accountController.Route(router, auth)

	fmt.Println("Server running at :8083")
	log.Fatal(http.ListenAndServe(":8083", handlers.CORS(origins, headers, methods)(router)))

}
