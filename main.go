package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Supply a URL and # of desired Go Routines.")
		fmt.Println("Example: 'go run main.go http://google.com 5'")
		return
	}

	URL := os.Args[1]
	concurrency, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Error: A number must be passed.")
		fmt.Println("Example: 'go run main.go http://google.com 5'")
		return
	}

	hasHTTP := strings.Index(URL, "http") == 0
	if !hasHTTP {
		fmt.Println("URL must start with 'http://'")
		return
	}

	URLChan := make(chan string, 100000)
	dataChan := make(chan string, 100000)

	go aggregate(dataChan)

	crawl(concurrency, URL, URLChan, dataChan)

	select {}
}

// Spins up go routines to scrape web pages found from
func crawl(concurrency int, seedURL string, URLChan chan string, dataChan chan string) {
	for i := 0; i < concurrency; i++ {
		go scrape(URLChan, dataChan)
	}

	URLChan <- seedURL
}

// Logs new urls
func aggregate(dataChan chan string) {
	count := 0
	fmt.Println("Logging abrasion:")
	for {
		select {
		case newURL := <-dataChan:
			fmt.Printf("%d: %s\n", count, newURL)
			count++
		}
	}
}
