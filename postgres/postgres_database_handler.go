package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"simple-bank-account/models"
)

type DatabaseHandler struct {
	pg *gorm.DB
}

//ToDo: Expose database operations as interfaces
//1. Withdraw amount
//2. top up account

func (d *DatabaseHandler) WithdrawAmount(amount float64) {

}

func (d *DatabaseHandler) CreateAccount(userId uuid.UUID, firstName string, LastName string) (*models.Account, error) {
	newAccount := InitializeAccount(userId)
	err := d.pg.Model(&models.Account{}).FirstOrCreate(&newAccount)
	if err != nil {
		return nil, fmt.Errorf("error creating user account %v", err)
	}
	return &newAccount, nil
}

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
