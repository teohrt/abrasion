package app

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Site        string
	RegexValue  string
	RegexSearch bool
	Verbose     bool
	DataChan    chan string
}

func Start(c *Config) {
	validate(c)
	go c.aggregate()
	c.scrape(c.Site)
}

func validate(c *Config) {
	hasHTTP := strings.Index(c.Site, "http") == 0
	if !hasHTTP {
		fmt.Println("URL must start with 'http://'")
		os.Exit(1)
	}
}

func (c *Config) log(s string) {
	if c.Verbose {
		fmt.Println(s)
	}
}
