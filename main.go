package main

import (
	"flag"
	"math"

	"github.com/teohrt/abrasion/app"
)

func main() {
	url := flag.String("url", "https://www.google.com", "The URL with which Abrasion begins scraping the web")
	limit := flag.Int("scrapeLimit", math.MaxInt32, "Sets the number of URLs to scrape. Defaults to MAXINT")
	emails := flag.Bool("getEmails", false, "Aggregate email addresses")
	verbose := flag.Bool("verbose", false, "Sets verbose logging")
	debug := flag.Bool("debug", false, "Sets debug level logging")
	flag.Parse()

	app.Start(&app.Config{
		SeedURL:     *url,
		ScrapeLimit: *limit,
		GetEmails:   *emails,
		Verbose:     *verbose,
		Debug:       *debug,
	})
}
