package handlers

import (
	"net/http"

	"github.com/terrytay/product-api/data"
)

// AddProduct takes a product and append it to products list
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}
