package middlewares

import (
	"net/http"
	"productservice/data"

	"github.com/gin-gonic/gin"
)

func ProductValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate req body Product obj
		prod := &data.Product{}

		err := prod.FromJSON(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unable to unmarshal json"})
			c.Abort()
			return
		} else {
			c.Set("payload", prod)
			c.Next()
		}


	}
}
