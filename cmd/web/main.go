package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// 1. Request URLs are automatically sanitized, e.g /foo/bar/..//baz redirects to /foo/baz
	// 2. If /foo/ is registered, then /foo redirects to /foo/
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home) // Restricting subtree path matching
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Static assets
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	logger.Info("starting server", "addr", *addr)

	// This runs an infinite loop and only returns when an error occurs
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
