package logs

import (
	"io"
	"log"
	"os"
)

func initLoggers(logHandle io.Writer) (*log.Logger, *log.Logger, *log.Logger) {
	Info := log.New(logHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Warning := log.New(logHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error := log.New(logHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	return Info, Warning, Error
}

func Open(path string) (*log.Logger, *log.Logger, *log.Logger, error) {
	logHandle, err := os.OpenFile(path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		log.Printf("Failed to open log file @ %s: %s\n", path, err)
		return nil, nil, nil, err
	}
	infoL, warnL, errL := initLoggers(logHandle)
	return infoL, warnL, errL, nil
}
