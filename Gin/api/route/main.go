package route

import (
	"main/api/controller"
	"main/api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.StaticFS("/public", http.Dir("public"))
	api := router.Group("/api/v1")

	api.POST("/login", controller.Login)

	auth := api.Use(middleware.AuthMiddleware())
	auth.POST("/users", controller.AddNewUser)
	//users
	auth.GET("/users/:id", controller.GetUserByID)
	auth.GET("/profile", controller.Profile)

	// posts
	auth.GET("/posts/export-csv", controller.ExportPostCsv)
	auth.GET("/posts", controller.GetPosts)
	auth.POST("/posts", controller.AddNewPost)
	auth.GET("/posts/:id", controller.GetPostByID)
	auth.PUT("/posts/:id", controller.UpdatePost)
	auth.DELETE("/posts/:id", controller.DeletePostByID)

	
	return router
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}