package main

import (
	"testing"

	"github.com/terrytay/microservices-with-go/product-api/sdk/client/client"
	"github.com/terrytay/microservices-with-go/product-api/sdk/client/client/products"
	"github.com/terrytay/microservices-with-go/product-api/sdk/client/models"
)

func TestOurClientListProducts(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	_, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOurClientListSingleProduct(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListSingleProductParams().WithID(1)
	_, err := c.Products.ListSingleProduct(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOurClientCreateProduct(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	var name string = "TestName"
	var price float32 = 6.99
	var sku string = "zxc-bnm-jkl"

	params := products.NewCreateProductParams().WithBody(&models.Product{Name: &name, Price: &price, SKU: &sku})
	_, err := c.Products.CreateProduct(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOurClientUpdateProduct(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	var name string = "TestName"
	var price float32 = 6.99
	var sku string = "zxc-bnm-jkl"

	params := products.NewUpdateProductParams().WithBody(&models.Product{ID: 1, Name: &name, Price: &price, SKU: &sku})
	_, err := c.Products.UpdateProduct(params)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOurClientDeleteProduct(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewDeleteProductParams().WithID(1)
	_, err := c.Products.DeleteProduct(params)
	if err != nil {
		t.Fatal(err)
	}
}
