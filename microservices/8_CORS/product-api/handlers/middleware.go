package handlers

import (
	"context"
	"fmt"
	"net/http"
	"product-api/data"
)

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
