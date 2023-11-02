package main

import (
	"github.com/renatomagalhaes/go-rest-api/models"
	"github.com/renatomagalhaes/go-rest-api/routes"
)

func main() {
	//Mock Teachers
	models.Teachers = []models.Teacher{
		{Id: 1, FirstName: "Albert", LastName: "Vignesh", Age: 28, Email: "teacher1@go-api-class.com", Class: 10},
		{Id: 2, FirstName: "Ana", LastName: "Kumar", Age: 22, Email: "teacher2@go-api-class.com", Class: 11},
		{Id: 3, FirstName: "Eric", LastName: "", Age: 25, Email: "teacher3@go-api-class.com", Class: 12},
		{Id: 4, FirstName: "Paul", LastName: "Mark", Age: 24, Email: "teacher4@go-api-class.com", Class: 10},
	}

	routes.HandleRequest()
}
