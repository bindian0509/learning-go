package main

import (
	"fmt"
	"net/http"
	"github.com/justinas/nosurf"
)

// WriteToConsole w.r.t. every request 
func WriteToConsole(next http.Handler) http.Handler {	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page...")
		next.ServeHTTP(w,r)
	})
}
// NoSurf adds CSRF protection to all POST request 
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
		
	})
	return csrfHandler;
}

// SessionLoad loads and save sessions on every requests 
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)	
}
