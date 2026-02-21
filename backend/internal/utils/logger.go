package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// GenerateNewLogger creates a new logger instance that writes to a dynamically
// created file (e.g., inside a "logs" directory).
// Returns the logger, the file pointer (so it can be closed with defer file.Close()), and an error if any.
func GenerateNewLogger(prefix string) (*log.Logger, *os.File, error) {
	logDir := "logs"

	// Ensure the logs directory exists
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Create a log file name with the current date, e.g., "app_2026-02-21.log"
	fileName := fmt.Sprintf("%s_%s.log", prefix, time.Now().Format("2006-01-02"))
	logFilePath := filepath.Join(logDir, fileName)

	// Open the file in append mode, or create it if it doesn't exist
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file %s: %w", logFilePath, err)
	}

	// Create the logger
	logger := log.New(file, fmt.Sprintf("[%s] ", prefix), log.Ldate|log.Ltime|log.Lshortfile)

	return logger, file, nil
}

// GenerateGinLoggerFile creates a log file specifically meant to be used with Gin router.
func GenerateGinLoggerFile() (*os.File, error) {
	logDir := "logs"

	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	fileName := fmt.Sprintf("gin_%s.log", time.Now().Format("2006-01-02"))
	logFilePath := filepath.Join(logDir, fileName)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file %s: %w", logFilePath, err)
	}

	return file, nil
}
