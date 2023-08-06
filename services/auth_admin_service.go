package services

import (
	"errors"
	"gocommerce/models"
	"gocommerce/repositories"
	"gocommerce/utils"

	"github.com/jinzhu/gorm"
)

var (
	ErrAdminNotFound    = errors.New("admin not found")
	ErrAdminInvalidPassword = errors.New("invalid password")
	ErrAdminInvalidCredentials = errors.New("invalid credentials")
)

type AuthAdminService struct {
	adminRepository *repositories.AdminRepository
	db *gorm.DB
}

func NewAuthAdminService(db *gorm.DB) *AuthAdminService {
	return &AuthAdminService{
		adminRepository: repositories.NewAdminRepository(db),
		db : db,
	}
}

func (s *AuthAdminService) RegisterAdmin(admin *models.Admin) error {
	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return err
	}

	admin.Password = hashedPassword
	return s.adminRepository.CreateAdmin(admin)
}

func (s *AuthAdminService) LoginAdmin(email, password string) (string, error) {
	// Get the user from the database by username
	admin, err := s.adminRepository.GetAdminByEmail(email)
	if err != nil {
		// Handle the "user not found" case
		if err == ErrUserNotFound {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	
	if err != nil {
		return "", nil 
	}

	// Check if the provided password matches the hashed password in the database
	passwordIsValid := utils.CheckPasswordHashAdmin(password, admin.Password)

	if !passwordIsValid {
		// Return the custom error for invalid password
		return "", ErrInvalidPassword
	}

	// Generate the JWT token
	token, err := utils.GenerateTokenAdmin(uint(admin.ID), admin.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthAdminService) GetAdminByID(userID uint) (*models.Admin, error) {
	var admin models.Admin
	err := s.db.Where("id = ?", userID).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
