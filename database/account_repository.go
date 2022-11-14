package database

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"simple-bank-account/models"
)

type AccountRepository interface {
	CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error)
	WithdrawAmount(amount float64) (float64, error)
	TopUpAccount(amount float64)
	UpdateBalance(amount float64, userId uuid.UUID) (*models.Account, error)
}

func NewAccountRepository(database *DatabaseHandler) AccountRepository {
	return &Repository{db: database}
}

func (d *DatabaseHandler) CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error) {
	newAccount := InitializeAccount(userId)
	err := d.pg.Model(&models.Account{}).FirstOrCreate(&newAccount)
	if err != nil {
		return nil, fmt.Errorf("error creating user account %v", err)
	}
	return &newAccount, nil
}

// WithdrawAmount deducts a specified amount from an account in a transactional operation and returns the withdrawn account balance or error
func (r Repository) WithdrawAmount(userId uuid.UUID, amount float64) (float64, error) {
	//Perform prerequisite checks before a withdrawal (1) account must have sufficient funds
	var acc models.Account
	err := r.db.pg.Model(models.Account{}).Where(models.Account{UserId: userId}).Scan(&acc).Error
	if err != nil {
		return 0, fmt.Errorf("withdrawal operation failed %v", err)
	}

	if acc.AvailableBalance < amount {
		return 0, errors.New("account balance is insufficient. top up and try again")
	}

	//2. Deduct the specified amount from the account
	acc.Debit(amount)
	var newAccount models.Account
	//3. Update the account balance
	err = r.db.pg.Model(models.Account{}).Where(models.Account{UserId: userId}).Updates(models.Account{AvailableBalance: amount}).Scan(&newAccount).Error
	if err != nil {
		return 0, fmt.Errorf("error updating account balance %v", err)
	}
	return amount, nil

}

//func (r Repository) TopUpAccount(amount float64) (float64, error) {
//
//}
