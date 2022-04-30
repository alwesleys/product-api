package handlers

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/alwesleys/product-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// get ID from URI
func getIDFromURI(path string) (int, error) {
	regex := regexp.MustCompile(`/([0-9]+)`)
	g := regex.FindAllStringSubmatch(path, -1)

	if len(g) != 1 {
		return -1, errors.New("invalid uri")
	}

	//g[0] = [/id id]
	//g[0][1] = id
	idString := g[0][1]

	id, err := strconv.Atoi(idString)

	if err != nil {
		return -1, errors.New("invalid uri")
	}

	return id, nil
}

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

// GET specific product
func (p *Products) getProductById(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Product by ID")
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

// POST add new product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST: Add Product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)
}

// PUT update product
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("PUT: Update Product ", id)
	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
