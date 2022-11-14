package repositories

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"simple-bank-account/models"
)

type AccountRepository interface {
	CreateAccount(userId uuid.UUID, accountType models.AccountType) (*models.Account, error)
	WithdrawFromAccount(userId uuid.UUID, amount float64) (float64, error)
	TopUpAccount(userId uuid.UUID, amount float64) (float64, error)
}

func (r Repository) CreateAccount(userId uuid.UUID, accountType models.AccountType) (*models.Account, error) {
	newAccount := InitializeAccount(userId, accountType)
	err := r.db.pg.Model(&models.Account{}).FirstOrCreate(&newAccount).Error
	if err != nil {
		return nil, fmt.Errorf("error creating user account %v", err)
	}
	return &newAccount, nil
}

// WithdrawFromAccount deducts a specified amount from an account in a transactional operation and returns the new account balance or error
func (r Repository) WithdrawFromAccount(userId uuid.UUID, amount float64) (float64, error) {
	//Perform prerequisite checks before a withdrawal (1) account must have sufficient funds (2) Withdrawal amount is greater than the set minimum

	//Start a transaction to ensure consistency of data.
	//We are setting the isolation level to RepeatableRead to guarantee that both the data read has been committed & that another read within the scope of the transaction would yield the same result.
	tx := r.db.pg.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	//1. Find the account of the specified user
	var acc models.Account
	err := tx.Model(models.Account{}).Where(models.Account{UserId: userId}).Scan(&acc).Error
	if err != nil {
		//We are rolling back all the changes if an error occurs in any step of our transaction
		tx.Rollback()
		return acc.AvailableBalance, fmt.Errorf("withdrawal operation failed %v", err)
	}

	if amount < models.MinimumWithdrawalAmount {
		tx.Rollback()
		return acc.AvailableBalance, fmt.Errorf("withdrawal amount less than the minimum %d", models.MinimumWithdrawalAmount)
	}

	if acc.AvailableBalance < amount {
		tx.Rollback()
		return acc.AvailableBalance, errors.New("account balance is insufficient. top up and try again")
	}

	//2. Deduct the specified amount from the account
	newAmount := acc.Debit(amount)
	//3. Update the account balance
	_, err = UpdateBalance(newAmount, userId, tx)
	if err != nil {
		tx.Rollback()
		return acc.AvailableBalance, err
	}
	return acc.AvailableBalance, tx.Commit().Error

}

// TopUpAccount tops up a specified amount to an account in a transactional operation and returns the new account balance or error
func (r Repository) TopUpAccount(userId uuid.UUID, amount float64) (float64, error) {
	//Start a transaction to ensure consistency of data.
	//We are setting the isolation level to RepeatableRead to guarantee that both the data read has been committed & that another read within the scope of the transaction would yield the same result.
	tx := r.db.pg.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	//1. Find the account of the specified user
	var acc models.Account
	err := tx.Model(models.Account{}).Where(models.Account{UserId: userId}).Scan(&acc).Error
	if err != nil {
		return acc.AvailableBalance, fmt.Errorf("withdrawal operation failed %v", err)
	}

	//2. Perform prerequisite top up checks
	if amount < models.MinimumTopUpAmount {
		tx.Rollback()
		return acc.AvailableBalance, fmt.Errorf("topup amount less than the minimum %d", models.MinimumTopUpAmount)
	}

	newAmount := acc.Credit(amount)
	//3. Update account balance
	_, err = UpdateBalance(newAmount, userId, tx)
	if err != nil {
		tx.Rollback()
		return acc.AvailableBalance, err
	}
	return acc.AvailableBalance, tx.Commit().Error
}

func UpdateBalance(amount float64, userId uuid.UUID, db *gorm.DB) (*models.Account, error) {
	var acc models.Account
	result := db.Model(models.Account{}).Where(models.Account{UserId: userId}).Updates(models.Account{AvailableBalance: amount}).Scan(&acc)
	if err := result.Error; err != nil {
		return &models.Account{}, fmt.Errorf("update operation failed %w", err)
	}
	return &acc, nil
}

// InitializeAccount initializes and returns a new account
func InitializeAccount(userId uuid.UUID, accountType models.AccountType) models.Account {
	return models.Account{
		Type:             accountType,
		AvailableBalance: 0,
		UserId:           userId,
	}

}

func NewAccountRepository(database *DatabaseHandler) AccountRepository {
	return &Repository{db: database}
}
