package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func main(){

	initRoutes()
	
	router.Run(":1337")
}

