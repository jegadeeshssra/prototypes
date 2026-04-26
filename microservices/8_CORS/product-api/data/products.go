package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for this user
	//
	// required: true
	// min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// When this method is called with an io.reader, it will automatically fill the necessary values from reqBody
func (p *Product) FromJSON(reqBody io.Reader) error {
	d := json.NewDecoder(reqBody)
	return d.Decode(p) // Reads the JSON data from the request body. Looks at the JSON keys in the incoming data. Matches them with the struct fields using the json:"..." tags. Automatically fills (populates) the matching fields in your Product struct. Returns an error if the JSON is invalid or types don't match.
}

func (p *Product) ToJson(id int, w io.Writer) error {
	e := json.NewEncoder((w))
	return e.Encode(p)
}

// Custom validation function for "sku" parameter
func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	fmt.Printf("Matches found: %v\n", matches)
	fmt.Printf("Number of matches: %d\n", len(matches))

	if len(matches) != 1 {
		return false
	}

	return true
}

// Validate method within the Product interface
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

// Created a separate Class with methods for this type []*Products
// every listOfProduct instance will contain this method ToJSON
// whoever have this type will have this method
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
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

func GetSingleProduct(id int) (*Product, error) {
	_, product, err := FindProduct((id))
	if err != nil {
		return nil, err
	}
	return product, nil
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

func DeleteProduct(id int) error {
	i, _, err := FindProduct((id))
	if err != nil {
		return err
	}
	// ... (variadic expansion) Expands slice into elements: append([A, B], D, E) -> [A, B, D, E]
	productList = append(productList[:i], productList[i+1:]...)

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
