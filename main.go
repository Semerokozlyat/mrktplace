package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"

	"mrktplace/controllers"
	"mrktplace/migrations"
	"mrktplace/models"
	"mrktplace/templates"
	"mrktplace/views"
)

func main() {
	// Setup DB
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFromFS(db, migrations.FS, ".")
	if err != nil {
		panic(fmt.Errorf("apply DB migrations: %s", err))
	}

	// Setup services
	userService := &models.UserService{
		DB: db,
	}
	sessService := &models.SessionService{
		DB: db,
	}

	// Setup middleware
	csrfKey := "23jfnrhy57lh6sbnydpe7503khtq230U"
	csrfMiddleware := csrf.Protect([]byte(csrfKey), csrf.Secure(false))

	userMiddleware := controllers.UserMiddleware{
		SessionService: sessService,
	}

	// Setup controllers
	usersC := controllers.Users{
		UserService:    userService,
		SessionService: sessService,
	}
	usersC.Templates.New = views.Must(views.ParseTemplateFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseTemplateFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))

	// Setup router
	router := chi.NewRouter()
	router.Use(csrfMiddleware, userMiddleware.SetUser)

	tpl := views.Must(views.ParseTemplateFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	router.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseTemplateFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseTemplateFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	router.Get("/faq", controllers.FAQ(tpl))

	router.Get("/signup", usersC.New)
	router.Get("/signin", usersC.SignIn)
	router.Post("/signin", usersC.ProcessSignIn)
	router.Post("/signout", usersC.ProcessSignOut)
	router.Post("/users", usersC.Create)

	router.Route("/users/me", func(r chi.Router) {
		r.Use(userMiddleware.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
