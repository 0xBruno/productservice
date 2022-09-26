package data

import (
	"encoding/json"
	"io"
	"time"
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

func GetProducts() Products {
	return productList
}


