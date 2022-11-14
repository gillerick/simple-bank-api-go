package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountStatus string

const (
	StatusActive            = AccountStatus("ACTIVE")
	StatusInactive          = AccountStatus("INACTIVE")
	StatusSuspended         = AccountStatus("SUSPENDED")
	MinimumWithdrawalAmount = 50 // Least amount that can be withdrawn from an account
	MinimumTopUpAmount      = 5  //Least amount that can be deposited into an account
)

// Account entity definition
type Account struct {
	Id               uuid.UUID
	Status           AccountStatus `gorm:"not null;default:ACTIVE"`
	AvailableBalance float64
	UserId           uuid.UUID

	gorm.Model
}

func (Account) TableName() string {
	return "account"
}

// Credit adds a specified amount to the account and returns the available balance
func (acc Account) Credit(amount float64) float64 {
	return acc.AvailableBalance + amount

}

// Debit subtracts a specified amount to the account and returns the available balance
func (acc Account) Debit(amount float64) float64 {
	return acc.AvailableBalance - amount
}

// IsBalanceLessThanAmount checks whether there is sufficient balance in account
func (acc Account) IsBalanceLessThanAmount(amount float64) bool {
	return acc.AvailableBalance < amount
}
