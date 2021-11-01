package handlers

import (
	"net/http"

	"github.com/terrytay/product-api/data"
)

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
