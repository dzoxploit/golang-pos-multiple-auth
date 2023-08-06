package services

import (
	"fmt"
	"gocommerce/models"
	"gocommerce/repositories"

	"github.com/jinzhu/gorm"
)

type TransactionService struct {
	transactionRepository *repositories.TransactionRepository
	productRepository *repositories.ProductRepository
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		transactionRepository: repositories.NewTransactionRepository(db),
		productRepository: repositories.NewProductRepository(db),
	}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return s.transactionRepository.CreateTransaction(transaction)
}

func (s *TransactionService) UpdateProductStock(transaction *models.Transaction) error {
    // Get the product by its ID
    product, err := s.productRepository.GetProductByID(transaction.ProductID)
    if err != nil {
        return err
    }

	fmt.Println(product.Quantity)
    // Subtract the transaction quantity from the product's stock
    product.Quantity -= transaction.Quantity

    // Update the product's stock in the repository
    err = s.productRepository.UpdateProduct(product)
    if err != nil {
        return err
    }

    return nil
}

func (s *TransactionService) ListTransactions() ([]*models.Transaction, error) {
	transactions, err := s.transactionRepository.ListTransactions()
	if err != nil {
		return nil, err
	}

	// Convert []models.Transaction to []*models.Transaction
	var transactionPointers []*models.Transaction
	for i := range transactions {
		transactionPointers = append(transactionPointers, &transactions[i])
	}

	return transactionPointers, nil
}
