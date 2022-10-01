package handlers

import (
	"net/http"
	"productservice/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	productsList := data.GetProducts()
	c.JSON(http.StatusOK, productsList)
}

func PostProduct(c *gin.Context) {

	p, _  := c.Get("payload")

	prod := p.(*data.Product)

	data.AddProduct(prod)
}

func PutProduct(c *gin.Context) {

	// validate route param
	stringId := c.Param("productId")

	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id is invalid"})
		return
	}

	//hacky way of using context and middleware for validation
	p, _  := c.Get("payload")
	prod := p.(*data.Product)

	// update product
	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
}

func DeleteProduct(c *gin.Context) {
	// validate route param is int
	stringId := c.Param("productId")

	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id is invalid"})
		return
	}

	err = data.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

}
