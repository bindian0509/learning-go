package main

import (
	"fmt"
	"net/http"
	"time"
)
func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, done_chan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done_chan <- true
}
func main() {


	go greet("How are you?")
	go greet("Nice to meet you!")
	go greet("I hope you're liking the course!")
	channel := make (chan bool)
	go slowGreet("How ... are ... you ...?", channel)
	fmt.Println(<-channel)

	/*

	links := []string{
		"https://www.google.fr",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amazon.in",
		"https://www.orkut.com",
	}

	fmt.Println(links)

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
	*/
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "website might be down. Error : ", err)
		c <- link
		return
	}
	fmt.Println(link, "is up!!!")
	c <- link
}
