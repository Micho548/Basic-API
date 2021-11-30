package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"rest-api/models"
	"strconv"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	models.Posts = append(models.Posts, post)
	json.NewEncoder(w).Encode(&post)
}
