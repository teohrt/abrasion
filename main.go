package main

import (
	"flag"

	"github.com/teohrt/abrasion/app"
)

func main() {
	seedSite := flag.String("url", "https://www.google.com", "The URL with which Abrasion begins scraping the web")
	getEmail := flag.Bool("getEmail", false, "Aggregate email addresses")
	verbose := flag.Bool("verbose", false, "Sets verbose logging")
	flag.Parse()

	app.Start(&app.Config{
		Site:     *seedSite,
		GetEmail: *getEmail,
		Verbose:  *verbose,
	})
}
