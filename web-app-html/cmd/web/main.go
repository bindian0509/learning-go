package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
	"github.com/bindian0509/learning-go/web-app-html/pkg/config"
	"github.com/bindian0509/learning-go/web-app-html/pkg/handlers"
	"github.com/bindian0509/learning-go/web-app-html/pkg/render"
)

// port for debugging
const portNum = ":8080"
// get the template cache from app config 
var app config.AppConfig
// session for variable shadowing 
var session *scs.SessionManager

// main - open the link - http://localhost:8080/
func main() {

	// change this to true in production
	app.InProduction = false
	// remove the fucking : to prevent re-initialization 
	session = scs.New()
	
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Error creating template cache ", err)
	}
	app.TemplateCache = tc 
	app.UseCache = false 
	
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting web-app application the first go web application on port %s", portNum)
	//http.ListenAndServe(portNum, nil)
	serve := &http.Server{
		Addr: portNum,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
