package main

import (
	"github.com/dortlii/p9r-backend/initializers"
	"github.com/dortlii/p9r-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Namespace{})
}
