package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	reg := []string{"a", "b", "c"}
	fmt.Println(strings.Join(reg[:], ""))

	var firstname *string = new(string)
	*firstname = "Thomas"

	ptr := &firstname

	fmt.Println(*firstname, firstname, ptr, *ptr)

	const pi float32 = 3.14
	fmt.Println(pi + 2)

	// call by reference
	color := "red"

	log.Println("light before", color)
	changeViaPointer(&color)
	log.Println("light after", color)
}

func changeViaPointer(s *string) {
	*s = "green"
}
