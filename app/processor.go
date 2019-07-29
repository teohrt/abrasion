package app

import (
	"fmt"
	"net/url"
)

// Scrapes new URLs and logs them
func (c *Config) Process() {
	visitedURLs := make(map[string]bool)

	fmt.Println("Abrasion is scraping...")

	go func() {
		for {
			select {
			case URL := <-c.URLChan:
				u, err := url.Parse(URL)
				if err != nil {
					c.Logger.Err("Error parsing URL. : " + URL)
					continue
				}

				// If hostname hasn't already been visited
				if _, exists := visitedURLs[u.Host]; !exists {
					if !c.GetEmail {
						c.Logger.Log(u.Host)
					}

					visitedURLs[u.Host] = true
					go c.Scrape(URL)
				}
			}
		}
	}()

	if c.GetEmail {
		go func() {
			for {
				select {
				case result := <-c.DataChan:
					c.Logger.Log(result)
				}
			}
		}()
	}
}
