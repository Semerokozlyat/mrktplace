package controllers

import "net/http"

type Template interface {
	Execute(rw http.ResponseWriter, data interface{})
}
