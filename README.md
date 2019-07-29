# Abrasion: The CLI Web Scraper
Abrasion is used to crawl the web. It will concurrently scrape all websites found, starting with a seed website. While scraping, Abrasion can also search and output matches to a regex template that defaults to match on email addresses, or can just simply print the structure of web.

## Installation Instructions
```bash
git clone https://github.com/teohrt/abrasion.git
cd abrasion/
make build
```

## Usage
```bash
# Crawl the web starting from google.com, the default seed URL, and output emails to a text file.
./abrasion -getEmail

# Scrape 100 URLs starting with given seed URL. Output all related URLs and debug logs to their respective files, as well as the console. 
./abrasion -url=https://reddit.com -scrapeLimit=100 -verbose -debug
```

## Flags
* url - The URL with which Abrasion begins scraping the web. Defaults to "https://www.google.com".
* verbose - Sets verbose logging in console. Defaults to false.
* debug - When set, a debug log file is written. Defaults to false.
* getEmail - Aggregate email addresses. Defaults to false.
* scrapeLimit - Sets the number of URLs to scrape. Defaults to MAXINT.

## Logs
Abrasion has 3 potential output streams. By default it outputs errors and results to their own respective text files and will also print to console if the -verbose flag is set.