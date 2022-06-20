package main

import (
	"errors"
	"log"
)

func main () {
	result, err := divide(100.00, 0)
	if(err != nil) {
		log.Println(err)
		return
	}
	log.Println("return of the division : ", result)
}

func divide (x,y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("Who the eff divides by zero...")
	}
	result = x/y 
	return result, nil
}