package controller

import (
	"main/domain"
	"main/domain/auth"
	"main/helper"
	"main/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loginRequest domain.LoginRequest

	if err := c.ShouldBind(&loginRequest); err != nil {
		helper.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := domain.Validate.Struct(&loginRequest)
	if err != nil {
		helper.ResponseError(c, http.StatusBadRequest, err.Error())
		return

	}
	var user models.User

	if err := models.DB.NewSelect().Model(&user).Where("email = ?", loginRequest.Email).Scan(c.Request.Context()); err != nil {
		helper.ResponseError(c, http.StatusNotFound, "User not found with the given email")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil {
		helper.ResponseError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	expirationTime := 5
	jwtSecret := os.Getenv("jwtSecret")
	jwtToken, err := auth.GenerateJwtToken(1, jwtSecret, expirationTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	res := domain.LoginResponse{
		AccessToken: jwtToken,
	}

	helper.ResponseJson(c, res)

}
