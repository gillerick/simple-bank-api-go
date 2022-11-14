package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	UserId    uuid.UUID
	FirstName string
	LastName  string

	gorm.Model
}

func (Customer) TableName() string {
	return "customer"
}
