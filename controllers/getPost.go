package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"

	"github.com/gorilla/mux"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range models.Posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(models.Post{})
}
