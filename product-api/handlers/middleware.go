package handlers

import (
	"context"
	"net/http"

	"github.com/terrytay/microservices-with-go/product-api/data"
	"github.com/terrytay/microservices-with-go/product-api/utils"
)

// MiddlwwareValidateProduct is middlware that validates the request body object
// with struct from data.Product{}
// Propogates in context upon successful validation with key KeyProduct{}
func (p *Products) MiddlwareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		prod := &data.Product{}

		err := utils.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			utils.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		errs := p.v.Validate(prod)
		if errs != nil {
			p.l.Println("[ERROR] validating product", err)

			w.WriteHeader(http.StatusUnprocessableEntity)
			utils.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
