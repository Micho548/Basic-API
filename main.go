package main

import (
	// "encoding/json"
	"fmt"

	"net/http"
	"rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Backend running...")

	router := mux.NewRouter()

	routes(router)

	http.ListenAndServe(":8000", router)
}

func routes(router *mux.Router) {
	router.HandleFunc("/post", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/create", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/update/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePost).Methods("DELETE")
}
