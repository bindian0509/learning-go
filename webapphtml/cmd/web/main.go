package main

import (
	"fmt"
	"net/http"
	"github.com/bindian0509/learning-go/webapphtml/pkg/handlers"
)

const portNum = ":8080"

// main - open the link - http://localhost:8080/
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting web-app application the first go web application on port %s", portNum))
	http.ListenAndServe(portNum, nil)
}
