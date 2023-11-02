package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/renatomagalhaes/go-rest-api/controllers"
)

func HandleRequest() {
	port := 8080

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/teachers", controllers.AllTeachers)

	log.Printf("Iniciando sevidor na porta %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
