package utils

import (
	"log"
	"net/http"
)

func Log(msg string, req *http.Request){
	log.Printf("%s %s %s %s", msg, req.Method, req.URL, req.UserAgent())
}