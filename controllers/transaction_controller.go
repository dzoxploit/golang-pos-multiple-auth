package controllers

import (
	"gocommerce/models"
	"gocommerce/services"
	"gocommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

func NewTransactionController(db *gorm.DB) *TransactionController {
	return &TransactionController{
		transactionService: services.NewTransactionService(db),
	}
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
    var transaction models.Transaction
    if err := ctx.ShouldBindJSON(&transaction); err != nil {
        utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Create the transaction
    createdTransaction, err := c.transactionService.CreateTransaction(&transaction)
    if err != nil {
        utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create transaction")
        return
    }

    // Subtract the stock quantity of the products involved in the transaction
    err = c.transactionService.UpdateProductStock(createdTransaction)
    if err != nil {
        utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update product stock")
        return
    }

    utils.SendSuccessResponse(ctx, http.StatusCreated, "Transaction created successfully", createdTransaction)
}

func (c *TransactionController) CreateTransactionByUserID(ctx *gin.Context) {
    var transaction models.Transaction

	email, ok := ctx.Get("email")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	// Convert userID to uint (assuming it's stored as uint in the database)
	emailString, ok := email.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}


    if err := ctx.ShouldBindJSON(&transaction); err != nil {
        utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
        return
    }

    userID, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

    userIDUint, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	transaction.UserID = userIDUint

    // Create the transaction
    createdTransaction, err := c.transactionService.CreateTransactionByUserID(&transaction)
    if err != nil {
        utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create transaction")
        return
    }

    // Subtract the stock quantity of the products involved in the transaction
    err = c.transactionService.UpdateProductStock(createdTransaction)
    if err != nil {
        utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update product stock")
        return
    }


	receiptEmail := emailString // Replace with the actual email of the user
	err = c.transactionService.SendEmailAfterTransaction(createdTransaction, receiptEmail)

    utils.SendSuccessResponse(ctx, http.StatusCreated, "Transaction created successfully", createdTransaction)
}

func (c *TransactionController) ListTransactions(ctx *gin.Context) {
	transactions, err := c.transactionService.ListTransactions()
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch transactions")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusOK, "Transactions fetched successfully", transactions)
}

func (c *TransactionController) GenerateCSVTransactions(ctx *gin.Context) {
	// Fetch the list of transactions from the service
	transactions, err := c.transactionService.ListTransactionJoin()
	if err != nil {
		// Handle the error, return error response
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	// Call the CSV generator function to convert transactions to CSV data
	csvData, err := c.transactionService.GenerateCSV(transactions)
	if err != nil {
		// Handle the error, return error response
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSV"})
		return
	}

	// Set the response headers for CSV download
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename=transactions.csv")
	ctx.Data(http.StatusOK, "text/csv", csvData)
}

func(c *TransactionController) ListTransactionsByUserID(ctx *gin.Context) {
	// Get the user ID from the context set by the authentication middleware
	userID, ok := ctx.Get("uid")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	// Convert userID to uint (assuming it's stored as uint in the database)
	userIDUint, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	// Use the user ID to fetch transactions from the database
	// You can implement this logic using your TransactionService or Repository
	transactions, err := c.transactionService.GetTransactionsByUserID(userIDUint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}
