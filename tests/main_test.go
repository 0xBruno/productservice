package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"productservice/handlers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPingHandler(t *testing.T) {

	mockResponse := `{"ping":"pong"}`

	r := SetUpRouter()

	r.GET("/ping", handlers.GetPing)

	req, _ := http.NewRequest("GET", "/ping", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}
