package main

import "fmt"

func main() {
	// data types declaration and variables

	//var card string = "Ace of Spades"
	/* card := "Ace of Diamonds"
	card = "Ace of Spades"
	card_five := newCard()

	fmt.Println(card)
	fmt.Println(card_five)

	deckSize = 52
	fmt.Println(deckSize)
	fmt.Println(newCard())
	//printState()

	// arrays and looping
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
	colors := []string{"Red", "Green", "Blue", "Yellow"}

	for index, color := range colors {
		fmt.Println(index, color)
	}

	fmt.Println(colors[0:2])
	fmt.Println(colors[:2])
	fmt.Println(colors[2:])
	fmt.Println(colors)
	// some oops
	kards := newDeck()
	kards.print()

	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	c := color("Red")
	fmt.Println(c.describe("ass!"))


	greeting := "Hi, there!"
	fmt.Println([]byte(greeting))


	cards := newDeck()
	//fmt.Println(cards.toString())

	cards.saveToFile("mycards.txt")


	//cards := newDeckFromFile("myCards.txt")

	cards := newDeck()
	cards.shuffle()
	cards.print()

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	for _, num := range nums {

		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}
	*/

	bharat := person{
		firstname: "Bharat",
		lastname:  "Verma",
		contact: contactInfo{
			email:   "bindian0509@gmail.com",
			zipCode: 122001,
		},
	}
	name := "Bill"
	fmt.Println(&name)

	fmt.Println(bharat)
	//bharatPointer := &bharat
	bharat.updateFirstName("Sonu") // - this can work as well thanks to Go!
	bharat.print()

	greetings := []string{"Hi", "there", "I", "want", "to", "rail", "you!"}
	fmt.Println(greetings)
	updateSlice(greetings) // this is blasphemy from Go!
	fmt.Println(greetings)
}
func updateSlice(s []string) {
	s[0] = "Bye"

}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// damn pointer you are scary
func (pointerToPerson *person) updateFirstName(firstname string) {
	(*pointerToPerson).firstname = firstname
}

/*
type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
} */

/*
func newCard() string {
	return "Five of Diamonds"
}
*/
