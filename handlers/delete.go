package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alwesleys/product-api/data"
	"github.com/gorilla/mux"
)

func (p *Products) DeleteProductByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to identify product", http.StatusBadRequest)
		return
	}

	p.l.Printf("DELETE Product %d", id)
	if err := data.DeleteProductByID(id); err != nil {
		http.Error(rw, fmt.Sprintf("Unable to delete product %d", id), http.StatusBadGateway)
		return
	}
}
