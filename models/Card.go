package models

import "gorm.io/gorm"

type Card struct {
	FirstName  string
	LastName   string
	CardNumber string
	gorm.Model
}

func (Card) TableName() string {
	return "customer_card"
}

// ToDO: Implement
func obfuscateCardNumber(cardNumber string) string {
	return cardNumber
}

// BeforeCreate is a hook that obfuscates a customer's card number before creation in the database
func (c *Card) BeforeCreate(tx *gorm.DB) error {
	c.CardNumber = obfuscateCardNumber(c.CardNumber)
	return nil
}
