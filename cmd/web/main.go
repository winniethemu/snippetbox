package main

import (
	"log"
	"net/http"
)

func main() {
	// 1. Request URLs are automatically sanitized, e.g /foo/bar/..//baz redirects to /foo/baz
	// 2. If /foo/ is registered, then /foo redirects to /foo/
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home) // Restricting subtree path matching
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	// This runs an infinite loop and only returns when an error occurs
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
