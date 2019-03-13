package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Supply a URL.")
		fmt.Println("Example: 'go run main.go http://google.com'")
		return
	}

	URL := os.Args[1]

	hasHTTP := strings.Index(URL, "http") == 0
	if !hasHTTP {
		fmt.Println("URL must start with 'http://'")
		return
	}

	dataChan := make(chan string)

	go aggregate(dataChan)

	scrape(URL, dataChan)
}
