package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatomagalhaes/go-rest-api/controllers"
)

func HandleRequest() {
	port := 8080

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/teachers", controllers.GetTeachers).Methods("Get")
	r.HandleFunc("/api/teachers/{id}", controllers.GetTeacher).Methods("Get")

	log.Printf("Iniciando sevidor na porta %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
