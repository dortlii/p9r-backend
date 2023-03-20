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

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/namespaces", controllers.NamespaceList)
	router.GET("/namespaces/:id", controllers.NamespaceById)

	router.POST("/namespaces/create", controllers.NamespaceCreate)

	router.DELETE("/namespaces/delete/:id", controllers.NamespaceDelete)

	return router
}

func main() {
	r := setupRouter()

	r.Run("localhost:3000")
}
