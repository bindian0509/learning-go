package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

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

	/* for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	} */

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

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
