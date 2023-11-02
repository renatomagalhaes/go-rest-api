package controllers

import (
	"encoding/json"
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
