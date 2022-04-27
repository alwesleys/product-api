package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alwesleys/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProduct(l)

	// create new serve mux
	// bind handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create new server
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// make channel for GoRoutine of the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// start server in a GoRoutine
	go func() {
		if err := s.ListenAndServe(); err != nil {
			l.Fatal(err)
		}
	}()

	l.Println("Server start")

	<-sigChan // receives the signal

	l.Println("Server stopped")

	c, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(c)
}
