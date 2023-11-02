package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/renatomagalhaes/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "[{\"Status\": \"OK\"}]")
}

func AllTeachers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Teachers)
}
