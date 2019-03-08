package main

import (
	"encoding/csv"
	"fmt"
	"log"
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

// Scrapes new URLs and logs them
func aggregate(dataChan chan string) {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Error with csv.", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fmt.Println("Abrasion in progress...")
	for {
		select {
		case newURL := <-dataChan:
			go scrape(newURL, dataChan)

			err := writer.Write([]string{newURL})

			if err != nil {
				fmt.Println("Cannot write URL to file. : ", newURL)
			}
		}
	}
}
