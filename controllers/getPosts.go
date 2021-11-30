package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"
)

// var posts []models.Post

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Posts)
}
