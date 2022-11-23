package controllers

import "net/http"

const (
	CookieSession = "session"
)

func newCookie(name, value string) *http.Cookie {
	c := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	return &c
}

func setCookie(rw http.ResponseWriter, name, value string) {
	c := newCookie(name, value)
	http.SetCookie(rw, c)
}
