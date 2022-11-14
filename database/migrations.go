package database

import (
	"github.com/prometheus/common/log"
	"gorm.io/gorm"
	"simple-bank-account/models"
)

func (d *DatabaseHandler) Migrate(pg *gorm.DB) error {
	log.Info("running db migrations")
	//ToDo: Add card table
	err := d.pg.AutoMigrate(models.Account{}, models.Account{}, models.Card{})
	if err != nil {
		log.Fatalf("could not run db migrations")
	}

	log.Info("db migrations ran successfully")
	return nil
}
