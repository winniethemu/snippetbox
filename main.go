package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// 1. Request URLs are automatically sanitized, e.g /foo/bar/..//baz redirects to /foo/baz
	// 2. If /foo/ is registered, then /foo redirects to /foo/
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home) // Restricting subtree path matching
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	// This runs an infinite loop and only returns when an error occurs
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
