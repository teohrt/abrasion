package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

const MAX_BUFFER_SIZE int = 100

type Logger interface {
	Log(s string)
	Err(s string)
	Close() error
	Flush() error
}

type loggerimpl struct {
	verbose           bool
	debug             bool
	resultFile        *os.File
	errorFile         *os.File
	resultWriter      *bufio.Writer
	errorWriter       *bufio.Writer
	resultBufferCount int
	errorBufferCount  int
	mutex             *sync.Mutex
}

func NewLogger(resultFileName string, errorFileName string, verbose bool, debug bool) (Logger, error) {
	rf, err := createFile(resultFileName)
	if err != nil {
		return nil, err
	}
	rw := bufio.NewWriter(rf)

	ef := &os.File{}
	ew := &bufio.Writer{}
	if debug {
		ef, err = createFile(errorFileName)
		if err != nil {
			return nil, err
		}
		ew = bufio.NewWriter(ef)
	}

	return &loggerimpl{
		verbose:           verbose,
		debug:             debug,
		resultFile:        rf,
		resultWriter:      rw,
		errorFile:         ef,
		errorWriter:       ew,
		resultBufferCount: 0,
		errorBufferCount:  0,
		mutex:             &sync.Mutex{},
	}, nil
}

func (l *loggerimpl) Log(s string) {
	l.mutex.Lock()

	if l.verbose {
		fmt.Println(s)
	}

	fmt.Fprintln(l.resultWriter, s)
	l.resultBufferCount++

	if l.resultBufferCount >= MAX_BUFFER_SIZE {
		if err := l.resultWriter.Flush(); err != nil {
			log.Fatal(err.Error())
		}

		l.resultBufferCount = 0
	}

	l.mutex.Unlock()
}

func (l *loggerimpl) Err(s string) {
	l.mutex.Lock()

	if l.verbose && l.debug {
		fmt.Println(s)
	}

	if l.debug {
		fmt.Fprintln(l.errorWriter, s)
		l.errorBufferCount++

		if l.errorBufferCount >= MAX_BUFFER_SIZE {
			if err := l.errorWriter.Flush(); err != nil {
				log.Fatal(err.Error())
			}

			l.errorBufferCount = 0
		}
	}

	l.mutex.Unlock()
}

func (l *loggerimpl) Close() error {
	err := l.resultFile.Close()
	if err != nil {
		return err
	}

	if l.debug {
		err = l.errorFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *loggerimpl) Flush() error {
	err := l.resultWriter.Flush()
	if err != nil {
		return err
	}

	if l.debug {
		err = l.errorWriter.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func createFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}
