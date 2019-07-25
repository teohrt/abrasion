package app

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/teohrt/abrasion/utils"
)

type Config struct {
	SeedURL    string
	RegexValue string
	Verbose    bool

	Client       http.Client
	DataChan     chan string
	URLChan      chan string
	ErrorLogger  utils.Logger
	ResultLogger utils.Logger
}

func Start(c *Config) {
	currentTime := time.Now().Format("2006-01-02 3:4:5 pm")
	errorFileName := "Abrasion_Error_log_" + currentTime + ".csv"
	resultFileName := "Abrasion_Result_log_" + currentTime + ".csv"

	c.ErrorLogger = utils.NewLogger(errorFileName, c.Verbose)
	c.ResultLogger = utils.NewLogger(resultFileName, c.Verbose)
	c.DataChan = make(chan string)
	c.URLChan = make(chan string)

	timeout := time.Duration(5 * time.Second)
	c.Client = http.Client{
		Timeout: timeout,
	}

	if err := validate(c); err != nil {
		os.Exit(1)
	}

	go c.Process()
	c.Scrape(c.SeedURL)

	select {} // Block forever
}

func validate(c *Config) error {
	_, err := url.Parse(c.SeedURL)
	if err != nil {
		c.ErrorLogger.Log("Error parsing URL. : " + c.SeedURL)
	}
	return err
}
