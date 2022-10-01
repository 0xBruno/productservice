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

	prod, _  := c.Get("payload")

	data.AddProduct(prod.(*data.Product))
}

func PutProduct(c *gin.Context) {

	// validate route param
	stringId := c.Param("productId")

	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id is invalid"})
		return
	}

	// validate req body Product obj
	prod := &data.Product{}

	err = prod.FromJSON(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to unmarshal json"})
		return
	}

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
