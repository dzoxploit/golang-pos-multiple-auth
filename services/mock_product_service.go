package services

import "gocommerce/models"

type MockProductService struct {
	// Your mock implementation...
}

func (s *MockProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	// Your mock implementation for CreateProduct...
	// For testing purposes, you can return a fixed product or error.
	return &models.Product{
		ID:       1,
		Name:     "Test Product",
		Price:    12.34,
		Quantity: 100,
	}, nil
}
