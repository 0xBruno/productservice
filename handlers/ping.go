package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context){
	
	c.JSON(http.StatusOK, gin.H{"ping":"pong"})

}