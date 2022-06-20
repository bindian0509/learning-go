package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

const portNum = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "THIS IS MY HOME PAGE .... YEAH RIGHT!! ")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	rand := RandomNum(9, 999)
	fmt.Fprintf(w, fmt.Sprintf("THIS IS ABOUT PAGE .... XoXo %d ", rand))
}

// AppValues adds two integers
func AppValues(x, y int) int {
	var sum int
	sum = x + y
	return sum
}

func RandomNum(min, max int) int {
	return rand.Intn(max-min) + min
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 69.87
	var y float32 = 12.00
	div, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
	}

	fmt.Fprintf(w, fmt.Sprintf("Some sort of dividing... where x = %f divided by y = %f gives us %f ", x, y, div))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Wrong divisor greater than zero expected ...")
		return 0, err
	}
	return x / y, nil
}

// main - open the link - http://localhost:8080/
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	/* http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes Written : %d", n))
	}) */

	fmt.Println(fmt.Sprintf("Starting web-app application the first go web application on port %s", portNum))
	http.ListenAndServe(portNum, nil)
}
