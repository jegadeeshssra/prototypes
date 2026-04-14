// Package classification of Product API
//
// Documentation for Product API
//
//		Schemes: http
//	 BasePath: /
//	 Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//	 Produces:
//	 - application/json
//
// swagger:meta
package handlers

import (
	"log"
	"product-api/data"
)

type Products struct {
	l *log.Logger
}

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// Create a new instance of Products which is a handler(contains ServerHTTP)
func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}
