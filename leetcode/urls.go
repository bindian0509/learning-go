package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchAll(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		u := url // shadow to avoid capturing the loop variable
		go func(target string) {
			defer wg.Done()

			resp, err := http.Get(target)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", target, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Printf("Error: %s returned status %d\n", target, resp.StatusCode)
				return
			}

			fmt.Printf("Fetched %s with status %d\n", target, resp.StatusCode)
		}(u)
	}

	wg.Wait()
}
