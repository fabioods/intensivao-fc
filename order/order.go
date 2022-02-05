package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"order/db"
	"order/queue"
	"os"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	UUID    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Order struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ProductID string    `json:"product_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

var productsURL string

func init() {
	productsURL = os.Getenv("PRODUCT_URL")
}

func main() {
	in := make(chan []byte)
	connection := queue.Connect()
	queue.StartConsuming(connection, in)

	for payload := range in {
		CreateOrder(payload)
		fmt.Println("Recibiendo payload: ", string(payload))
	}
}

func CreateOrder(payload []byte) {
	var order Order
	json.Unmarshal(payload, &order)
	order.UUID = uuid.New().String()
	order.Status = "pending"
	order.CreatedAt = time.Now()
	saveOrder(order)
}

func saveOrder(order Order) {
	json, _ := json.Marshal(order)
	connection := db.Connect()
	err := connection.Set(context.Background(), order.UUID, string(json), 0).Err()
	if err != nil {
		panic(err.Error())
	}
}

func getProductByID(productID string) Product {
	response, err := http.Get(productsURL + "/product/" + productID)
	if err != nil {
		fmt.Println("Error de http", err.Error())
	}
	data, _ := ioutil.ReadAll(response.Body)
	var product Product
	json.Unmarshal(data, &product)
	return product
}
