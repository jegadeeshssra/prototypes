package handlers

import (
	"net/http"
	"product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productResponse

func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to parse json", http.StatusInternalServerError)
	}
}
