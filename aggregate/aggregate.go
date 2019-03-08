package aggregate

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	s "../scrape"
)

// Aggregate scrapes new URLs and logs them
func Aggregate(dataChan chan string) {
	w := newCSVWriter("result.csv")
	defer w.Flush()

	fmt.Println("Abrasion in progress...")
	for {
		select {
		case newURL := <-dataChan:
			go s.Scrape(newURL, dataChan)

			err := w.Write([]string{newURL})

			if err != nil {
				fmt.Println("Cannot write URL to file. : ", newURL)
			}
		}
	}
}

func newCSVWriter(filename string) *csv.Writer {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Error with csv.", err)
	}

	return csv.NewWriter(file)
}
