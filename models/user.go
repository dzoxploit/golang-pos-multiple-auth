package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Add any additional fields or methods as needed for User model
