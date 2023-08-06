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

func (r *TransactionRepository) CreateTransactionByUserID(transaction *models.Transaction) (*models.Transaction, error) {
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

func (r *TransactionRepository) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) ListTransactionJoin() ([]models.TransactionJoin, error) {
	var transactions []models.TransactionJoin
	if err := r.db.Table("transactions").
		Select("transactions.*, products.name as product, users.username").
		Joins("JOIN products ON products.id = transactions.product_id").
		Joins("JOIN users ON users.id = transactions.user_id").
		Scan(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

