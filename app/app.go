package app

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/teohrt/abrasion/utils"
)

type Config struct {
	Site         string
	RegexValue   string
	Verbose      bool
	DataChan     chan string
	URLChan      chan string
	ErrorLogger  utils.Logger
	ResultLogger utils.Logger
}

func Start(c *Config) {
	currentTime := time.Now().Format("2006-01-02 3:4:5 pm")
	errorFileName := "Abrasion_Error_log_" + currentTime + ".csv"
	resultFileName := "Abrasion_Result_" + currentTime + ".csv"

	c.ErrorLogger = utils.NewLogger(errorFileName, c.Verbose)
	c.ResultLogger = utils.NewLogger(resultFileName, c.Verbose)

	validate(c)

	go c.Process()
	c.Scrape(c.Site)
}

func validate(c *Config) {
	hasHTTP := strings.Index(c.Site, "http") == 0
	if !hasHTTP {
		fmt.Println("URL must start with 'http://'")
		os.Exit(1)
	}
}
