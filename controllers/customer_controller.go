package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"net/http"
	"simple-bank-account/models"
	"simple-bank-account/services"
)

type CustomerController struct {
	customerService services.CustomerService
}

func (c CustomerController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if RequestValidator(request) != nil {
		http.Error(writer, "bad user input", http.StatusBadRequest)
	}

	//Routes all POST requests to the saveCustomer service method
	if request.Method == http.MethodPost {
		var customer models.Customer

		//Map request body from JSON to Customer entity
		err := json.NewDecoder(request.Body).Decode(&customer)
		if err != nil {
			http.Error(writer, "bad user input", http.StatusBadRequest)
			return
		}
		log.Infof("received new create customer request %s", customer)

		account, err := c.customerService.SaveCustomer(customer)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(account)
		if err != nil {
			http.Error(writer, "json encoding failed", http.StatusUnprocessableEntity)
		}
		return
	}

	//Retrieves customerId path parameter from the http request
	customerId := mux.Vars(request)["customerId"]
	if request.Method == http.MethodGet {
		retrievedCustomer, err := c.customerService.FindCustomerByUserId(uuid.MustParse(customerId))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Infof("retrieving customer details for customer id: %s", customerId)

		err = json.NewEncoder(writer).Encode(retrievedCustomer)
		if err != nil {
			http.Error(writer, "json encoding failed", http.StatusUnprocessableEntity)
		}
		return

	}

}

func NewCustomerController(customerService services.CustomerService) *CustomerController {
	return &CustomerController{customerService: customerService}
}
