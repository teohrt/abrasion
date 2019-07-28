package app

import (
	"strings"

	"golang.org/x/net/html"
)

// Extracts all links and emails from a page and sends them back to processor over channels
func (c *Config) Scrape(URL string) {
	res, err := c.Client.Get(URL)
	if err != nil {
		c.Logger.Err(err.Error())
		return
	}
	defer res.Body.Close()

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

			link, ok := getHrefLink(token)
			if !ok {
				continue
			}

			hasHTTP := strings.Index(link, "http") == 0
			if hasHTTP {
				c.URLChan <- link
			} else {
				if c.GetEmail {
					isEmail := strings.Index(link, "mailto:") == 0
					if isEmail {

						linkRunes := []rune(link)
						linkSubstring := linkRunes[7:] // remove "mailto:"
						email := string(linkSubstring)

						if c.Regex.MatchString(email) {
							c.DataChan <- email
						} else {
							c.Logger.Err("BAD EMAIL: " + email)
						}
					}
				}
			}
		}
	}
	c.Wg.Done()
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
