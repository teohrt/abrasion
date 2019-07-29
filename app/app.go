package app

import (
	"log"
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

	Client   http.Client
	Wg       *sync.WaitGroup
	DataChan chan string
	URLChan  chan string
	Logger   utils.Logger
	Regex    *regexp.Regexp
}

func Start(c *Config) {
	initApp(c)
	defer c.Logger.Close()

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

	currentTime := time.Now().Format("2006-01-02_3:4:5_pm")
	errorFileName := "Abrasion_Error_log_" + currentTime + ".txt"
	resultFileName := "Abrasion_Result_log_" + currentTime + ".txt"
	logger, err := utils.NewLogger(resultFileName, errorFileName, c.Verbose)
	if err != nil {
		log.Fatal("Failed creating logger " + err.Error())
	}

	c.Logger = logger
	c.DataChan = make(chan string)
	c.URLChan = make(chan string)
	c.Client = http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	if c.GetEmail {
		c.Regex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	}
}

func validate(c *Config) error {
	_, err := url.Parse(c.SeedURL)
	if err != nil {
		c.Logger.Err("Error parsing seed URL. : " + c.SeedURL)
	}
	return err
}
