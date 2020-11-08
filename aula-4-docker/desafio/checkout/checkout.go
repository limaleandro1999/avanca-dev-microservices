package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/wesleywillians/go-rabbitmq/queue"
)

type Order struct {
	Coupon   string
	CcNumber string
}

type Result struct {
	Status string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env")
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":9090", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, Result{})
}

func process(w http.ResponseWriter, r *http.Request) {
	coupon := r.FormValue("coupon")
	ccNumber := r.FormValue("cc-number")

	order := Order{
		Coupon:   coupon,
		CcNumber: ccNumber,
	}

	jsonOrder, err := json.Marshal(order)
	if err != nil {
		log.Fatal("Error parsing to json")
	}

	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	err = rabbitMQ.Notify(string(jsonOrder), "application/json", "orders_ex", "")
	if err != nil {
		log.Fatal("Error sending message to queue")
	}
	t := template.Must(template.ParseFiles("templates/process.html"))
	t.Execute(w, "")
}
