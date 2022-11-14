package controllers

import (
	"github.com/google/uuid"
	"simple-bank-account/models"
	"simple-bank-account/services"
)

type AccountsHandler struct {
	accountsService services.AccountsService
}

type AccountsController interface {
	CreateAccount(userId uuid.UUID) (*models.Account, error)
	WithdrawFromAccount(userId uuid.UUID, amount float64) (float64, error)
	TopUpAccount(userId uuid.UUID, amount float64) (float64, error)
	UpdateBalance(amount float64, userId uuid.UUID) (*models.Account, error)
}

func NewAccountsController(accountsService services.AccountsService) *AccountsHandler {
	return &AccountsHandler{accountsService: accountsService}
}
