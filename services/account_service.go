package services

import (
	"fmt"
	"github.com/google/uuid"
	"simple-bank-account/models"
	"simple-bank-account/repositories"
)

type AccountsService struct {
	repository repositories.AccountRepository
}

type SimpleBankAccountsService interface {
	CreateAccount(userId uuid.UUID) (*models.Account, error)
	WithdrawFromAccount(userId uuid.UUID, amount float64) (float64, error)
	TopUpAccount(userId uuid.UUID, amount float64) (float64, error)
}

func NewAccountsService(repository repositories.AccountRepository) *AccountsService {
	return &AccountsService{repository: repository}
}

func (s AccountsService) CreateAccount(userId uuid.UUID) (*models.Account, error) {
	account, err := s.repository.CreateAccount(userId)
	if err != nil {
		return &models.Account{}, fmt.Errorf("account creation failed with error: %w", err)
	}
	return account, nil
}

func (s AccountsService) WithdrawFromAccount(userId uuid.UUID, amount float64) (float64, error) {
	balance, err := s.repository.WithdrawFromAccount(userId, amount)
	if err != nil {
		return balance, fmt.Errorf("account withdrawal failed with error: %w", err)
	}
	return balance, nil
}

func (s AccountsService) TopUpAccount(userId uuid.UUID, amount float64) (float64, error) {
	balance, err := s.repository.WithdrawFromAccount(userId, amount)
	if err != nil {
		return balance, fmt.Errorf("account top-up failed with error: %w", err)
	}
	return balance, nil
}
