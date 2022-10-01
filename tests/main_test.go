package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"productservice/handlers"
	"productservice/middlewares"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func MockRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPingHandler(t *testing.T) {

	mockResponse := `{"ping":"pong"}`

	r := MockRouter()

	r.GET("/ping", handlers.GetPing)

	req, _ := http.NewRequest("GET", "/ping", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestPostHandler(t *testing.T){

	mockRequest := strings.NewReader(`{"name":"frappe", "description":"le frappe", "price":1.69}`)

	r := MockRouter()

	r.POST("/products", middlewares.ProductValidator(), handlers.PostProduct)

	req, _ := http.NewRequest("POST", "/products", mockRequest)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestPutHandler(t *testing.T){

	mockRequest := strings.NewReader(`{"name":"McFrappe", "description":"le frappe", "price":1.69, "sku":"frappe123", "createdOn":"today", "updatedOn":"today", "DeletedOn":"today"}`)
	mockResponse := `{"msg":" product 1 successfully updated"}`
	
	r := MockRouter()

	r.PUT("/products/:productId", middlewares.ProductValidator(), handlers.PutProduct)

	req, _ := http.NewRequest("PUT", "/products/1", mockRequest)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

