package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts()
	res.Header().Add("Content-Type", "application/json")
	err := lp.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to parse json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	X - 404: errorResponse

func (p *Products) GetSingleProduct(res http.ResponseWriter, req *http.Request) {
	// retrieves the id value from the mux.Vars(req)
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod, err := data.GetSingleProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(res, "Product not found", http.StatusInternalServerError)
		return
	}

	err = prod.ToJson(id, res)
	if err != nil {
		http.Error(res, "Unable to parse json", http.StatusInternalServerError)
	}
}
