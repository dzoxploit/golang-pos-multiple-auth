package config

import (
	"gocommerce/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDB() (*gorm.DB, error) {
	// Replace the database connection details with your actual MySQL database credentials
	dbURI := "root:@tcp(localhost:3306)/gopos?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// Enable debug mode during development for logging SQL queries
	db.LogMode(true)

	// AutoMigrate will automatically create the database tables based on the models
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Transaction{})

	return db, nil
}
