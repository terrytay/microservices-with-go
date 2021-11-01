// Package classfication of Product API

// Documentation for Product API
//
// 	Schemes: http
// 	BasePath: /
//	Version: 1.0.0
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
// swagger:meta
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/terrytay/product-api/data"
)

type Products struct {
	l *log.Logger
	v *data.Validation
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

type KeyProduct struct{}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	idString := chi.URLParam(r, "id")

	// convert the id into an integer and return
	id, err := strconv.Atoi(idString)
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
