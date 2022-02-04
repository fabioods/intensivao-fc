package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Product struct {
	UUID    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product `json:"products"`
}

func loadData() []byte {
	jsonFile, err := os.Open("product.json")
	if err != nil {
		fmt.Printf("Error reading file: %s", err.Error())
		return nil
	}
	defer jsonFile.Close()
	data, _ := ioutil.ReadAll(jsonFile)
	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write([]byte(products))
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data := loadData()

	var products Products
	json.Unmarshal(data, &products)
	for _, product := range products.Products {
		if product.UUID == id {
			productFound, _ := json.Marshal(product)
			w.Write(productFound)
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", ListProducts)
	r.HandleFunc("/product/{id}", GetProductByID)
	http.ListenAndServe(":8081", r)
}
