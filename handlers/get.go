package handlers

import (
	"net/http"
	"strconv"

	"github.com/alwesleys/product-api/data"
	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// returns a list of products
// responses:
//		200: productsResponseWrapper

// GET all products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Products")
	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Error getting products", http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /products/{id} products getProduct
// returns a specific product
// responses:
//		200: productResponseWrapper

// GET specific product
func (p *Products) GetProductById(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Product by ID")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Invalid product ID", http.StatusBadRequest)
		return
	}

	lp, err := data.GetProductByID(id)

	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}

	err = lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Error getting product", http.StatusInternalServerError)
		return
	}
}
