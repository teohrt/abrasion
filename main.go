package main

import (
	"flag"
	"math"

	"github.com/teohrt/abrasion/app"
)

func main() {
	url := flag.String("url", "https://www.google.com", "The URL with which Abrasion begins scraping the web")
	getEmail := flag.Bool("getEmail", false, "Aggregate email addresses")
	verbose := flag.Bool("verbose", false, "Sets verbose logging")
	debug := flag.Bool("debug", false, "Sets debug level logging")
	scrapeLimit := flag.Int("scrapeLimit", math.MaxInt32, "Sets the number of URLs to scrape. Defaults to MAXINT")
	flag.Parse()

	app.Start(&app.Config{
		SeedURL:     *url,
		ScrapeLimit: *scrapeLimit,
		GetEmail:    *getEmail,
		Verbose:     *verbose,
		Debug:       *debug,
	})
}
