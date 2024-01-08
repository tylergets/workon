// logger/logger.go
package logger

import (
	"fmt"
	"github.com/fatih/color"
)

type Logger struct {
	StepID    string
	colorFunc func(format string, a ...interface{})
}

func NewLogger(stepID int) *Logger {
	logger := &Logger{StepID: fmt.Sprintf("Step[%d]", stepID)}

	// Assign a unique color based on the step ID
	switch stepID % 5 {
	case 0:
		logger.colorFunc = color.Cyan
	case 1:
		logger.colorFunc = color.Green
	case 2:
		logger.colorFunc = color.Yellow
	case 3:
		logger.colorFunc = color.Red
	case 4:
		logger.colorFunc = color.Magenta
	}

	return logger
}

func (l *Logger) log(format string, a ...interface{}) {
	l.colorFunc(fmt.Sprintf("%s: %s\n", l.StepID, format), a...)
}

func (l *Logger) Info(format string, a ...interface{}) {
	l.log(format, a...)
}

func (l *Logger) Success(format string, a ...interface{}) {
	l.log(format, a...)
}

func (l *Logger) Warn(format string, a ...interface{}) {
	l.log(format, a...)
}

func (l *Logger) Error(format string, a ...interface{}) {
	l.log(format, a...)
}
