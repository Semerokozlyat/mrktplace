package controllers

import (
	"mrktplace/views"
	"net/http"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tpl.Execute(rw, nil)
	}
}
