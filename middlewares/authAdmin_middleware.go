package middlewares

import (
	"fmt"
	"gocommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        clientToken := c.Request.Header.Get("Authorization")
        if clientToken == "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Token Not Found"})
            c.Abort()
            return
        }

		fmt.Println(clientToken)
 
        claims, err := utils.ValidateTokenAdmin(clientToken)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            c.Abort()
            return
        }
        
        c.Set("uid", claims.Aid)
        c.Next()
    }
}