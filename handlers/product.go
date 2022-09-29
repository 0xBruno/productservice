package handlers

import (
	"log"
	"microservices/data"
	"microservices/utils"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		utils.Log("", req)
		p.getProducts(resp, req)
		return

	} else if req.Method == http.MethodPost {
		utils.Log("", req)
		p.addProduct(resp, req)
		return
	} else if req.Method == http.MethodPut {

		utils.Log("", req)

		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(req.URL.Path, -1)

		if len(g) != 1 {
			utils.Log("Invalid URI more than one id", req)
			http.Error(resp, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			utils.Log("Invalid URI more than one capture group", req)
			http.Error(resp, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			utils.Log("Invalid URI unable to convert to numer", req)
			http.Error(resp, "Invalid URI", http.StatusBadRequest)
			return
		}
		
		p.updateProduct(id, resp, req)
		return 
		
	}
	
	resp.WriteHeader(http.StatusMethodNotAllowed)
}


func (p *Products) updateProduct(id int, resp http.ResponseWriter, req *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(req.Body)
	

	if err != nil {
		http.Error(resp, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(resp, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(resp, "Error", http.StatusInternalServerError)
		return
	}

}

func (p *Products) addProduct(resp http.ResponseWriter, req *http.Request) {

	prod := &data.Product{}

	err := prod.FromJSON(req.Body)

	if err != nil {
		http.Error(resp, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(prod)

}

func (p *Products) getProducts(resp http.ResponseWriter, req *http.Request) {
	productsList := data.GetProducts()
	productsList.ToJSON(resp)
}
