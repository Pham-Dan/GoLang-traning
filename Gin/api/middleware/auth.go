package middleware

import (
	"main/domain/auth"
	"main/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        jwtSecret := os.Getenv("jwtSecret")
        tokenString := c.GetHeader("Authorization")
        userId, err := auth.GetUserIdFromJwtToken(tokenString, jwtSecret)
        if err != nil {
            helper.ResponseError(c,http.StatusUnauthorized, "Unauthorized")
            c.Abort()
        }
        c.Set("userId", userId)
        c.Next()
    }
}