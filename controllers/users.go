package controllers

import (
	"fmt"
	"net/http"

	"mrktplace/models"
)

type Users struct {
	Templates struct {
		New Template
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
