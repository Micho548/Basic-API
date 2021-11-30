package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"

	"github.com/gorilla/mux"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Posts {
		if item.ID == params["id"] {
			models.Posts = append(models.Posts[:index], models.Posts[index+1:]...)

			var post models.Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			models.Posts = append(models.Posts, post)
			json.NewEncoder(w).Encode(&post)

			return
		}
	}

	json.NewEncoder(w).Encode(models.Posts)
}
