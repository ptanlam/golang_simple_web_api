package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ptanlam/golang_simple_web_api/app/handlers"
	"github.com/ptanlam/golang_simple_web_api/app/repositories"
	"github.com/ptanlam/golang_simple_web_api/app/services"
)

func Start() {
	router := mux.NewRouter()

	cr := repositories.NewCustomerRepositoryStub()
	cs := services.NewCustomerService(cr)
	ch := handlers.NewCustomerHandler(cs)

	router.
		HandleFunc("/customers", ch.GetAllCustomers).
		Methods(http.MethodGet)
	router.
		HandleFunc("/customers/{customerId}", handlers.GetCustomerById).
		Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}
