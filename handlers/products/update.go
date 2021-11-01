package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/terrytay/product-api/data"
)

// UpdateProducts takes a product and update the products list
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	// Obtain id (string) from URL param and convert to id (int)
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
		return
	}

	// Update product
	err = data.UpdateProduct(id, &prod)
	if errors.Is(err, data.ErrProductNotFound) {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
