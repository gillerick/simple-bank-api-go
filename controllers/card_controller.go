package controllers

import (
	"net/http"
	"simple-bank-account/services"
)

type CardHandler struct {
	cardService services.CardService
}

func (c CardHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

type CardController interface {
}

func NewCardController(cardService services.CardService) *CardHandler {
	return &CardHandler{cardService: cardService}
}
