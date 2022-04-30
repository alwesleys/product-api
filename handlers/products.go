package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/alwesleys/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct {
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			return
		}

		// validate product
		if err := prod.Validate(); err != nil {
			http.Error(
				rw,
				fmt.Sprintf("Error reading product: %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}

// List of wrappers for swagger documentation

// A list of products returned in a response
// swagger:response productsResponseWrapper
type productsResponseWrapper struct {
	// All products in the system
	// in:body
	Body []data.Product
}

// A specific product returned in a response
// swagger:response productResponseWrapper
type productResponseWrapper struct {
	// Specified product by client
	// in:body
	Body data.Product
}

// swagger:response noContent
type productsNoContentWrapper struct {
}

// swagger:parameters deleteProduct
// swagger:parameters getProduct
type productIDParamWrapper struct {
	// The id of the product to get/update/delete from DB
	// in: path
	// required: true
	ID int `json:"id"`
}
