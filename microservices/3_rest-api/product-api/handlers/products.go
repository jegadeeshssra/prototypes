package handlers

import (
	"log"
	"net/http"
	"product-api/data"
)

// Create a new instance of Products which is a handler(contains ServerHTTP)
func NewProductsHandler(l *log.Logger) *Products {
	return &Products{l}
}

type Products struct {
	l *log.Logger
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	p.l.Printf("%v %v", req.URL.Path, req.Method)
	if req.Method == http.MethodGet {
		p.GetProducts(res, req)
		return
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to parse json", http.StatusInternalServerError)
	}
}
