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

type AccountsController struct {
	accountsService services.AccountsService
}

func (a AccountsController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//Retrieves userId path parameter from the http request
	userId := mux.Vars(request)["userId"]
	if RequestValidator(request) != nil {
		http.Error(writer, "bad user input", http.StatusBadRequest)
	}

	//Routes all POST requests to the createAccount service method
	if request.Method == http.MethodPost {
		var account models.Account

		//Map request body from JSON to Account entity
		err := json.NewDecoder(request.Body).Decode(&account)
		if err != nil {
			http.Error(writer, "bad user input", http.StatusBadRequest)
			return
		}

		log.Infof("received new create account request %s", account)

		createdAcc, err := a.accountsService.CreateAccount(uuid.MustParse(userId), account.Type)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Infof("creating account for user id: %s", userId)

		err = json.NewEncoder(writer).Encode(createdAcc)
		if err != nil {
			http.Error(writer, "json encoding failed", http.StatusUnprocessableEntity)
		}
		return
	}

}

func NewAccountsController(accountsService services.AccountsService) *AccountsController {
	return &AccountsController{accountsService: accountsService}
}

func RequestValidator(r *http.Request) error {
	//ToDo: Implement
	return nil
}
