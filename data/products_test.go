package data

import "testing"

func TestProductValidation(t *testing.T) {

	t.Run("Product keys not valid", func(t *testing.T) {
		p := Product{
			Name:  "Milo",
			Price: 5,
			SKU:   "dv-dd-ee",
		}
		err := p.Validate()

		if err != nil {
			t.Fatal(err)
		}
	})
}
