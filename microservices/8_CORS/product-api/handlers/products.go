package handlers

import (
	"log"
)

type Products struct {
	l *log.Logger
}

// Create a new instance of Products which is a handler(contains ServerHTTP)
func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}
