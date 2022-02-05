package main

import (
	"checkout/queue"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var productsURL string

type Product struct {
	UUID    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Order struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ProductID string `json:"product_id"`
}

func init() {
	productsURL = os.Getenv("PRODUCT_URL")
}

func displayCheckout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	response, err := http.Get(productsURL + "/product/" + id)
	if err != nil {
		fmt.Println("Error de http", err.Error())
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	var product Product
	json.Unmarshal(data, &product)
	t := template.Must(template.ParseFiles("templates/checkout.html"))
	t.Execute(w, product)
}

func finish(w http.ResponseWriter, r *http.Request) {
	var order Order
	order.Name = r.FormValue("name")
	order.Email = r.FormValue("email")
	order.Phone = r.FormValue("phone")
	order.ProductID = r.FormValue("product_id")

	data, _ := json.Marshal(order)

	connection := queue.Connect()
	queue.Notify(data, "checkout_ex", "", connection)

	fmt.Println(string(data))
	w.Write([]byte(("Gracias por tu compra")))
}

func main() {
	r := mux.NewRouter()
	//preceder finish antes de id, pois id espera um parametro, e sendo colocado antes, entenderia que /finish é um parametro
	r.HandleFunc("/finish", finish)
	r.HandleFunc("/{id}", displayCheckout)

	http.ListenAndServe(":8082", r)
}
