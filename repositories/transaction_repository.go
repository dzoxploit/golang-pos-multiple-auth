package repositories

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	// Create the transaction in the database
	if err := r.db.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}


func (r *TransactionRepository) ListTransactions() ([]models.Transaction, error) {
	// Get all transactions from the database
	var transactions []models.Transaction
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
