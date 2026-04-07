package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"regexp"
	"strconv"
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
	if req.Method == http.MethodPost {
		p.AddProduct(res, req)
		return
	}
	if req.Method == http.MethodPut {
		// expect the id in path
		p.l.Println("Invalid URI more than one id")
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(req.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.UpdateProduct(id, res, req)
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

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Bad Request", http.StatusBadRequest)
	}
	data.AddProduct(product)
}

func (p *Products) UpdateProduct(id int, res http.ResponseWriter, req *http.Request) {
	product := &data.Product{}
	err := product.FromJSON((req.Body))
	if err != nil {
		http.Error(res, "Bad Request", http.StatusBadRequest)
	}
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
