package main

import (
	"log"
	"net/http"
	"os"
	"product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := http.Server{
		Addr:    ":9090",
		Handler: sm,
	}

	s.ListenAndServe()
}
