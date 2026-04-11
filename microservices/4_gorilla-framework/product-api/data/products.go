package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// When this method is called with an io.reader, it will automatically fill the necessary values from reqBody
func (p *Product) FromJSON(reqBody io.Reader) error {
	d := json.NewDecoder(reqBody)
	return d.Decode(p) // Reads the JSON data from the request body. Looks at the JSON keys in the incoming data. Matches them with the struct fields using the json:"..." tags. Automatically fills (populates) the matching fields in your Product struct. Returns an error if the JSON is invalid or types don't match.
}

// productList is slice(dynamic) of type ptr(Product struct type) containing ptrs to individual product
// example data source
var productList = []*Product{ // [] - dynamic Arr , *Product - this slice stores ptr to product struct type
	&Product{ // Create a Product struct with these values, then give me a pointer to it.
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func generateID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = generateID()
	productList = append(productList, p)
}

// Custom error
var ErrProductNotFound = fmt.Errorf("Product Not Found")

func UpdateProduct(id int, p *Product) error {
	i, _, err := FindProduct((id))
	if err != nil {
		return err
	}
	p.ID = id // replace the default value of 0
	productList[i] = p

	return nil
}

func FindProduct(id int) (int, *Product, error) {
	for index, prod := range productList {
		if prod.ID == id {
			return index, prod, nil
		}
	}
	return -1, nil, ErrProductNotFound
}

// Created a separate Class with methods for this type []*Products
// every listOfProduct instance will contain this method ToJSON
// whoever have this type will have this method
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
