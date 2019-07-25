package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(s string)
	Err(s string)
	Flush()
}

type loggerimpl struct {
	verbose      bool
	resultLogger *csv.Writer
	errorLogger  *csv.Writer
}

func NewLogger(resultFileName string, errorFileName string, verbose bool) Logger {
	return &loggerimpl{
		verbose:      verbose,
		resultLogger: newCSVWriter(resultFileName),
		errorLogger:  newCSVWriter(errorFileName),
	}
}

func (l *loggerimpl) Log(s string) {
	if l.verbose {
		fmt.Println(s)
	}

	if err := l.resultLogger.Write([]string{s}); err != nil {
		log.Fatal("failed writing to file. Msg: " + s + "\n" + err.Error())
	}
}

func (l *loggerimpl) Err(s string) {
	if l.verbose {
		fmt.Println(s)
	}

	if err := l.errorLogger.Write([]string{s}); err != nil {
		log.Fatal("failed writing to file. Msg: " + s + "\n" + err.Error())
	}
}

func (l *loggerimpl) Flush() {
	l.resultLogger.Flush()
	l.errorLogger.Flush()
}

func newCSVWriter(filename string) *csv.Writer {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("failed creating file", err)
	}

	return csv.NewWriter(file)
}
