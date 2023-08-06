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

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		authService: services.NewAuthService(db),
	}
}

func NewAuthControllerTest(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := c.authService.Register(&user)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SendSuccessResponse(ctx, http.StatusCreated, "User created successfully", nil)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}


	// Authenticate the user
	jwtToken, err := c.authService.Login(user.Username, user.Password)
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

func (c *AuthController) ValidateToken(ctx *gin.Context) {
	// Get the token from the request header
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Authorization token missing")
		return
	}

	fmt.Println(tokenString)

	// Validate the token using the ValidateToken function from utils
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid token")
		return
	}

	// You can perform further actions based on the user's identity
	// For example, fetch user details from the database based on userID
	// user, err := c.authService.GetUserByID(claims.Uid)

	// Respond with user information (This is just for demonstration)
	utils.SendSuccessResponse(ctx, http.StatusOK, "Valid token", gin.H{"user_id": claims.Uid})
}