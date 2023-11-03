package routes

import (
	"github.com/dimasmds/crm-backend-udacity/pkg/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func GetConsumersRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/customers", handlers.GetCustomersHandler).Methods("GET")
	router.HandleFunc("/customers", handlers.PostCustomerHandler).Methods("POST")
	router.HandleFunc("/customers/{id}", handlers.GetCustomerByIdHandler).Methods("GET")
	router.HandleFunc("/customers/{id}", handlers.UpdateCustomerByIdHandler).Methods("PUT")
	router.HandleFunc("/customers/{id}", handlers.DeleteCustomerByIdHandler).Methods("DELETE")

	router.PathPrefix("/").Handler(http.StripPrefix("/", handlers.GetStaticFileHandler()))

	return router
}
