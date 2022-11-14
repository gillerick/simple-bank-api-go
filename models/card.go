package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Card entity definition
type Card struct {
	FirstName  string
	LastName   string
	CardNumber string
	UserId     uuid.UUID
	gorm.Model
}

func (Card) TableName() string {
	return "customer_card"
}

// ToDO: Implement
func obfuscateCardNumber(cardNumber string) string {
	return cardNumber
}

// BeforeCreate is a hook that obfuscates a customer's card number before creation in the repositories
func (c *Card) BeforeCreate(tx *gorm.DB) error {
	c.CardNumber = obfuscateCardNumber(c.CardNumber)
	return nil
}
