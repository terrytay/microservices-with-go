package handlers

import (
	"net/http"

	"github.com/terrytay/product-api/data"
)

//	swagger:route GET /products products getProducts
//	Returns a list of products
//	responses:
//		200: productsResponse

// GetProducts return the list of products
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
		return
	}
}
