package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	var err error
	dbDSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable")
	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not create a database connection %v", err)
	}
	return db, nil
}
