# Abrasion: The CLI Web Scraper
Abrasion is used to crawl the web given a website as a starting point. It will concurrently scrape all websites found, starting with a seeded website. While scraping, Abrasion can also search and output matches to a regex template that defaults to match on email addresses, or can just simply print the structure of the world wide web.

## Installation Instructions
```bash
git clone https://github.com/teohrt/abrasion.git
cd abrasion/
make build
```

## Usage
```bash
# Crawl the web, starting at the specified entry point, aggregating emails
./abrasion -url=https://reddit.com -getEmail

# Logs the first 100 sites encountered with the given entry point
./abrasion -url=https://reddit.com -scrapeLimit=100 -verbose 
```

## Flags
* url - The URL with which Abrasion begins scraping the web. Defaults to "https://www.google.com" when not set.
* verbose - Sets verbose logging. Defaults to false when not set.
* getEmail - Aggregate email addresses. Defaults to false when not set.
* scrapeLimit - Sets the number of URLs to scrape. Defaults to MAXINT when not set.

## Logs
Abrasion has 3 potential output streams. By default it outputs errors and results to their own respective .CSV files and will also print to console if the -verbose flag is set.