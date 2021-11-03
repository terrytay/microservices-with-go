package handlers

import (
	"errors"
	"net/http"

	"github.com/terrytay/microservices-with-go/product-api/data"
	"github.com/terrytay/microservices-with-go/product-api/utils"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
// responses:
//	204: noContentResponse
//  404: errorResponse
//  500: errorResponse

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := p.getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	// Update product
	err := data.DeleteProduct(id)
	if errors.Is(err, data.ErrProductNotFound) {
		p.l.Println("[ERROR] deleting record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}
	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
