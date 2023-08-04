package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLoger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)
	return &Logger{
		debug:   log.New(writer, "[DEBUG] -> ", logger.Flags()),
		info:    log.New(writer, "[INFO] -> ", logger.Flags()),
		warning: log.New(writer, "[WARNING] -> ", logger.Flags()),
		err:     log.New(writer, "[ERROR] -> ", logger.Flags()),
		writer:  writer,
	}
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// No formated logs

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Print(string(colorGreen))
	l.debug.Println(v...)
	l.debug.Println(string(colorWhite))
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(string(colorBlue))
	l.info.Println(v...)
	l.info.Println(string(colorWhite))
}
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(string(colorYellow))
	l.warning.Println(v...)
	l.warning.Println(string(colorWhite))
}
func (l *Logger) Err(v ...interface{}) {
	l.err.Println(string(colorRed))
	l.err.Println(v...)
	l.err.Println(string(colorWhite))
}

// Formated logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.err.Println(string(colorRed))
	l.err.Printf(format, v...)
	l.err.Println(string(colorWhite))
}

func (l *Logger) InfoF(format string, v ...interface{}) {
	l.info.Println(string(colorBlue))
	l.info.Printf(format, v...)
	l.info.Println(string(colorWhite))
}
func (l *Logger) WarnF(format string, v ...interface{}) {
	l.warning.Println(string(colorYellow))
	l.warning.Printf(format, v...)
	l.warning.Println(string(colorWhite))
}
func (l *Logger) ErrF(format string, v ...interface{}) {
	l.err.Println(string(colorRed))
	l.err.Printf(format, v...)
	l.err.Println(string(colorWhite))
}