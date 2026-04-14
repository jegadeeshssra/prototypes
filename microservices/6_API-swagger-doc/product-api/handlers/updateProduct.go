package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

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
