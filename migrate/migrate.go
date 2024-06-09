package main

import (
	"go-gin-gorm-crud/initializers"
	"go-gin-gorm-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
