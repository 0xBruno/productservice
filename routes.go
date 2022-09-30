package main

import "productservice/handlers"

func initRoutes() {

	// Wire up routes
	router.GET("/", handlers.GetIndex)
	router.GET("/ping", handlers.GetPing)
	router.GET("/products", handlers.GetProducts)
	router.POST("/products", handlers.PostProduct)
	router.PUT("/products/:productId", handlers.PutProduct)
	router.DELETE("/products/:productId", handlers.DeleteProduct)
}
