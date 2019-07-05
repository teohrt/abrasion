package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(s string)
	FlushLogger()
}

type loggerimpl struct {
	verbose bool
	logger  *csv.Writer
}

func NewLogger(fileName string, verbose bool) Logger {
	return &loggerimpl{
		verbose: verbose,
		logger:  newCSVWriter(fileName),
	}
}

func (l *loggerimpl) Log(s string) {
	if l.verbose {
		fmt.Println(s)
	}

	if err := l.logger.Write([]string{s}); err != nil {
		log.Fatal("failed writing to file. Msg: " + s + "\n" + err.Error())
	}
}

func (l *loggerimpl) FlushLogger() {
	l.logger.Flush()
}

func newCSVWriter(filename string) *csv.Writer {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("failed creating file", err)
	}

	return csv.NewWriter(file)
}
