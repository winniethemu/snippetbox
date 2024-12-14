package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// 1. Request URLs are automatically sanitized, e.g /foo/bar/..//baz redirects to /foo/baz
	// 2. If /foo/ is registered, then /foo redirects to /foo/
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home) // Restricting subtree path matching
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// Static assets
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
