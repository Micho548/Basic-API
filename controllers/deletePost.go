package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"

	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Posts {
		if item.ID == params["id"] {
			models.Posts = append(models.Posts[:index], models.Posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.Posts)
}
