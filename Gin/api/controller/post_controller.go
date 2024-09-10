package controller

import (
	"encoding/csv"
	"fmt"
	"log"
	"main/domain"
	"main/helper"
	"main/models"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	err := models.DB.NewSelect().Model(&posts).Order("id ASC").Scan(c.Request.Context())
	if err != nil {
		helper.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	host := c.Request.Host
	scheme := "http"

	for i, post := range posts {
		posts[i].Image = fmt.Sprintf("%s://%s", scheme, host) + "/public/" + post.Image
	}

	for _, v := range posts {
		log.Println(v.Image)
	}
	helper.ResponseSuccess(c, "", posts)
}

func AddNewPost(c *gin.Context) {
	var newPostRequest domain.CreatePostRequest

	if err := c.ShouldBind(&newPostRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": newPostRequest})
		return
	}
	errorsRe := map[string]any{}
	if err := domain.Validate.Struct(&newPostRequest); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorsRe[err.Field()] = err.ActualTag()
		}

		helper.ResponseError(c, http.StatusBadRequest, errorsRe)
		return

	}
	file, _ := c.FormFile("image")
	var filename string
	if file != nil {
		filename = filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, "public/"+filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	post := models.Post{
		Title:   newPostRequest.Title,
		Content: newPostRequest.Content,
		Image:   filename,
		UserID:  1,
	}
	_, err := models.DB.NewInsert().Model(&post).Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new post"})
		return
	}
	helper.ResponseSuccess(c, "New Post Created", post)
}

func UpdatePost(c *gin.Context) {
	var postRequest domain.UpdatePostRequest

	if err := c.ShouldBind(&postRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": postRequest})
		return
	}

	if err := domain.Validate.Struct(&postRequest); err != nil {
		helper.ResponseError(c, http.StatusBadRequest, err.Error())
		return

	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	post := models.Post{
		ID:      uint64(id),
		Title:   postRequest.Title,
		Content: postRequest.Content,
		UserID:  1,
	}
	_, err := models.DB.NewUpdate().Model(&post).WherePK().Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new post"})
		return
	}
	helper.ResponseSuccess(c, "New Post Created", post)
}

func GetPostByID(c *gin.Context) {
	postID, valid := ValidatePostID(c)

	if !valid {
		return
	}

	var post models.Post
	err := models.DB.NewSelect().Model(&post).Where("id = ?", postID).Scan(c.Request.Context())

	if err != nil {
		helper.ResponseError(c, http.StatusInternalServerError, err.Error())

		return
	}
	helper.ResponseSuccess(c, "success", post)

}

func DeletePostByID(c *gin.Context) {

	postID, valid := ValidatePostID(c)
	if !valid {
		return
	}

	var post models.Post
	_, err := models.DB.NewDelete().Model(&post).Where("id = ?", postID).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post successfully deleted!"})
}

func ValidatePostID(c *gin.Context) (string, bool) {
	postID := c.Param("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be present"})
		return "", false
	}
	return postID, true
}
func ExportPostCsv(c *gin.Context) {
	var posts []models.Post
	err := models.DB.NewSelect().Model(&posts).Order("id ASC").Scan(c.Request.Context())
	if err != nil {
		helper.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Create a buffer to write the CSV data
	var sb strings.Builder
	writer := csv.NewWriter(&sb)

	// Write header
	err = writer.Write([]string{"ID", "Title", "Content"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write data rows
	for _, post := range posts {
		err = writer.Write([]string{
			strconv.FormatInt(int64(post.ID), 10),
			post.Title,
			post.Content,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// Flush writer and set response headers
	writer.Flush()
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=users.csv")
	c.Data(http.StatusOK, "text/csv;charset=UTF-8", []byte(sb.String()))

}
