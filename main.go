package main

import (
	"go-gin-gorm-crud/controllers"
	"go-gin-gorm-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()
	r.POST("/createPost", controllers.PostsCreate)
	r.GET("/posts", controllers.FetchPosts)
	r.Run()
}
