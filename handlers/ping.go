package handlers

import (
	"net/http"
	"microservices/utils"
	
)

type Ping struct {

}

func NewPing() *Ping {
	return &Ping{}
}

func (p* Ping) ServeHTTP(resp http.ResponseWriter, req *http.Request){
	utils.Log("",req)
	resp.Write([]byte("OK\r\n"))
}