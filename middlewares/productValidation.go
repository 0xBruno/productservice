package middlewares

import (
	"net/http"
	"productservice/data"

	"github.com/gin-gonic/gin"
)

func ValidateProductMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product data.Product
		
		if err := c.ShouldBindJSON(&product); err == nil {
			c.Set("product", product)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
}
