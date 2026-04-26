package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
//	201: noContentResponse

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(res http.ResponseWriter, req *http.Request) {

	// retrieves the id value from the mux.Vars(req)
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(res, "Product not found", http.StatusInternalServerError)
		return
	}
}
