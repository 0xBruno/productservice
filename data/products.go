package data

import (
	"encoding/json"
	"io"
	"time"
	"fmt"
)

type Product struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
	SKU         string `json:"sku"`
	CreatedOn   string `json:"--"`
	UpdatedOn   string `json:"--"`
	DeletedOn   string `json:"--"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

var productList = []*Product{
	&Product{
		ID:	1,
		Name: "Latte",
		Description:"Frothy milky coffee",
		Price: 2.45,
		SKU: "Latte1",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:	2,
		Name: "Espresso",
		Description:"short and strong coffee without milk",
		Price: 3.69,
		SKU: "Esspresso1",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

func UpdateProduct(id int, p *Product) error {

	pos, err := findProduct(id)

	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil 
}

var ErrProductNotFound = fmt.Errorf("Product not found! ")

func findProduct(id int) (pos int, e error) {
	for i, p := range productList {

		if p.ID == id {
			return i, nil 
		}
	}

	return -1, ErrProductNotFound
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product){
	p.ID = getNextID()
	productList = append(productList, p)
}

func DeleteProduct(id int) error { 
	newProductList := []*Product{}

	var productFound bool
	// Iterate over productList, if ID does not match 
	// add to newProductList. After looping set productList 
	// to newProductList productFound flag set to return error 
	// if not found
	for _, prod := range productList {
		if prod.ID == id {
			productFound = true

		} else {
			newProductList = append(newProductList, prod)
		}
	}

	if !productFound { 
		return fmt.Errorf("product not found")
	}

	productList = newProductList

	return nil 
}

func getNextID() int  {
	lp := productList[len(productList) - 1]
	return lp.ID + 1 
}