package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := json.Marshal(data.GetProducts())

	if err != nil {
		http.Error(rw, "Error getting products", http.StatusInternalServerError)
	}

	rw.Write(d)
}
