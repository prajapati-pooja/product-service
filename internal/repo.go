package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ProductType string  `json:"type"`
}

type Repository interface {
	Get() []Product
}

type productRepository struct {
}

func NewProductRepository() Repository {
	return productRepository{}
}

func (ps productRepository) Get() []Product {
	bytes, e := ioutil.ReadFile("internal/products.json")
	if e != nil {
		log.Printf("not able to read products file %s", e.Error())
	}
	var products []Product
	e = json.Unmarshal(bytes,&products)
	if e != nil {
		log.Printf("not able to unmarshel products, %s", e.Error())
	}
	return products
}

func SayHello() string {
	return "Hello"
}
