package services

import "errors"

type MockAuthService struct{}

func (s *MockAuthService) Login(email, username, password string) (string, error) {
	// Implement the login functionality for the mock service.
	// For testing purposes, you can return a fixed access token or error.
	var ErrInvalidCredentials = errors.New("invalid credentials")
	
	if username == "testuser" && password == "testpassword" {
		return "mock_access_token", nil
	}
	return "", ErrInvalidCredentials
}