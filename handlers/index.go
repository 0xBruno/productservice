package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Index is a simple handler
type Index struct {
	l *log.Logger
}

// NewIndex creates a new index handler with the given logger
func NewIndex(l *log.Logger) *Index {
	return &Index{l}
}

// ServeHTTP implements the go http.Handler Interface
// https://pkg.go.dev/net/http#Handler
func (i* Index) ServeHTTP(resp http.ResponseWriter, req *http.Request){
	
	// Read the body 
	d, err := ioutil.ReadAll(req.Body)
	sData := string(d)

	if err != nil {
		http.Error(resp, "Oops", http.StatusBadRequest)
		log.Fatalln(err)
		return 
	}
	// Write the response 
	resp.Write(d)
	i.l.Println(sData)

}