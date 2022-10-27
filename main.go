package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(rw, "<h1>Start page</h1>")
}

func contactHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(rw, "<h1>Contact page.</h1>"+
		"<p>Email at <a href=\"mailto:example@example.com\">example@example.com</p>")
}

func main() {
	router := chi.NewRouter()
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.NotFound(func(rw http.ResponseWriter, r *http.Request) {
		http.Error(rw, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting HTTP server on :3000...")
	http.ListenAndServe(":3000", router)
}
