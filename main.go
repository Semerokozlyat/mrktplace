package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"mrktplace/controllers"
	"mrktplace/models"
	"mrktplace/templates"
	"mrktplace/views"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	tpl := views.Must(views.ParseTemplateFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	router.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseTemplateFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseTemplateFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	router.Get("/faq", controllers.FAQ(tpl))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	us := &models.UserService{
		DB: db,
	}
	usersC := controllers.Users{
		UserService: us,
	}
	usersC.Templates.New = views.Must(views.ParseTemplateFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	router.Get("/signup", usersC.New)
	router.Post("/users", usersC.Create)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
