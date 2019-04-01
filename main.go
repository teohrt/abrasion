package main

import (
	"flag"
)

func main() {

	s := flag.String("url", "https://www.google.com", "The seed URL Abrasion enters the web with.")
	r := flag.String("regex", "", "The regular expression Abrasion searches for.")
	c := make(chan string)
	flag.Parse()

	config := Config{
		site:     s,
		regex:    r,
		dataChan: c,
	}

	run(config)
}
