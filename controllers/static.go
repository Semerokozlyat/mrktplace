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

func FAQ(tpl views.Template) http.HandlerFunc {
	qaData := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Do you provide support?",
			Answer:   "No",
		},
		{
			Question: "Do you plan to provide support in the future?",
			Answer:   "Yes",
		},
		{
			Question: "Where your office is located?",
			Answer:   "Our entire team is remote",
		},
	}
	return func(rw http.ResponseWriter, r *http.Request) {
		tpl.Execute(rw, qaData)
	}
}
