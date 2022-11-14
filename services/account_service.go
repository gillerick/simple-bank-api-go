package services

import (
	"github.com/google/uuid"
	"simple-bank-account/models"
)

type DataStore interface {
	CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error)
	WithdrawAmount(amount float64)
	DepositAmount(amount float64)
	UpdateBalance(amount float64, userId uuid.UUID) (*models.Account, error)
}

type AccountService struct {
	repository DataStore
}

func NewAccountService(repository DataStore) *AccountService {
	return &AccountService{repository: repository}
}

// WithdrawAmount
func WithdrawAmount(amount float64) {

}

func DepositAmount(amount float64) {

}
