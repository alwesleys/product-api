// Package classification Product API
//
// Documentation for Product API
//
//	Schemes: http
// 	BasePath: /
// 	Version: 1.0.0
//	Host: github.com/alwesleys
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	-application/json
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alwesleys/product-api/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProduct(l)

	// create new serve mux
	// bind handlers
	sm := mux.NewRouter()

	const root = "/products"

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(root, ph.GetProducts)
	getRouter.HandleFunc(root+"/{id:[0-9]+}", ph.GetProductById)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(root+"/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc(root, ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	delRouter := sm.Methods(http.MethodDelete).Subrouter()
	delRouter.HandleFunc(root+"/{id:[0-9]+}", ph.DeleteProductByID)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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
