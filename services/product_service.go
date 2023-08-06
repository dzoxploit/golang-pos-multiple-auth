package services

import (
	"gocommerce/models"
	"gocommerce/repositories"
	"strconv"

	"github.com/jinzhu/gorm" // Add strconv package for string to uint conversion
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		productRepository: repositories.NewProductRepository(db),
	}
}

func (s *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return s.productRepository.CreateProduct(product)
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	// Convert the string ID to uint
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	return s.productRepository.GetProductByID(uint(productID))
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.productRepository.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id string) error {
	// Convert the string ID to uint
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	return s.productRepository.DeleteProduct(uint(productID))
}

func (s *ProductService) ListProducts() ([]*models.Product, error) {
	products, err := s.productRepository.ListProducts()
	if err != nil {
		return nil, err
	}

	// Convert []models.Product to []*models.Product
	var productPointers []*models.Product
	for i := range products {
		productPointers = append(productPointers, &products[i])
	}

	return productPointers, nil
}