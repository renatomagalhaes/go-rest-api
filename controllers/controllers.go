package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renatomagalhaes/go-rest-api/database"
	"github.com/renatomagalhaes/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "[{\"Status\": \"OK\"}]")
}

func GetTeachers(w http.ResponseWriter, r *http.Request) {
	var teacher []models.Teacher
	database.DB.Find(&teacher)
	json.NewEncoder(w).Encode(teacher)
}

func GetTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var teacher models.Teacher
	database.DB.First(&teacher, id)
	json.NewEncoder(w).Encode(teacher)
}

func NewTeacher(w http.ResponseWriter, r *http.Request) {
	var NewTeacher models.Teacher
	json.NewDecoder(r.Body).Decode(&NewTeacher)
	database.DB.Create(&NewTeacher)
	json.NewEncoder(w).Encode(NewTeacher)
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var teacher models.Teacher
	database.DB.Delete(&teacher, id)
	json.NewEncoder(w).Encode(teacher)
}

func EditTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var teacher models.Teacher
	database.DB.First(&teacher, id)
	json.NewDecoder(r.Body).Decode(&teacher)
	database.DB.Save(&teacher)
	json.NewEncoder(w).Encode(teacher)
}
