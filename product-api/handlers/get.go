package handlers

import (
	"net/http"

	"github.com/terrytay/microservices-with-go/product-api/data"
	"github.com/terrytay/microservices-with-go/product-api/utils"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetProducts return the list of products
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()
	err := utils.ToJSON(prods, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// GetProduct return the product by ID
func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = utils.ToJSON(prod, w)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
