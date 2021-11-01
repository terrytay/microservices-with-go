package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/terrytay/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

// MiddlwwareValidateProduct is middlware that validates the request body object
// with struct from data.Product{}
// Propogates in context upon successful validation with key KeyProduct{}
func (p *Products) MiddlwareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
