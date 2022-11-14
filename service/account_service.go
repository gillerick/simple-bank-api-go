package service

import (
	"fmt"
	"github.com/google/uuid"
	"simple-bank-account/models"
)

type Repository interface {
	CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error)
	WithdrawAmount(amount float64)
	DepositAmount(amount float64)
	UpdateBalance(amount float64, userId uuid.UUID) (*models.Account, error)
}

//type repository struct {
//	db *postgres.DatabaseHandler
//}

type AccountService struct {
	repository Repository
}

type CustomerService struct {
	repository Repository
}

func NewCustomerService(repository Repository) *CustomerService {
	return &CustomerService{repository: repository}
}

func NewAccountService(repository Repository) *AccountService {
	return &AccountService{repository: repository}
}

// CreateAccount creates an account and returns it if successful, otherwise returns an error
func (d *AccountService) CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error) {
	account, err := d.repository.CreateAccount(userId, firstName, LastName)
	if err != nil {
		return nil, fmt.Errorf("account could not be created %v", err)
	}
	return account, nil
}

// WithdrawAmount
func WithdrawAmount(amount float64) {

}

func DepositAmount(amount float64) {

}
