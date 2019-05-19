package app

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

// Scrapes new URLs and logs them
func (c *Config) aggregate() {
	currentTime := time.Now()
	fileName := "AbrasionResult-" + currentTime.Format("2006-01-02 3:4:5 pm") + ".csv"

	outputWriter := newCSVWriter(fileName)
	defer outputWriter.Flush()

	visitedURLs := make(map[string]bool)

	fmt.Println("Abrasion is scraping...")
	for {
		select {
		case URLString := <-c.DataChan:
			u, err := url.Parse(URLString)
			if err != nil {
				c.log("Error parsing URL. : " + URLString)
				continue
			}

			// If hostname hasn't already been visited
			if _, exists := visitedURLs[u.Host]; !exists {
				c.log(u.Host)

				visitedURLs[u.Host] = true

				go c.scrape(URLString)

				err = outputWriter.Write([]string{URLString})
				if err != nil {
					c.log("Cannot write URL to file. : " + URLString)
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
