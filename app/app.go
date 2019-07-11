package app

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/teohrt/abrasion/utils"
)

type Config struct {
	Site         string
	GetEmail     bool
	Verbose      bool
	DataChan     chan string
	URLChan      chan string
	ErrorLogger  utils.Logger
	ResultLogger utils.Logger
	Regex        *regexp.Regexp
}

func Start(c *Config) {
	currentTime := time.Now().Format("2006-01-02 3:4:5 pm")
	errorFileName := "Abrasion_Error_log_" + currentTime + ".csv"
	resultFileName := "Abrasion_Result_" + currentTime + ".csv"

	c.ErrorLogger = utils.NewLogger(errorFileName, c.Verbose)
	c.ResultLogger = utils.NewLogger(resultFileName, c.Verbose)
	c.DataChan = make(chan string)
	c.URLChan = make(chan string)

	if c.GetEmail {
		exp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		c.Regex = exp
	}

	validate(c)

	go c.Process()
	c.Scrape(c.Site)

	select {} // Block forever
}

func validate(c *Config) {
	hasHTTP := strings.Index(c.Site, "http") == 0
	if !hasHTTP {
		fmt.Println("URL must start with 'http://'")
		os.Exit(1)
	}
}
