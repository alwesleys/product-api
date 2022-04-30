package handlers

import (
	"net/http"

	"github.com/alwesleys/product-api/data"
)

// POST add new product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST: Add Product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}
