package main

import (
	"flag"

	"github.com/teohrt/abrasion/app"
)

func main() {
	seedSite := flag.String("url", "https://www.google.com", "The URL with which Abrasion begins scraping the web")
	regexValue := flag.String("regex", "", "The regular expression Abrasion searches for")
	verbose := flag.Bool("verbose", false, "Sets verbose logging")
	flag.Parse()

	ch := make(chan string)
	searchWithRegex := *regexValue != ""

	config := &app.Config{
		Site:        *seedSite,
		RegexValue:  *regexValue,
		RegexSearch: searchWithRegex,
		Verbose:     *verbose,
		DataChan:    ch,
	}

	app.Start(config)
}
