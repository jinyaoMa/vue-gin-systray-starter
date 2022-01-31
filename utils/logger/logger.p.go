package logger

import (
	"io"
	"log"
	"os"
)

func newLog(out io.Writer, prefix string) *log.Logger {
	return log.New(out, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}

func getLogFile(path string) (file *os.File) {
	var err error
	file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log: " + path)
	}
	return
}
