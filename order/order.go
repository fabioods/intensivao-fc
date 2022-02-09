package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"order/db"
	"order/queue"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
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

func main() {
	var param string
	flag.StringVar(&param, "opt", "", "")
	flag.Parse()

	in := make(chan []byte)
	connection := queue.Connect()

	switch param {
	case "checkout":
		queue.StartConsuming("checkout_queue", connection, in)
		for payload := range in {
			notifyOrderCreated(CreateOrder(payload), connection)
			fmt.Println("Recibiendo payload checkout: ", string(payload))
		}
	case "payment":
		queue.StartConsuming("payment_queue", connection, in)
		var order Order
		for payload := range in {
			json.Unmarshal(payload, &order)
			saveOrder(order)
			fmt.Println("Recibiendo payload payment: ", string(payload))
		}
	}

}

func CreateOrder(payload []byte) Order {
	var order Order
	json.Unmarshal(payload, &order)
	order.UUID = uuid.New().String()
	order.Status = "pending"
	order.CreatedAt = time.Now()
	saveOrder(order)
	return order
}

func saveOrder(order Order) {
	json, _ := json.Marshal(order)
	connection := db.Connect()
	err := connection.Set(context.Background(), order.UUID, string(json), 0).Err()
	if err != nil {
		panic(err.Error())
	}
}

func notifyOrderCreated(order Order, ch *amqp.Channel) {
	json, _ := json.Marshal(order)
	queue.Notify(json, "order_ex", "", ch)

}
