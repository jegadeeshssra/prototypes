package handlers

import (
	"net/http"
	"product-api/data"
)

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	// getting the deserialized from json to the data.Product{}
	// temp - interface which contains the Product type values
	temp := req.Context().Value(KeyProduct{})
	// This will assert the interface that the type of value it contains is Product{} type
	product := temp.(*data.Product)
	data.AddProduct(product)
}
