package config

import (
	"io"
	"log"
)

// Logs configuration
type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)
	return &Logger{
		debug: log.New(writer,"DEBUG: ", log.Flags()),
		info: log.New(writer,"INFO: ", log.Flags()),
        warning: log.New(writer,"WARNING: ", log.Flags()),
        err: log.New(writer,"ERROR: ", log.Flags()),
		writer:	writer,
}

// create non-formatted log message

func (l *Logger) Debug(v ...interface{}) {
    l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v...interface{}) {
    l.warning.Println(v...)
}	

func (l *Logger) Error(v...interface{}) {
    l.err.Println(v...)
}

