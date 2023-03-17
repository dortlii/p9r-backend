package controllers

import (
	"net/http"

	"github.com/dortlii/p9r-backend/initializers"
	"github.com/dortlii/p9r-backend/models"
	"github.com/gin-gonic/gin"
)

func NamespaceCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Name string `json:"name"`
	}

	c.ShouldBindJSON(&body)

	// create the namespace
	namespace := models.Namespace{Name: body.Name}

	result := initializers.DB.Create(&namespace) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed to create kubernetes namespace",
			"error":  result.Error,
		})
		return
	}

	// return it
	c.JSON(http.StatusOK, gin.H{
		"name": namespace,
	})
}

func NamespaceList(c *gin.Context) {
	// Get all namespaces
	var namespaces []models.Namespace
	initializers.DB.Find(&namespaces)

	// response
	c.JSON(200, gin.H{
		"namespaces": namespaces,
	})
}

func NamespaceById(c *gin.Context) {
	// Get ID off url
	id := c.Param("id")

	// Get the post
	var namespace models.Namespace
	initializers.DB.First(&namespace, id)

	// repond
	c.JSON(http.StatusOK, gin.H{
		"namespace": namespace,
	})
}

func NamespaceDelete(c *gin.Context) {
	// Get ID off url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Namespace{}, id)

	// respond
	c.JSON(http.StatusOK, gin.H{
		"status": "deleted namespace",
	})
}
