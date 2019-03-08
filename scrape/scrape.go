package scrape

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Scrape extracts all links from a page and concurrently returns them over the datachan
func Scrape(URL string, dataChan chan string) {

	res, err := http.Get(URL)

	if err != nil {
		fmt.Println("Failed to crawl: " + URL)
		return
	}
	defer res.Body.Close()

	tokenizer := html.NewTokenizer(res.Body)

	for {
		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken:
			break

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
				dataChan <- newURL
			}
		}
	}
}

func getHrefLink(t html.Token) (string, bool) {
	URL := ""
	ok := false

	for _, a := range t.Attr {
		if a.Key == "href" {
			URL = a.Val
			ok = true
		}
	}

	return URL, ok
}
