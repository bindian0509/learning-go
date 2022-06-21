package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bindian0509/learning-go/web-app-html/pkg/config"
	"github.com/bindian0509/learning-go/web-app-html/pkg/handlers"
	"github.com/bindian0509/learning-go/web-app-html/pkg/render"
)

const portNum = ":8080"

// main - open the link - http://localhost:8080/
func main() {
	// get the template cache from app config 
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Error creating template cache ", err)
	}
	app.TemplateCache = tc 
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting web-app application the first go web application on port %s", portNum))
	http.ListenAndServe(portNum, nil)
}
