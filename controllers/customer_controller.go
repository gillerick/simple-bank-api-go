package controllers

import (
	"net/http"
	"simple-bank-account/services"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func (c CustomerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

type CustomerController interface {
	//UpdateBalance(amount float64, userId uuid.UUID) (*models.Account, error)
}

func NewCustomerController(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}
