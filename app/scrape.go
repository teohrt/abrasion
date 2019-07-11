package app

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Extracts all links from a page and concurrently returns them over the datachan
func (c *Config) Scrape(URL string) {
	res, err := http.Get(URL)
	if err != nil {
		c.ErrorLogger.Log(err.Error())
		return
	}
	defer res.Body.Close()

	if c.GetEmail {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.ErrorLogger.Log("failed reading response body: " + err.Error())
		} else {
			c.scrapeForEmails(string(bodyBytes))
			res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset the res io.Reader for re-use in the next function
		}
	}

	c.scrapeForURLs(res)
}

func (c *Config) scrapeForEmails(body string) {
	//fmt.Println(body)
	matches := c.Regex.FindAllString(body, -1)
	for _, v := range matches {
		c.DataChan <- v
	}
}

func (c *Config) scrapeForURLs(res *http.Response) {
	tokenizer := html.NewTokenizer(res.Body)

	loop := true
	for {
		if loop == false {
			break
		}

		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken:
			loop = false
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
				c.URLChan <- newURL
			}
		}
	}
}

func getHrefLink(t html.Token) (URL string, ok bool) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			URL = a.Val
			ok = true
			return
		}
	}
	return
}
