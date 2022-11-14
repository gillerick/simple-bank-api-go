package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountStatus string
type AccountType string

const (
	StatusActive    = AccountStatus("ACTIVE")
	StatusInactive  = AccountStatus("INACTIVE")
	StatusSuspended = AccountStatus("SUSPENDED")
	CurrentAccount  = AccountType("CURRENT")
	SavingsAccount  = AccountType("SAVINGS")

	MinimumWithdrawalAmount = 50 // Least amount that can be withdrawn from an account
	MinimumTopUpAmount      = 5  //Least amount that can be deposited into an account
)

// Account entity definition. It has a composite primary key of UserId, Status & Type.
// This allows us to have multiple accounts of different types and status for the same customer
type Account struct {
	Type             AccountType   `json:"account_type" gorm:"primaryKey"`
	Status           AccountStatus `gorm:"not null;default:ACTIVE;primaryKey"`
	AvailableBalance float64
	UserId           uuid.UUID `gorm:"primaryKey"`

	gorm.Model
}

func (Account) TableName() string {
	return "customer_account"
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
