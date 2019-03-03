package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Supply a url")
		return
	}

	url := os.Args[1]
	hasHTTP := strings.Index(url, "http") == 0
	if !hasHTTP {
		log.Println("URL must start with 'http://'")
		return
	}

	var URLs []string
	URLChan := make(chan string)
	finished := make(chan bool)

	go crawl(url, URLChan, finished)

	for i := 0; i < 1; {
		select {
		case newURL := <-URLChan:
			URLs = append(URLs, newURL)
		case <-finished:
			i++
		}
	}

	printURLs(URLs)
}

// Extract all links from a page and concurrently returns them over the linkChan
func crawl(url string, linkChan chan string, finished chan bool) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to crawl: " + url)
		finished <- false
		return
	}
	defer res.Body.Close()

	tokenizer := html.NewTokenizer(res.Body)

	for {
		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken:
			finished <- false
			return

		case tokenType == html.StartTagToken:
			token := tokenizer.Token()

			isATag := token.Data == "a"
			if !isATag {
				continue
			}

			newURL, ok := getHrefLink(token)

			if !ok {
				continue
			}

			hasHTTP := strings.Index(newURL, "http") == 0
			if hasHTTP {
				linkChan <- newURL
			}
		}
	}
}

func getHrefLink(t html.Token) (string, bool) {
	url := ""
	ok := false

	for _, a := range t.Attr {
		if a.Key == "href" {
			url = a.Val
			ok = true
		}
	}

	return url, ok
}

func printURLs(links []string) {
	fmt.Printf("%d links found!\n\n", len(links))
	for _, l := range links {
		fmt.Println(l)
	}
}
