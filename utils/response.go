package utils

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

func SendErrorResponse(ctx *gin.Context, statusCode int, message string) {
	response := ErrorResponse{
		Message: message,
	}
	ctx.JSON(statusCode, response)
}

func SendSuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		Message: message,
		Data:    data,
	}
	ctx.JSON(statusCode, response)
}