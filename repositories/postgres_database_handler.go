package repositories

import (
	"gorm.io/gorm"
)

type DatabaseHandler struct {
	pg *gorm.DB
}

type Repository struct {
	db *DatabaseHandler
}

func NewDatabaseHandler(db *gorm.DB) *DatabaseHandler {
	return &DatabaseHandler{pg: db}
}
