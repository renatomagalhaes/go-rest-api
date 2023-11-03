package main

import (
	"github.com/renatomagalhaes/go-rest-api/database"
	"github.com/renatomagalhaes/go-rest-api/routes"
)

func main() {
	database.DbConnect()
	routes.HandleRequest()
}
