package main

import (
	"fmt"
	"log"
	"mrktplace/views"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(rw http.ResponseWriter, templatePath string) {
	t, err := views.ParseTemplateFile(templatePath)
	if err != nil {
		log.Printf("parsing template file: %v", err)
		http.Error(rw, "Failed to parse template file", http.StatusInternalServerError)
		return
	}
	t.Execute(rw, nil)
}

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(rw, tplPath)
}

func contactHandler(rw http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(rw, tplPath)
}

func faqHandler(rw http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(rw, tplPath)
}

func main() {
	router := chi.NewRouter()
	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.NotFound(func(rw http.ResponseWriter, r *http.Request) {
		http.Error(rw, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting HTTP server on :3000...")
	http.ListenAndServe(":3000", router)
}
