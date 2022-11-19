package controllers

import "net/http"

type Template interface {
	Execute(rw http.ResponseWriter, r *http.Request, data interface{})
}
