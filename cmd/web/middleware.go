package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// uses Cookies to make sure the csrf Token is available per page
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		// for the whole page
		Path: "/",
		// if HTTPS is in use, then true, otherwise it's false
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
