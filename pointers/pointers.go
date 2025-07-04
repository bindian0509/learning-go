package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", *agePointer)
	fmt.Println("Age Pointer :", agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
	editAgeToAdultYears(agePointer)
	fmt.Println("New Age:", age)

	agePointer = nil
	fmt.Println("Age Pointer :", agePointer)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}

func getAdultYears(age *int) int {
	return *age - 18
}
