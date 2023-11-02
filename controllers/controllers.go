package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/renatomagalhaes/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "[{\"Status\": \"OK\"}]")
}

func GetTeachers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Teachers)
}

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, teacher := range models.Teachers {
		if strconv.Itoa(teacher.Id) == id {
			json.NewEncoder(w).Encode(teacher)
		}
	}
}
