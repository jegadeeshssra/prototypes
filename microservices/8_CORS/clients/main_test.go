package main

import (
	"fmt"
	"testing"

	"product-api-client/client"
	"product-api-client/client/products"
)

// Defines a test function that will be executed by the Go testing framework
func TestOurClient(t *testing.T) {

	// Creates a default client configuration and sets the API server host to localhost:9090
	cfg := client.DefaultTransportConfig().WithHost("localhost:8090")
	// Initializes a new Swagger-generated HTTP client using the provided configuration
	c := client.NewHTTPClientWithConfig(nil, cfg)

	// Constructs request parameters for the ListProducts API endpoint
	params := products.NewListProductsParams()
	// Invokes the ListProducts API call and captures the response and any error
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	t.Log("Success: received products")
}
