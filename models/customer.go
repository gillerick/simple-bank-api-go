package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Customer entity definition
// The Customer table has a one-to-many relationship with Account and Card entities
type Customer struct {
	UserId    uuid.UUID `gorm:"not null;unique;type:uuid"`
	FirstName string
	LastName  string
	Email     string  `gorm:"not null;unique"`
	Account   Account `gorm:"foreignKey:user_id;references:user_id"`
	Card      []Card  `gorm:"foreignKey:user_id;references:user_id"`

	gorm.Model
}

func (Customer) TableName() string {
	return "customer"
}
