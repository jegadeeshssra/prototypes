package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

// Create a new instance of Products which is a handler(contains ServerHTTP)
func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to parse json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	// getting the deserialized from json to the data.Product{}
	// temp - interface which contains the Product type values
	temp := req.Context().Value(KeyProduct{})
	// This will assert the interface that the type of value it contains is Product{} type
	product := temp.(*data.Product)
	data.AddProduct(product)
}

func (p *Products) UpdateProduct(res http.ResponseWriter, req *http.Request) {

	// retrieves the id value from the mux.Vars(req)
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Unable to convert id", http.StatusBadRequest)
		return
	}

	// getting the deserialized from json to the data.Product{}
	// temp - interface which contains the Product type values
	temp := req.Context().Value(KeyProduct{})
	// This will assert the interface that the type of value it contains is Product{} type
	product := temp.(*data.Product)

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(res, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

// returns an http.Handler which satisfies the http.Handler interface
// Internally , it will return an object of http.Handlerfunc and then object.ServeHTTP will execute our anonymous function.
func (p *Products) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		singleProduct := &data.Product{}
		err := singleProduct.FromJSON(req.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(res, "Error reading the product", http.StatusBadRequest)
		}

		err = singleProduct.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(res, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(req.Context(), KeyProduct{}, singleProduct)
		req = req.WithContext(ctx)

		// This will call the UpdateProduct
		next.ServeHTTP(res, req)

	})
}
