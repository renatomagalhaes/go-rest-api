package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	ID   int    `json:"id,string"`
	Code string `json:"-"`
	Name string `json:"name,omitempty"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	listCustomers := []Customer{
		{1, "001", "Albert"},
		{2, "002", "Ana"},
		{3, "003", ""},
		{13, "013", "Paul"},
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(&listCustomers)
}

func HandleRequest() {
	port := 8080
	http.HandleFunc("/", Home)
	log.Printf("Iniciando sevidor na porta %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func main() {
	HandleRequest()
}
