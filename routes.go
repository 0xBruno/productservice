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
		api.DELETE("/products/:productId", handlers.DeleteProduct)

		// Require data.Product
		api.POST("/products", middlewares.ValidateProductMiddleware(), handlers.PostProduct)
		api.PUT("/products/:productId", middlewares.ValidateProductMiddleware(),handlers.PutProduct)
	}
}
