package services

import (
	"fmt"
	"simple-bank-account/models"
	"simple-bank-account/repositories"
)

type CardService struct {
	repository repositories.CardRepository
}

type SimpleBankCardService interface {
	SaveCard(card models.Card) (models.Card, error)
	DeleteCard(customer models.Card) error
}

func NewCardService(repository repositories.CardRepository) *CardService {
	return &CardService{repository: repository}
}

func (s CardService) SaveCard(card models.Card) (models.Card, error) {
	savedCard, err := s.repository.SaveCard(card)
	if err != nil {
		return models.Card{}, fmt.Errorf("save card failed with error: %w", err)
	}
	return savedCard, nil
}

func (s CardService) DeleteCard(card models.Card) error {
	err := s.repository.DeleteCard(card)
	if err != nil {
		return fmt.Errorf("card deletion failed with error: %w", err)
	}
	return nil
}
