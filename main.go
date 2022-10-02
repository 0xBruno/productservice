package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func main(){

	initRoutes()
	
	//gin.SetMode(gin.ReleaseMode)
	
	router.Run(":1337")
}

