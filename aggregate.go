package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

// Scrapes new URLs and logs them
func (c *config) aggregate() {
	currentTime := time.Now()
	fileName := "AbrasionResult-" + currentTime.Format("2006-01-02 3:4:5 pm") + ".csv"

	logger := newCSVWriter(fileName)
	defer logger.Flush()

	visitedURLs := make(map[string]bool)

	fmt.Println("Abrasion in progress...")
	for {
		select {
		case URLString := <-c.dataChan:
			u, err := url.Parse(URLString)
			if err != nil {
				fmt.Println("Error parsing URL. :", URLString)
				continue
			}

			// If hostname hasn't already been visited
			if _, exists := visitedURLs[u.Host]; !exists {
				fmt.Println(u.Host)
				visitedURLs[u.Host] = true

				go c.scrape(URLString)

				err = logger.Write([]string{URLString})

				if err != nil {
					fmt.Println("Cannot write URL to file. : ", URLString)
				}
			}
		}
	}
}

func newCSVWriter(filename string) *csv.Writer {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error with csv.", err)
	}

	return csv.NewWriter(file)
}
