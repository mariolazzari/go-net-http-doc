package main

import (
	"log"
	"net/http"
)

func main() {
	router := &http.NewServeMux()
	router.HandleFunc("/item/{id}", getItem)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on :8080")
	server.ListenAndServe()
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Item ID: " + id))
}
