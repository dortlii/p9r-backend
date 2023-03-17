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
