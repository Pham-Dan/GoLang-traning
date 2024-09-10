package controller

import (
	"main/domain"
	"main/helper"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddNewUser(c *gin.Context) {
	var userRequest domain.CreateUserRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if userRequest.Name == "" || userRequest.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required fields"})
		return
	}
    hash, _ := HashPassword(userRequest.Password) // ignore error for the sake of simplicity
	userRequest.Password = hash
	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	_, err := models.DB.NewInsert().Model(&user).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetUserByID(c *gin.Context) {
	userID, valid := ValidateUserID(c)

	if !valid {
		return
	}

	var user models.User

	err := models.DB.NewSelect().
		Model(&user).Relation("Post").Where("id = ?", userID).Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No record found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func ValidateUserID(c *gin.Context) (string, bool) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be present"})
		return "", false
	}
	return userID, true
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func Profile(c *gin.Context) {
	var user models.User
	userId, ok := c.Get("userId")
	if !ok {
		helper.ResponseError(c, http.StatusNotFound, "")
		return
	}

	err := models.DB.NewSelect().
		Model(&user).Relation("Post").Where("id = ?", userId).Scan(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No record found"})
		return
	}
	c.JSON(http.StatusOK, user)
}