package services

import (
	"fmt"
	"github.com/google/uuid"
	"simple-bank-account/models"
	"simple-bank-account/repositories"
)

type CustomerService struct {
	repository repositories.CustomerRepository
}

type SimpleBankCustomerService interface {
	SaveCustomer(models.Customer) (models.Customer, error)
	FindCustomerByUserId(customerId uuid.UUID) (models.Customer, error)
	DeleteCustomer(customer models.Customer) error
}

func (s CustomerService) SaveCustomer(customer models.Customer) (models.Customer, error) {
	savedCustomer, err := s.repository.SaveCustomer(customer)
	if err != nil {
		return models.Customer{}, fmt.Errorf("customer creation failed with error: %w", err)
	}
	return savedCustomer, nil
}

func (s CustomerService) FindCustomerByUserId(customerId uuid.UUID) (models.Customer, error) {
	balance, err := s.repository.FindCustomerByUserId(customerId)
	if err != nil {
		return balance, fmt.Errorf("account retrieval failed with error: %w", err)
	}
	return balance, nil
}

func (s CustomerService) DeleteCustomer(customer models.Customer) error {
	err := s.repository.DeleteCustomer(customer)
	if err != nil {
		return fmt.Errorf("customer deletion failed with error: %w", err)
	}
	return nil
}

func NewCustomerService(repository repositories.CustomerRepository) *CustomerService {
	return &CustomerService{repository: repository}
}
