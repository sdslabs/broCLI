package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/shiena/ansicolor"
)

// Colors
const (
	Red    = "\x1b[31m"
	Green  = "\x1b[32m"
	Yellow = "\x1b[33m"
	Blue   = "\x1b[34m"
	Reset  = "\x1b[0m"
)

// Logger type
type Logger struct {
	writer io.Writer
	prefix string
}

// NewLogger creates a new logger
func NewLogger(w io.Writer) *Logger {
	return &Logger{
		writer: w,
		prefix: fmt.Sprintf("%s>%s%s>%s%s>%s ", Red, Reset, Green, Reset, Blue, Reset),
	}
}

// Write writes to the logger writer
func (l *Logger) Write(str string) error {
	w := ansicolor.NewAnsiColorWriter(l.writer)
	_, err := fmt.Fprintf(w, str)
	return err
}

// Printf prints
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Write(fmt.Sprintf(format, v...))
}

// Print prints
func (l *Logger) Print(v ...interface{}) {
	l.Write(fmt.Sprint(v...))
}

// Println prints
func (l *Logger) Println(v ...interface{}) {
	l.Write(fmt.Sprintln(v...))
}

// Logf logs
func (l *Logger) Logf(format string, v ...interface{}) {
	l.Write(l.prefix + fmt.Sprintf(format, v...) + "\n")
}

// Log logs
func (l *Logger) Log(v ...interface{}) {
	l.Write(l.prefix + fmt.Sprint(v...) + "\n")
}

// Debugf logs debug info
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Write(l.prefix + Green + fmt.Sprintf(format, v...) + Reset + "\n")
}

// Debug logs debug info
func (l *Logger) Debug(v ...interface{}) {
	l.Write(l.prefix + Green + fmt.Sprint(v...) + Reset + "\n")
}

// Infof logs info
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Write(l.prefix + Blue + fmt.Sprintf(format, v...) + Reset + "\n")
}

// Info logs info
func (l *Logger) Info(v ...interface{}) {
	l.Write(l.prefix + Blue + fmt.Sprint(v...) + Reset + "\n")
}

// Warnf warns
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Write(l.prefix + Yellow + fmt.Sprintf(format, v...) + Reset + "\n")
}

// Warn warns
func (l *Logger) Warn(v ...interface{}) {
	l.Write(l.prefix + Yellow + fmt.Sprint(v...) + Reset + "\n")
}

// Errorf logs error
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Write(l.prefix + Red + fmt.Sprintf(format, v...) + Reset + "\n")
}

// Error logs error
func (l *Logger) Error(v ...interface{}) {
	l.Write(l.prefix + Red + fmt.Sprint(v...) + Reset + "\n")
}

// Fatalf is same as Errorf followed by os.Exit(1)
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Errorf(format, v...)
	os.Exit(1)
}

// Fatal is same as Error followed by os.Exit(1)
func (l *Logger) Fatal(v ...interface{}) {
	l.Error(v...)
	os.Exit(1)
}

// Panicf is same as Errorf followed by panic()
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Errorf(format, v...)
	panic(fmt.Errorf(format, v...))
}

// Panic is same as Error followed by panic()
func (l *Logger) Panic(v ...interface{}) {
	l.Error(v...)
	panic(fmt.Errorf("%s", v...))
}
