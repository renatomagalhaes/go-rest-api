package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/renatomagalhaes/go-rest-api/controllers"
	"github.com/renatomagalhaes/go-rest-api/middleware"
)

func HandleRequest() {
	port := 8080

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/teachers", controllers.GetTeachers).Methods("Get")
	r.HandleFunc("/api/teachers/{id}", controllers.GetTeacher).Methods("Get")
	r.HandleFunc("/api/teachers", controllers.NewTeacher).Methods("Post")
	r.HandleFunc("/api/teachers/{id}", controllers.DeleteTeacher).Methods("Delete")
	r.HandleFunc("/api/teachers/{id}", controllers.EditTeacher).Methods("Put")

	log.Printf("Iniciando sevidor na porta %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
