package main

import (
	"bytes"
	"encoding/json"
	"github.com/dortlii/p9r-backend/controllers"
	"github.com/dortlii/p9r-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

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
