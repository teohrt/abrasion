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

	outputWriter := newCSVWriter(fileName)
	defer outputWriter.Flush()

	visitedURLs := make(map[string]bool)

	fmt.Println("Abrasion is scraping...")
	for {
		select {
		case URLString := <-c.dataChan:
			u, err := url.Parse(URLString)
			if err != nil {
				c.logMsg("Error parsing URL. : " + URLString)
				continue
			}

			// If hostname hasn't already been visited
			if _, exists := visitedURLs[u.Host]; !exists {
				c.logMsg(u.Host)

				visitedURLs[u.Host] = true

				go c.scrape(URLString)

				err = outputWriter.Write([]string{URLString})
				if err != nil {
					c.logMsg("Cannot write URL to file. : " + URLString)
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

func (c *config) logMsg(s string) {
	if c.verbose {
		fmt.Println(s)
	}
}
