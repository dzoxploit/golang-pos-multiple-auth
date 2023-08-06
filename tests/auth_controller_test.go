package test

import (
	"bytes"
	"encoding/json"
	"gocommerce/controllers"
	"gocommerce/services"
	"gocommerce/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestAuthController_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()

	mockAuthService := &services.AuthService{}
    authController := controllers.NewAuthControllerTest(mockAuthService)

    // Define the route
    router.POST("/login", authController.Login)

	// Define the test case
	t.Run("Login with valid credentials", func(t *testing.T) {
		// Create a request
		reqBody := map[string]string{
			"username": "testuser",
			"password": "testpassword",
		}
		reqJSON, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Check the response
		assert.Equal(t, http.StatusOK, w.Code)

		var response utils.SuccessResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, "Login successful", response.Message)
		// Additional checks for the response data if needed
	})
	// Add more test cases for different scenarios
}
