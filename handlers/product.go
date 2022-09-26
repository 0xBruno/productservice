package handlers

import (
	"log"
	"net/http"

	"github.com/0xBruno/go_microservices/data"
)

type Products struct{ 
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(resp http.ResponseWriter, req *http.Request){
	
	if req.Method == http.MethodGet {
		p.getProducts(resp, req)
		return 
		
	} else if req.Method == http.MethodPut { 
		return 
	}

	resp.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts( resp http.ResponseWriter, req *http.Request){ 
	productsList := data.GetProducts()
	productsList.ToJSON(resp)
}