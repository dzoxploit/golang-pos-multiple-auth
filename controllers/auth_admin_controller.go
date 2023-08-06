package controllers

import (
	"fmt"
	"gocommerce/models"
	"gocommerce/services"
	"gocommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AuthAdminController struct {
	authAdminService *services.AuthAdminService
}

func NewAuthAdminController(db *gorm.DB) *AuthAdminController {
	return &AuthAdminController{
		authAdminService: services.NewAuthAdminService(db),
	}
}

func NewAuthAdminControllerTest(authAdminService *services.AuthAdminService) *AuthAdminController {
	return &AuthAdminController{
		authAdminService: authAdminService,
	}
}

func (c *AuthAdminController) RegisterAdmin(ctx *gin.Context) {
	var admin models.Admin
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := c.authAdminService.RegisterAdmin(&admin)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create admin")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusCreated, "Admin created successfully", nil)
}

func (c *AuthAdminController) LoginAdmin(ctx *gin.Context) {
	var admin models.Admin
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}


	// Authenticate the user
	jwtToken, err := c.authAdminService.LoginAdmin(admin.Email, admin.Password)
	if err != nil {
		if err == services.ErrUserNotFound {
			utils.SendErrorResponse(ctx, http.StatusUnauthorized, "User not found")
		} else if err == services.ErrInvalidPassword {
			utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid password")
		} else {
			utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to authenticate")
		}
		return
	}

	// Include the token in the response
	utils.SendSuccessResponse(ctx, http.StatusOK, "Login successful", gin.H{"token": jwtToken})
}

func (c *AuthAdminController) ValidateTokenAdmin(ctx *gin.Context) {
	// Get the token from the request header
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Authorization token missing")
		return
	}

	fmt.Println(tokenString)

	// Validate the token using the ValidateToken function from utils
	claims, err := utils.ValidateTokenAdmin(tokenString)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid token")
		return
	}

	// You can perform further actions based on the user's identity
	// For example, fetch user details from the database based on userID
	// user, err := c.authService.GetUserByID(claims.Uid)

	// Respond with user information (This is just for demonstration)
	utils.SendSuccessResponse(ctx, http.StatusOK, "Valid token", gin.H{"user_id": claims.Aid})
}