package main

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	site     *string
	regex    *string
	dataChan chan string
}

func validate(c *Config) {
	hasHTTP := strings.Index(*c.site, "http") == 0
	if !hasHTTP {
		fmt.Println("URL must start with 'http://'")
		os.Exit(1)
	}
}

func run(config Config) {
	validate(&config)

	go aggregate(config)

	scrape(*config.site, config.dataChan)
}
