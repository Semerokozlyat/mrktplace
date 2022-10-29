package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(rw http.ResponseWriter, templatePath string) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("parsing template file: %v", err)
		http.Error(rw, "Failed to parse template file", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(rw, nil)
	if err != nil {
		log.Printf("executing template file: %v", err)
		http.Error(rw, "Failed to execute template file", http.StatusInternalServerError)
		return
	}
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(rw, tplPath)
}

func contactHandler(rw http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(rw, tplPath)
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
