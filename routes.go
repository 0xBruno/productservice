package main

import (
	"productservice/handlers"

	"productservice/middlewares"
)

func initRoutes() {

	// Wire up routes
	router.GET("/", handlers.GetIndex)
	
	api := router.Group("/api")
	{ 
		api.GET("/ping", handlers.GetPing)
		api.GET("/products", handlers.GetProducts)
		api.POST("/products", middlewares.ProductValidator(), handlers.PostProduct)
		api.PUT("/products/:productId", middlewares.ProductValidator(), handlers.PutProduct)
		api.DELETE("/products/:productId", handlers.DeleteProduct)
	}
}
