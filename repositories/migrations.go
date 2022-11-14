package repositories

import (
	"github.com/prometheus/common/log"
	"simple-bank-account/models"
)

func (d *DatabaseHandler) Migrate() error {
	log.Info("running db migrations")
	err := d.pg.AutoMigrate(models.Customer{}, models.Account{}, models.Card{})
	if err != nil {
		log.Fatalf("could not run db migrations")
	}

	log.Info("db migrations ran successfully")
	return nil
}
