package repositories

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) CreateAdmin(admin *models.Admin) error {
	// Create the user in the database
	if err := r.db.Create(admin).Error; err != nil {
		return err
	}

	return nil
}

func (r *AdminRepository) GetAdminByEmail(email string) (*models.Admin, error) {
	// Get the user from the database based on the username
	var admin models.Admin
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}
