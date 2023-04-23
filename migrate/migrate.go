package main

import (
	"github.com/TaronVardanyan/Go-CRUD/initializers"
	"github.com/TaronVardanyan/Go-CRUD/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}