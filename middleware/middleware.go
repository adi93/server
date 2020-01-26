package middleware

import (
	"log"
	"net/http"

	"server/api"
	"server/config"
)

//Middleware applies a bunch of middleware to a handler
func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

// RedirectHTTPS transfers https to http
func RedirectHTTPS(handler http.Handler) http.Handler {
	if config.HTTPSMode() {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			handler.ServeHTTP(w, r)
		})
	}
	return handler
}

// RequiresLogin ensures that user is logged in before accessing
func RequiresLogin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.LoginWorks() {
			if !api.IsLoggedIn(r) {
				http.Redirect(w, r, "/login", 302)
				return
			}
		}
		handler.ServeHTTP(w, r)
	})
}

// Logger provides a wrapper to log all requests
func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Host, r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
