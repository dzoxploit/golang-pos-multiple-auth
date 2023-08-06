package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gocommerce/config"
	"gocommerce/controllers"
	"gocommerce/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func generateRandomKey() (string, error) {
	key := make([]byte, 32) // 32 bytes for a secure key
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

func main() {
	secretKey, err := generateRandomKey()
	if err != nil {
		fmt.Println("Failed to generate JWT secret key:", err)
		return
	}

	// Set the secret key as an environment variable
	os.Setenv("JWT_SECRET_KEY", secretKey)
	
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()

	// Set up the database connection
	db, err := config.NewDB()
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// Initialize the controllers
	authController := controllers.NewAuthController(db)
	productController := controllers.NewProductController(db)
	transactionController := controllers.NewTransactionController(db)


	
	// Register the API routes

	authRoutes := r.Group("/auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
			authRoutes.GET("/validate", authController.ValidateToken)
		}

	api := r.Group("/api")
	{

		userRoutes := api.Group("/users")
		userRoutes.Use(middlewares.Authenticate())
		{
			userRoutes.GET("/", productController.ListProducts)
			userRoutes.GET("/transactions", transactionController.ListTransactionsByUserID)
			userRoutes.POST("/transaction", transactionController.CreateTransactionByUserID)
		}

		adminRoutes := api.Group("/admins")
		adminRoutes.Use(middlewares.AuthenticateAdmin())
		{
			adminRoutes.GET("/", productController.ListProducts)
			adminRoutes.POST("/create", productController.CreateProduct)
			adminRoutes.GET("/:id", productController.GetProduct)
			adminRoutes.PUT("/:id", productController.UpdateProduct)
			adminRoutes.DELETE("/:id", productController.DeleteProduct)
			
			adminRoutes.POST("/transaction", transactionController.CreateTransaction)
			adminRoutes.GET("/transactions", transactionController.ListTransactions)
			adminRoutes.GET("/transactions/generate-csv", transactionController.GenerateCSVTransactions)
		}
	}

	// Start the server
	r.Run(":7000")
}
