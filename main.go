package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"microservices/handlers"
)


func main(){

	// Dependencies 
	l := log.New(os.Stdout, "LOG:", log.LstdFlags)

	// Create handlers
	hIndex		:= handlers.NewIndex(l)
	hPing 		:= handlers.NewPing()
	hGoodbye 	:= handlers.NewGoodbye(l)
	hProducts	:= handlers.NewProducts(l)

	// Create a new serve mux & register handlers 
	sm := http.NewServeMux()
	sm.Handle("/", hIndex)
	sm.Handle("/ping", hPing)
	sm.Handle("/goodbye", hGoodbye)
	sm.Handle("/products/", hProducts)

	
	// Configure server 
	s := &http.Server{
		Addr: ":1337",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	
	// Start the server
	go func(){
		fmt.Printf("[*] HTTP Server started at %s\n", s.Addr)
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()
	
	// Handle shutdown signals gracefully
	sigChan := make(chan os.Signal, 100)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	
	sig := <- sigChan
	fmt.Println("\r\nReceived terminate, graceful shutdown", sig)

	// Ensure request completes within 30 seconds
	timeCtx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	s.Shutdown(timeCtx)
}

