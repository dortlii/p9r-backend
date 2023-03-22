package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dortlii/p9r-backend/controllers"
	"github.com/dortlii/p9r-backend/models"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestNamespaceCreate(t *testing.T) {
	router := setupRouter()
	router.POST("/namespaces/create", controllers.NamespaceCreate)
	namespace := models.Namespace{
		Name: "testing-namespace",
	}

	jsonValue, _ := json.Marshal(namespace)
	req, _ := http.NewRequest("POST", "/namespaces/create", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNamespaceList(t *testing.T) {
	router := setupRouter()
	router.GET("/namespaces", controllers.NamespaceList)
	req, _ := http.NewRequest("GET", "/namespaces", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var namespaces []models.Namespace
	json.Unmarshal(w.Body.Bytes(), &namespaces)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, namespaces)
}
