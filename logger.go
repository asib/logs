package logs

import (
	"io"
	"log"
	"os"
	"time"
)

type Logger struct {
	Timeout time.Duration
	Last    time.Time
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func (l *Logger) InfoPrintln(v ...interface{}) {
	l.Info.Println(v...)
}

func (l *Logger) InfoPrintf(format string, v ...interface{}) {
	l.Info.Printf(format, v...)
}

func (l *Logger) InfoFatalln(v ...interface{}) {
	l.Info.Fatalln(v...)
}

func (l *Logger) InfoFatalf(format string, v ...interface{}) {
	l.Info.Fatalf(format, v...)
}

func (l *Logger) WarningPrintln(v ...interface{}) {
	if time.Since(l.Last) >= l.Timeout {
		l.Warning.Println(v...)
		l.Last = time.Now()
	}
}

func (l *Logger) WarningPrintf(format string, v ...interface{}) {
	if time.Since(l.Last) >= l.Timeout {
		l.Warning.Printf(format, v...)
		l.Last = time.Now()
	}
}

func (l *Logger) WarningFatalln(v ...interface{}) {
	l.Warning.Fatalln(v...)
}

func (l *Logger) WarningFatalf(format string, v ...interface{}) {
	l.Warning.Fatalf(format, v...)
}

func (l *Logger) ErrorPrintln(v ...interface{}) {
	if time.Since(l.Last) >= l.Timeout {
		l.Error.Println(v...)
		l.Last = time.Now()
	}
}

func (l *Logger) ErrorPrintf(format string, v ...interface{}) {
	if time.Since(l.Last) >= l.Timeout {
		l.Error.Printf(format, v...)
		l.Last = time.Now()
	}
}

func (l *Logger) ErrorFatalln(v ...interface{}) {
	l.Error.Fatalln(v...)
}

func (l *Logger) ErrorFatalf(format string, v ...interface{}) {
	l.Error.Fatalf(format, v...)
}

func initLogger(logHandle io.Writer,
	timeout time.Duration) *Logger {
	return &Logger{
		timeout,
		time.Now().Add(-timeout),
		log.New(logHandle,
			"INFO: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		log.New(logHandle,
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile),
		log.New(logHandle,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func NewLogger(path string, timeout time.Duration) (*Logger, error) {
	logHandle, err := os.OpenFile(path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		log.Printf("Failed to open log file @ %s: %s\n", path, err)
		return nil, err
	}
	logger := initLogger(logHandle, timeout)
	return logger, nil
}
