package handlers

import (
	"net/http"

	"github.com/alwesleys/product-api/data"
)

// swagger:route POST /products products addProduct
// adds the product provided
// responses:
//		200: noContent

// POST add new product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST: Add Product")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}
