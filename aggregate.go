package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
)

// Scrapes new URLs and logs them
func aggregate(dataChan chan string) {
	logger := newCSVWriter("result.csv")
	defer logger.Flush()

	visited := make(map[string]int)

	fmt.Println("Abrasion in progress...")
	for {
		select {
		case URLString := <-dataChan:
			u, err := url.Parse(URLString)
			if err != nil {
				fmt.Println("Error parsing URL. :", URLString)
				continue
			}

			// If hostname hasn't already been visited
			if _, exists := visited[u.Host]; !exists {
				fmt.Println(u.Host)
				visited[u.Host] = 1

				go scrape(URLString, dataChan)

				err = logger.Write([]string{URLString})

				if err != nil {
					fmt.Println("Cannot write URL to file. : ", URLString)
				}
			} else {
				// Incremenet hit count for popularity tracking
				visited[u.Host] = visited[u.Host] + 1
			}
		}
	}
}

func newCSVWriter(filename string) *csv.Writer {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Error with csv.", err)
	}

	return csv.NewWriter(file)
}
