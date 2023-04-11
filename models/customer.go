package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Customer entity definition
// The Customer table has a one-to-many relationship with Account and Card entities
type Customer struct {
	UserId    uuid.UUID `gorm:"not null;unique;type:uuid" json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Accounts  []Account `gorm:"foreignKey:user_id;references:user_id"`
	Cards     []Card    `gorm:"foreignKey:user_id;references:user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Customer) TableName() string {
	return "customer"
}

// BeforeCreate initializes a customer's unique ID
func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	c.UserId = uuid.New()
	return nil
}
