package main

import (
	"fmt"
	"mrktplace/controllers"
	"mrktplace/views"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	tpl := views.Must(views.ParseTemplateFile(filepath.Join("templates", "home.gohtml")))
	router.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseTemplateFile(filepath.Join("templates", "contact.gohtml")))
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseTemplateFile(filepath.Join("templates", "faq.gohtml")))
	router.Get("/faq", controllers.StaticHandler(tpl))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
