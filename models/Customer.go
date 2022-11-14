package models

import "gorm.io/gorm"

type Customer struct {
	FirstName string
	LastName  string

	gorm.Model
}

func (Customer) TableName() string {
	return "customer"
}
