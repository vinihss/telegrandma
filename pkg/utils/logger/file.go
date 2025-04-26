package logger

import (
	"log"
	"os"
	"path/filepath"
)

const defaultLogPath = "/var/log/drstein"

func start() {
	// Ensure the log directory exists
	err := os.MkdirAll(defaultLogPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Open the log file
	logFile, err := os.OpenFile(filepath.Join(defaultLogPath, "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Set log output to the file
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func init() {

}
