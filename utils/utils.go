package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger interface {
	Log(s string)
	Err(s string)
	Close()
}

type loggerimpl struct {
	verbose    bool
	resultFile *os.File
	errorFile  *os.File
}

func NewLogger(resultFileName string, errorFileName string, verbose bool) (Logger, error) {
	r, err := createFile(resultFileName)
	if err != nil {
		return nil, err
	}
	e, err := createFile(errorFileName)
	if err != nil {
		return nil, err
	}

	return &loggerimpl{
		verbose:    verbose,
		resultFile: r,
		errorFile:  e,
	}, nil
}

func (l *loggerimpl) Log(s string) {
	if l.verbose {
		fmt.Println(s)
	}

	_, err := io.WriteString(l.resultFile, s+"\n")
	if err != nil {
		log.Fatal("failed writing to file. Msg: " + s + "\n" + err.Error())
	}
	err = l.resultFile.Sync()
	if err != nil {
		log.Fatal("failed syncing file. Msg: " + s + "\n" + err.Error())
	}
}

func (l *loggerimpl) Err(s string) {
	if l.verbose {
		fmt.Println(s)
	}

	_, err := io.WriteString(l.errorFile, s+"\n")
	if err != nil {
		log.Fatal("failed writing to file. Msg: " + s + "\n" + err.Error())
	}
	err = l.errorFile.Sync()
	if err != nil {
		log.Fatal("failed syncing file. Msg: " + s + "\n" + err.Error())
	}
}

func (l *loggerimpl) Close() {
	l.resultFile.Close()
	l.errorFile.Close()
}

func createFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}
