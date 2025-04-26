package logger

import (
	"fmt"
)

type (
	LogInterface interface {
		Info()
		Error()
		Warning()
	}
)

type ErrorLog struct {
	Message string
}

type InfoLog struct {
	Message string
}

var adapter LogInterface

func (e ErrorLog) Error() string {
	return fmt.Sprintf("Error: %v", e.Message)
}

func (e ErrorLog) Info() string {
	return fmt.Sprintf("Error: %v", e.Message)
}

func logFile() {
	// Log the error message
}
