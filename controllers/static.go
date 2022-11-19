package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tpl.Execute(rw, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
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
		tpl.Execute(rw, r, qaData)
	}
}
