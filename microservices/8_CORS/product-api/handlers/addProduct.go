package handlers

import (
	"net/http"
	"product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  X - 422: errorValidation
//  X - 501: errorResponse

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	// getting the deserialized from json to the data.Product{}
	// temp - interface which contains the Product type values
	temp := req.Context().Value(KeyProduct{})
	// This will assert the interface that the type of value it contains is Product{} type
	product := temp.(*data.Product)
	data.AddProduct(product)
}
