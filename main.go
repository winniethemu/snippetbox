package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on :4000")

	// This runs an infinite loop and only returns when an error occurs
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
