package handlers

import (
	"github.com/gin-gonic/gin"
)


func GetIndex(c *gin.Context){
	
	c.Writer.Write([]byte("🙏\r\n"))

}