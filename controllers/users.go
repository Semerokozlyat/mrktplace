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
	u.Templates.New.Execute(rw, nil)
}

func (u Users) Create(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Email: ", r.FormValue("email"))
	fmt.Fprint(rw, "Password: ", r.FormValue("password"))
}
