package main

import (
	"fmt"
	"net/http"
)

const portNum = ":8080"

// main - open the link - http://localhost:8080/
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting web-app application the first go web application on port %s", portNum))
	http.ListenAndServe(portNum, nil)
}
