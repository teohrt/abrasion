package app

import (
	"fmt"
	"net/url"
	"sync"
)

// Scrapes new URLs and logs them
func (c *Config) Process(wg *sync.WaitGroup) {
	defer c.ErrorLogger.FlushLogger()
	defer c.ResultLogger.FlushLogger()

	visitedURLs := make(map[string]bool)

	fmt.Println("Abrasion is scraping...")
	for {
		select {
		case URLString := <-c.URLChan:
			u, err := url.Parse(URLString)
			if err != nil {
				c.ErrorLogger.Log("Error parsing URL. : " + URLString)
				continue
			}

			// If hostname hasn't already been visited
			if _, exists := visitedURLs[u.Host]; !exists {
				c.ResultLogger.Log(u.Host)

				visitedURLs[u.Host] = true

				go c.Scrape(URLString)
			}
		}
	}

	wg.Done()
}
