package main

import (
	"flag"
)

func main() {
	seed := flag.String("url", "https://www.google.com", "The seed URL Abrasion begins scraping the web with.")
	reg := flag.String("regex", "", "The regular expression Abrasion searches for.")
	v := flag.Bool("verbose", false, "Sets verbose logging.")
	ch := make(chan string)
	flag.Parse()

	searchWithRegex := *reg != ""

	config := &config{
		site:        *seed,
		regexValue:  *reg,
		regexSearch: searchWithRegex,
		verbose:     *v,
		dataChan:    ch,
	}

	run(config)
}
