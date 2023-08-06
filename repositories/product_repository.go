package repositories

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	// Create the product in the database
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	// Get the product from the database based on the ID
	var product models.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	// Update the product in the database
	if err := r.db.Save(product).Error; err != nil {
		return err
	}

	return nil
}


func (r *ProductRepository) DeleteProduct(id uint) error {
	// Delete the product from the database based on the ID
	if err := r.db.Where("id = ?", id).Delete(models.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) ListProducts() ([]models.Product, error) {
	// Get all products from the database
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
