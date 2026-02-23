package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func GenerateNewLogger(prefix string) (*log.Logger, *os.File, error) {
	logDir := "logs"

	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create log directory: %w", err)
	}
	fileName := fmt.Sprintf("%s_%s.log", prefix, time.Now().Format("2006-01-02"))
	logFilePath := filepath.Join(logDir, fileName)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file %s: %w", logFilePath, err)
	}

	logger := log.New(file, fmt.Sprintf("[%s] ", prefix), log.Ldate|log.Ltime|log.Lshortfile)

	return logger, file, nil
}

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
