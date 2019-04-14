package main

import (
	"fmt"
	"os"
	"strings"
)

type config struct {
	site        string
	regexValue  string
	regexSearch bool
	verbose     bool
	dataChan    chan string
}

func validate(c *config) {
	hasHTTP := strings.Index(c.site, "http") == 0
	if !hasHTTP {
		fmt.Println("URL must start with 'http://'")
		os.Exit(1)
	}
}

func run(c *config) {
	validate(c)

	go c.aggregate()

	c.scrape(c.site)
}
