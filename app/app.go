package app

import (
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/teohrt/abrasion/utils"
)

type Config struct {
	SeedURL     string
	ScrapeLimit int
	GetEmail    bool
	Verbose     bool

	Client       http.Client
	Wg           *sync.WaitGroup
	DataChan     chan string
	URLChan      chan string
	ErrorLogger  utils.Logger
	ResultLogger utils.Logger
	Regex        *regexp.Regexp
}

func Start(c *Config) {
	initApp(c)

	if err := validate(c); err != nil {
		os.Exit(1)
	}

	go c.Process()
	go c.Scrape(c.SeedURL)

	c.Wg.Wait()
}

func initApp(c *Config) {
	c.Wg = &sync.WaitGroup{}
	c.Wg.Add(c.ScrapeLimit)

	currentTime := time.Now().Format("2006-01-02 3:4:5 pm")
	errorFileName := "Abrasion_Error_log_" + currentTime + ".csv"
	resultFileName := "Abrasion_Result_log_" + currentTime + ".csv"

	c.ErrorLogger = utils.NewLogger(errorFileName, c.Verbose)
	c.ResultLogger = utils.NewLogger(resultFileName, c.Verbose)
	c.DataChan = make(chan string)
	c.URLChan = make(chan string)

	c.Client = http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	if c.GetEmail {
		c.Regex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	}
}

func validate(c *Config) error {
	_, err := url.Parse(c.SeedURL)
	if err != nil {
		c.ErrorLogger.Log("Error parsing seed URL. : " + c.SeedURL)
	}
	return err
}
