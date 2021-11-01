package handlers

import (
	"errors"
	"net/http"

	"github.com/terrytay/microservices-with-go/product-api/data"
	"github.com/terrytay/microservices-with-go/product-api/utils"
)

// swagger:route PUT /products products updateProduct
// Update a products details
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// UpdateProducts takes a product and update the products list
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if errors.Is(err, data.ErrProductNotFound) {
		w.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: "Product not found in database"}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
