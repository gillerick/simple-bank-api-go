package repositories

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"simple-bank-account/models"
)

// CardRepository exposes methods for performing customer-related db operations
type CardRepository interface {
	SaveCard(cardId models.Card) (models.Card, error)
	DeleteCard(customer models.Card) error
}

// SaveCard creates a new card if it doesn't already exist in the database
func (r Repository) SaveCard(card models.Card) (models.Card, error) {
	result := r.db.pg.Model(models.Card{}).Create(&card)
	if err := result.Error; err != nil {
		// we check if the error is a postgres unique constraint violation
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
			return card, errors.New("card already exists")
		}
		return models.Card{}, fmt.Errorf("card could not be created %w", err)
	}
	return card, nil
}

// DeleteCard deletes a specified card from the database
func (r Repository) DeleteCard(card models.Card) error {
	err := r.db.pg.Delete(&card).Error
	if err != nil {
		return fmt.Errorf("card could not be deleted %w", err)
	}
	return nil
}

func NewCardRepository(database *DatabaseHandler) CardRepository {
	return &Repository{db: database}
}
