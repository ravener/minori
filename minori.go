// Minori is a minimal logging package for Golang.
package minori

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"io"
	"os"
)

var LogLevel = 6

// SetLevel sets the global log level.
// Panics if level is invalid.
func SetLevel(level int) {
	if level > 6 || level < 0 {
		panic("Invalid Log Level.")
	}

	LogLevel = level
}

type Logger struct {
	Name  string
	Out   io.Writer
	Level int // -1 to use the global LogLevel
}

const (
	OFF   = 0
	FATAL = 1
	PANIC = 2
	ERROR = 3
	WARN  = 4
	INFO  = 5
	DEBUG = 6
)

func (l *Logger) log(level int, msg string) {
	logl := LogLevel
	if l.Level != -1 {
		logl = l.Level
	}
	if level > logl {
		return
	}

	// Align the messages nicely.
	ws := ""
	if level == INFO || level == WARN {
		ws = " "
	}

	fmt.Fprintf(l.Out, "\x1b[%dm[%s]%s\x1b[0m \x1b[35m[%s]\x1b[0m %s\n",
		getColorByLevel(level), getMessageByLevel(level), ws,
		l.Name, msg,
	)
}

func (l *Logger) Panic(v interface{}) {
	l.log(PANIC, fmt.Sprint(v))
	panic(v)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	p := fmt.Sprintf(format, v...)
	l.log(PANIC, p)
	panic(p)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.log(FATAL, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log(FATAL, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Info(v ...interface{}) {
	l.log(INFO, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.log(INFO, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.log(ERROR, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log(ERROR, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.log(DEBUG, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log(DEBUG, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.log(WARN, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log(WARN, fmt.Sprintf(format, v...))
}

func GetLogger(name string) *Logger {
	return &Logger{Name: name, Out: colorable.NewColorableStdout(), Level: -1}
}

func GetLoggerOutput(name string, writer io.Writer) *Logger {
	return &Logger{Name: name, Out: colorable.NewNonColorable(writer), Level: -1}
}

func GetLoggerLevel(name string, level int) *Logger {
	return &Logger{Name: name, Out: colorable.NewColorableStdout(), Level: level}
}

func GetLoggerLevelOutput(name string, level int, writer io.Writer) *Logger {
	return &Logger{Name: name, Out: colorable.NewNonColorable(writer), Level: level}
}

func getMessageByLevel(level int) string {
	switch level {
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case DEBUG:
		return "DEBUG"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	default:
		return ""
	}
}

func getColorByLevel(level int) int {
	switch level {
	case DEBUG:
		return 37
	case WARN:
		return 33
	case ERROR, FATAL, PANIC:
		return 31
	default:
		return 36
	}
}
