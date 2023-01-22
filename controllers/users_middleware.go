package controllers

import (
	"net/http"

	"mrktplace/models"
	"mrktplace/requestcontext"
)

type UserMiddleware struct {
	SessionService *models.SessionService
}

func (um UserMiddleware) SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("session")
		if err != nil {
			next.ServeHTTP(rw, r)
			return
		}
		user, err := um.SessionService.User(tokenCookie.Value)
		if err != nil {
			next.ServeHTTP(rw, r)
			return
		}
		reqCtx := r.Context()
		r = r.WithContext(requestcontext.WithUser(reqCtx, user))
		next.ServeHTTP(rw, r)
	})
}

func (um UserMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := requestcontext.User(r.Context())
		if user == nil {
			http.Redirect(rw, r, "/signin", http.StatusFound)
			return
		}
		next.ServeHTTP(rw, r)
	})
}
