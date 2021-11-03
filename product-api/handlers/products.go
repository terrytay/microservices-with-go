package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/terrytay/microservices-with-go/product-api/data"
)

type Products struct {
	l        *log.Logger
	v        *data.Validation
	urlParam func(r *http.Request, key string) string
}

func NewProducts(l *log.Logger, v *data.Validation, u func(r *http.Request, key string) string) *Products {
	return &Products{l, v, u}
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
func (p Products) getProductID(r *http.Request) int {
	// parse the product id from the url
	idString := p.urlParam(r, "id")

	// convert the id into an integer and return
	id, err := strconv.Atoi(idString)
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
