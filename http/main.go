package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://insurance.mobikwik.com/api/home/ping")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
