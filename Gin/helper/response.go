package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseJson(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func ResponseSuccess(c *gin.Context,message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

func ResponseError(c *gin.Context, code int, err interface{}) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
