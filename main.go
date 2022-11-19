package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"

	"mrktplace/controllers"
	"mrktplace/models"
	"mrktplace/templates"
	"mrktplace/views"
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
	usersC.Templates.SignIn = views.Must(views.ParseTemplateFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))
	router.Get("/signup", usersC.New)
	router.Get("/signin", usersC.SignIn)
	router.Post("/signin", usersC.ProcessSignIn)
	router.Post("/users", usersC.Create)
	router.Get("/users/me", usersC.CurrentUser)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	csrfKey := "23jfnrhy57lh6sbnydpe7503khtq230U"
	csrfMiddleware := csrf.Protect([]byte(csrfKey), csrf.Secure(false))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", csrfMiddleware(router))
}
