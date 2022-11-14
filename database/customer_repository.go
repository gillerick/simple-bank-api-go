package database

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"simple-bank-account/models"
)

// CustomerRepository exposes methods for performing customer-related db operations
type CustomerRepository interface {
	Add(models.Customer) (models.Customer, error)
	FindByUserId(customerId uuid.UUID) (models.Customer, error)
	Delete(customer models.Customer) error
}

// Add creates a new customer if they don't already exist in the database
func (r Repository) Add(customer models.Customer) (models.Customer, error) {
	result := r.db.pg.Model(models.Customer{}).Create(&customer)
	if err := result.Error; err != nil {
		// we check if the error is a postgres unique constraint violation
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
			return customer, errors.New("customer already exists")
		}
		return models.Customer{}, fmt.Errorf("customer could not be created %w", err)
	}
	return customer, nil
}

// FindByUserId searches a customer by their unique ID
func (r Repository) FindByUserId(customerId uuid.UUID) (models.Customer, error) {
	var customer models.Customer
	result := r.db.pg.Where(models.Customer{UserId: customerId}).First(&customer)
	// check if no record found.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Customer{}, errors.New("customer does not exist")
	}
	//Handle any other error
	if err := result.Error; err != nil {
		return models.Customer{}, fmt.Errorf("customer could not be retrieved %w", err)
	}
	return customer, nil
}

func (r Repository) Delete(customer models.Customer) error {
	err := r.db.pg.Delete(&customer).Error
	if err != nil {
		return fmt.Errorf("customer could not be deleted %w", err)
	}
	return nil
}

func NewCustomerRepository(database *DatabaseHandler) CustomerRepository {
	return &Repository{db: database}
}
