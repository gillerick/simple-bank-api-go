package database

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"simple-bank-account/models"
)

type DatabaseHandler struct {
	pg *gorm.DB
}

type Repository struct {
	db *DatabaseHandler
}

//ToDo: Expose db operations as interfaces
//1. Withdraw amount
//2. top up account

// WithdrawAmount deducts a specified amount from an account in a transactional operation
func (d *DatabaseHandler) WithdrawAmount(userId uuid.UUID, amount float64) error {
	//Perform prerequisite checks before a withdrawal (1) account must have sufficient funds
	var acc models.Account
	err := d.pg.Model(models.Account{}).Where(models.Account{UserId: userId}).Scan(&acc).Error
	if err != nil {
		return fmt.Errorf("withdrawal operation failed %v", err)
	}

	if acc.AvailableBalance < amount {
		return errors.New("account balance is insufficient. top up and try again")
	}

	//2. Deduct the specified amount from the account
	acc.Debit(amount)
	var newAccount models.Account
	//err := d.pg.Model(models.Account{}).
	err = d.pg.Model(models.Account{}).Where(models.Account{UserId: userId}).Updates(models.Account{AvailableBalance: amount}).Scan(&newAccount).Error
	if err != nil {
		return fmt.Errorf("error updating account balance %v", err)
	}
	return nil

}

//func (d *DatabaseHandler) CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error) {
//	newAccount := InitializeAccount(userId)
//	err := d.pg.Model(&models.Account{}).FirstOrCreate(&newAccount)
//	if err != nil {
//		return nil, fmt.Errorf("error creating user account %v", err)
//	}
//	return &newAccount, nil
//}

func (d *DatabaseHandler) TopUpAccount(amount float64) {

}

func (d *DatabaseHandler) UpdateBalance(amount float64, userId uuid.UUID) (*models.Account, error) {
	var newAccount models.Account
	err := d.pg.Model(models.Account{}).Where(models.Account{UserId: userId}).Updates(models.Account{AvailableBalance: amount}).Scan(&newAccount).Error
	if err != nil {
		return nil, fmt.Errorf("error updating account balance %v", err)
	}
	return &newAccount, nil
}

// InitializeAccount initializes and returns a new account
func InitializeAccount(userId uuid.UUID) models.Account {
	accountId := uuid.New()
	return models.Account{
		Id:               accountId,
		AvailableBalance: 0,
		UserId:           userId,
	}

}

func NewDatabaseHandler(db *gorm.DB) *DatabaseHandler {
	return &DatabaseHandler{pg: db}
}
