package controllers

import (
	"fmt"
	"net/http"

	"mrktplace/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}

	UserService    *models.UserService
	SessionService *models.SessionService
}

func (u Users) New(rw http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(rw, r, data)
}

func (u Users) Create(rw http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, "User cannot be created", http.StatusInternalServerError)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(rw, r, "/signin", http.StatusFound)
		return
	}
	setCookie(rw, CookieSession, session.Token)
	http.Redirect(rw, r, "/users/me", http.StatusFound)
	fmt.Fprintf(rw, "User created: %+v", user)
}

func (u Users) SignIn(rw http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(rw, r, data)
}

func (u Users) ProcessSignIn(rw http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, "User authentication failed", http.StatusUnauthorized)
		return
	}
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Failed to create user session", http.StatusInternalServerError)
		return
	}
	setCookie(rw, CookieSession, session.Token)
	http.Redirect(rw, r, "/users/me", http.StatusFound)
}

func (u Users) CurrentUser(rw http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Failed to read session cookie: " + err.Error())
		http.Redirect(rw, r, "/signin", http.StatusFound)
		return
	}
	user, err := u.SessionService.User(tokenCookie.Value)
	if err != nil {
		fmt.Println("Failed to get user by session cookie: " + err.Error())
		http.Redirect(rw, r, "/signin", http.StatusFound)
		return
	}
	fmt.Fprintf(rw, "Current user: %+v", user)
}

func (u Users) ProcessSignOut(rw http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(rw, r, "/signin", http.StatusFound)
		return
	}
	err = u.SessionService.Delete(tokenCookie.Value)
	if err != nil {
		fmt.Println("Failed to delete session by session cookie: " + err.Error())
		http.Error(rw, "failed to delete session by session cookie", http.StatusInternalServerError)
		return
	}
	deleteCookie(rw, "session")
	http.Redirect(rw, r, "/signin", http.StatusFound)
}
