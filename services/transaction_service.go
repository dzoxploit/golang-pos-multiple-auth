package services

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"gocommerce/models"
	"gocommerce/repositories"
	"strconv"

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

func (s *TransactionService) CreateTransactionByUserID(transaction *models.Transaction) (*models.Transaction, error) {
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

func (s *TransactionService) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	return s.transactionRepository.GetTransactionsByUserID(userID)
}

func (s *TransactionService) generateCSV(transactions []*models.Transaction) ([]byte, error) {
	// Create a buffer to store CSV data
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write the CSV header
	header := []string{"Transaction ID", "Description", "Amount", "Date"}
	writer.Write(header)

	// Write transaction data
	for _, transaction := range transactions {
		row := []string{
			strconv.Itoa(int(transaction.ID)),
			strconv.FormatFloat(transaction.Amount, 'f', 2, 64),
			transaction.OrderDate.Format("2006-01-02"),
		}
		writer.Write(row)
	}

	// Flush the writer to make sure all data is written to the buffer
	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

