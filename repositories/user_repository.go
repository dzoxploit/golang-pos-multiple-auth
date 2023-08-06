package repositories

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	// Create the user in the database
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	// Get the user from the database based on the username
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
