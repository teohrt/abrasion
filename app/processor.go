package app

import (
	"fmt"
	"net/url"
)

// Scrapes new URLs and logs them
func (c *Config) Process() {
	defer c.ErrorLogger.FlushLogger()
	defer c.ResultLogger.FlushLogger()

	visitedURLs := make(map[string]bool)

	fmt.Println("Abrasion is scraping...")

	go func() {
		for {
			select {
			case URL := <-c.URLChan:
				u, err := url.Parse(URL)
				if err != nil {
					c.ErrorLogger.Log("Error parsing URL. : " + URL)
					continue
				}

				// If hostname hasn't already been visited
				if _, exists := visitedURLs[u.Host]; !exists {
					if !c.GetEmail {
						c.ResultLogger.Log(u.Host)
					}

					visitedURLs[u.Host] = true

					go c.Scrape(URL)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case result := <-c.DataChan:
				if c.GetEmail {
					c.ResultLogger.Log(result)
				}
			}
		}
	}()
}
