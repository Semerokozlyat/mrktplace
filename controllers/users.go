package controllers

import (
	"mrktplace/views"
	"net/http"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(rw http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(rw, nil)
}
