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

	UserService *models.UserService
}

func (u Users) New(rw http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(rw, data)
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
	fmt.Fprintf(rw, "User created: %+v", user)
}

func (u Users) SignIn(rw http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(rw, data)
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
	fmt.Fprintf(rw, "User authenticated: %+v", user)
}
