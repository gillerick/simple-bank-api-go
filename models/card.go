package models

import (
	"github.com/ggwhite/go-masker"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Card entity definition. It has a composite primary key of CardId, CardNumber and UserId, allowing a single user to
// have multiple cards of distinct numbers
type Card struct {
	CardId       uuid.UUID `gorm:"type:uuid;unique;primaryKey"`
	CustomerName string
	ExpiryDate   string
	CardNumber   string    `gorm:"primaryKey"`
	UserId       uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Card) TableName() string {
	return "customer_card"
}

// BeforeCreate is a hook that masks a customer's card number & generates a UUID before creation of card entry in the database
// For instance, a MasterCard of number 1234567890123456 will be saved as 123456******3456
func (c *Card) BeforeCreate(tx *gorm.DB) error {
	c.CardId = uuid.New()
	c.CardNumber = masker.CreditCard(c.CardNumber)
	return nil
}
