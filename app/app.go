package app

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"sync"
	"syscall"
	"time"

	"github.com/teohrt/abrasion/utils"
)

type Config struct {
	SeedURL     string
	ScrapeLimit int
	GetEmails   bool
	Verbose     bool
	Debug       bool

	Client   http.Client
	Wg       *sync.WaitGroup
	DataChan chan string
	URLChan  chan string
	Logger   utils.Logger
	Regex    *regexp.Regexp
}

func Start(c *Config) {
	initApp(c)

	if err := validate(c); err != nil {
		os.Exit(1)
	}

	go c.Process()
	go c.Scrape(c.SeedURL)

	c.Wg.Wait()
	cleanup(c)
}

func initApp(c *Config) {
	c.Wg = &sync.WaitGroup{}
	c.Wg.Add(c.ScrapeLimit)

	currentTime := time.Now().Format("2006-01-02_3:4:5_pm")
	errorFileName := "Abrasion_Error_log_" + currentTime + ".txt"
	resultFileName := "Abrasion_Result_log_" + currentTime + ".txt"
	logger, err := utils.NewLogger(resultFileName, errorFileName, c.Verbose, c.Debug)
	if err != nil {
		log.Fatal("Failed creating logger " + err.Error())
	}

	c.Logger = logger
	c.DataChan = make(chan string)
	c.URLChan = make(chan string)
	c.Client = http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	if c.GetEmails {
		c.Regex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	}

	handleKill(c)
}

func validate(c *Config) error {
	_, err := url.Parse(c.SeedURL)
	if err != nil {
		c.Logger.Err("Error parsing seed URL. : " + c.SeedURL)
	}
	return err
}

func cleanup(c *Config) {
	c.Logger.Flush()
	c.Logger.Close()
	fmt.Println("\nStopped scrapping.")
	os.Exit(1)
}

// Handles cleanup on interupt signal
func handleKill(c *Config) {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		cleanup(c)
	}()
}
