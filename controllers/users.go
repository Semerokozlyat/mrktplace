package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(rw http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(rw, data)
}

func (u Users) Create(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Email: ", r.FormValue("email"))
	fmt.Fprint(rw, "Password: ", r.FormValue("password"))
}
