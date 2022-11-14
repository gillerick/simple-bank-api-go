package database

import (
	"fmt"
	"github.com/google/uuid"
	"simple-bank-account/models"
)

type AccountRepository interface {
	CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error)
	WithdrawAmount(amount float64)
	DepositAmount(amount float64)
	UpdateBalance(amount float64, userId uuid.UUID) (*models.Account, error)
}

type repository struct {
	db *DatabaseHandler
}

func NewAccountRepository(database *DatabaseHandler) AccountRepository {
	return &repository{db: database}
}

func (d *DatabaseHandler) CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error) {
	newAccount := InitializeAccount(userId)
	err := d.pg.Model(&models.Account{}).FirstOrCreate(&newAccount)
	if err != nil {
		return nil, fmt.Errorf("error creating user account %v", err)
	}
	return &newAccount, nil
}
