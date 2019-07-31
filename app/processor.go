package app

import (
	"fmt"
	"net/url"
)

// Scrapes new URLs and logs them
func (c *Config) Process() {
	fmt.Println("Abrasion is scraping...")

	go func() {
		visitedURLs := make(map[string]bool)
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
					if !c.GetEmails {
						c.Logger.Log(u.Host)
					}

					visitedURLs[u.Host] = true
					go c.Scrape(URL)
				}
			}
		}
	}()

	if c.GetEmails {
		go func() {
			emails := make(map[string]bool)
			for {
				select {
				case e := <-c.DataChan:
					if _, exists := emails[e]; !exists {
						emails[e] = true
						c.Logger.Log(e)
					}
				}
			}
		}()
	}
}
