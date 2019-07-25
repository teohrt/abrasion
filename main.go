package main

import (
	"flag"

	"github.com/teohrt/abrasion/app"
)

func main() {
	url := flag.String("url", "https://www.google.com", "The URL with which Abrasion begins scraping the web")
	regexValue := flag.String("regex", "", "The regular expression Abrasion searches for")
	verbose := flag.Bool("verbose", false, "Sets verbose logging")
	flag.Parse()

	app.Start(&app.Config{
		SeedURL:    *url,
		RegexValue: *regexValue,
		Verbose:    *verbose,
	})
}
