package main

import (
	"github.com/dortlii/p9r-backend/controllers"
	"github.com/dortlii/p9r-backend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/namespaces", controllers.NamespaceList)
	r.POST("/namespaces", controllers.NamespaceCreate)

	r.Run("localhost:3000")
}
