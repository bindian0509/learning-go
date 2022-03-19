package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/bindian0509/learning-go/helpers"
)

type User struct {
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber string
	BirthDate   time.Time
}
type Animal interface {
	Says() string
	NoOfLegs() int
}

type Cat struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name       string
	Color      string
	NoOfTeeths int
}

func main() {

	user := User{
		FirstName: "Bharat",
		LastName:  "Verma",
	}
	log.Println(user.FirstName)
	log.Println(user.LastName)
	log.Println(user.BirthDate)
	log.Println(user.Age)

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

	myMap := make(map[string]string)
	myMap["AK47"] = "7.62"
	myMap["M16"] = "5.56"

	log.Println(myMap)

	var mySlice []int

	mySlice = append(mySlice, 27)
	mySlice = append(mySlice, 81)
	mySlice = append(mySlice, 9)
	mySlice = append(mySlice, 729)

	log.Println(mySlice)
	sort.Ints(mySlice)
	log.Println(mySlice)

	nums := []int{3, 6, 9, 1, 3, 0, 5, 3, 4}
	log.Println(nums[2:])

	sentence := "My Name is Anthony Gonsalves"

	for i, l := range sentence {
		log.Println("index : ", i, "and value", l)
	}

	cat := Cat{
		Name:  "Michael",
		Breed: "Persian Cat",
	}
	PrintInfo(&cat)

	// package import
	sometype := helpers.SomeType{
		TypeName:    "ABC",
		TypeAddress: "FINE",
	}

	log.Println(sometype)

	// channels

	intChan := make(chan int)
	defer close(intChan)
	go CalcValue(intChan)
	num := <-intChan
	log.Println("int chan magic num ", num)
}

const numPool = 100

func CalcValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func PrintInfo(a Animal) {
	fmt.Println("This animal say : ", a.Says(), " and have legs ", a.NoOfLegs())
}

func (c *Cat) Says() string {
	return "meow"
}

func (c *Cat) NoOfLegs() int {
	return 4
}

func changeViaPointer(s *string) {
	*s = "green"
}

// make public with a caps start
func Whatever() {

	log.Println("Whatever")
}
