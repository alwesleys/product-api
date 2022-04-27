package handlers

import (
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
	if r.Method == http.MethodGet { // GET
		p.getProducts(rw, r)
		return
	} else if r.Method == http.MethodPost { // POST
		p.addProduct(rw, r)
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// GET all products
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Error getting products", http.StatusInternalServerError)
	}
}

// POST add new product
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {

	var newP data.Product
	err := newP.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	}

	data.AddProduct(&newP)
}
