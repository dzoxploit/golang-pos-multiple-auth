package services

import (
	"errors"
	"gocommerce/models"
	"gocommerce/repositories"
	"gocommerce/utils"

	"github.com/jinzhu/gorm"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthService struct {
	userRepository *repositories.UserRepository
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		userRepository: repositories.NewUserRepository(db),
		db : db,
	}
}

func (s *AuthService) Register(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepository.CreateUser(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	// Get the user from the database by username
	user, err := s.userRepository.GetUserByUsername(username)
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
	passwordIsValid := utils.CheckPasswordHash(password, user.Password)

	if !passwordIsValid {
		// Return the custom error for invalid password
		return "", ErrInvalidPassword
	}

	// Generate the JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := s.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
