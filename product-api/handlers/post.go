package handlers

import (
	"net/http"

	"github.com/terrytay/microservices-with-go/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// AddProduct takes a product and append it to products list
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	// fetch product from context
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	w.Header().Add("Content-Type", "application/json")
	data.AddProduct(prod)
}
