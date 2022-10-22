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
	
	product := c.MustGet("product").(data.Product)
	data.AddProduct(&product)

}

func PutProduct(c *gin.Context) {

	// validate route param
	stringId := c.Param("productId")

	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id is invalid"})
		return
	}

	product := c.MustGet("product").(data.Product)

	// update product
	err = data.UpdateProduct(id, &product)

	if err == data.ErrProductNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg":" product " + strconv.Itoa(id) + " successfully updated"})
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
