package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/terrytay/product-api/data"
)

//	swagger:route DELETE /products/{id} products deleteProduct
//	Deletes a product by its ID
//	responses:
//		201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Product")

	// Obtain id (string) from URL param and convert to id (int)
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
		return
	}

	// Update product
	err = data.DeleteProduct(id)
	if errors.Is(err, data.ErrProductNotFound) {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
