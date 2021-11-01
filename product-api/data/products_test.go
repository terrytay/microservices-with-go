package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terrytay/microservices-with-go/product-api/utils"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1.22,
		SKU:   "abc-def-ghi",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		Name: "abc",
		SKU:  "abc-def-ghi",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestValidProductDoesNOTReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc-efg-hji",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 0)
}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		&Product{
			Name: "abc",
		},
	}

	b := bytes.NewBufferString("")
	err := utils.ToJSON(ps, b)
	assert.NoError(t, err)
}
