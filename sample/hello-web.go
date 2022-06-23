package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError) // tag 3.0
	w.Write([]byte("500 - Something bad happened!"))
	// fmt.Fprintf(w, "I am a GO application version 3 running in a Docker image.")
}

func main() {
	fmt.Println("Trivial web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
