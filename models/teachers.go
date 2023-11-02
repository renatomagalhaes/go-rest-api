package models

type Teacher struct {
	ID        int    `json:"id,string"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Class     int    `json:"class"`
}

var Teachers []Teacher
