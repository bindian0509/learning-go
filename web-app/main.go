package main

import (
	"fmt"
	"net/http"
)
// open the link - http://localhost:8080/

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes Written : %d", n))
	})

	http.ListenAndServe(":8080", nil)
}