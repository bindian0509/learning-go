package main

import (
	"errors"
	"fmt"
)



// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

const (
	SuccessMsg string = "Operation successful!"
	PanicMsg   string = "Operation panicked!"
)

// Person is a name and age tuple.
type Person struct {
	Name string
	Age  int
}

// Persons saves a map of Person.
type Persons struct {
	m map[string]Person
}

// NewPersons initializes the Persons type.
func NewPersons() *Persons {
	return &Persons{
		m: make(map[string]Person),
	}
}

// PanickingAdd adds a new Person to the Persons receiver and
// panics in the case of invalid ages.
func (p *Persons) PanickingAdd(np Person) {
	if np.Age < 0 || np.Age > 100 {
		panic(fmt.Sprintf("add:invalid age %d for name %s", np.Age, np.Name))
	}
	p.m[np.Name] = np
}

// PopulateData should add data to the Persons type without panicking.
func PopulateData(data []Person) (result string) {
	result = SuccessMsg
	// Your code here
	p := NewPersons()
	for _, np := range data {
		p.PanickingAdd(np)
	}
	return
}



func meanCalc  (nums [] int ) (*float64, error) {
	if len(nums) == 0 {
		return nil, errors.New("empty or nil array")
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	mean := float64(sum) / float64(len(nums))
	return &mean, nil
	//return nil, errors.New("not implemented")
}


func main () {
	arr := [] int {1,2,3,4,5,6,7,8,9}
	mean, err := meanCalc(arr)
	if err == nil {
		fmt.Println("mean : ", *mean, " for the arr : ", arr )
	} else {
		fmt.Println("error : ", err, " for array : ", arr)
	}
	arr2 := [] int {}
	mean2, err2 := meanCalc(arr2)
	if err2 == nil {
		fmt.Println("mean : ", *mean2)
	} else {
		fmt.Println("error : ", err2, " for array : ", arr2)
	}
}
