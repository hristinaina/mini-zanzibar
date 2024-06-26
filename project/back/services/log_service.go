package services

import (
	"log"
	"os"
	"sync"
)

type LogService struct {
	logger *log.Logger
	mu     sync.Mutex
}

func NewLogService(fileName string) (*LogService, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	logger := log.New(file, "[MyApp] ", log.Ldate|log.Ltime|log.Lshortfile)
	return &LogService{
		logger: logger,
	}, nil
}

func (ls *LogService) Info(message string) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.logger.Printf("[INFO] %s\n", message)
}

func (ls *LogService) Error(message string) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.logger.Printf("[ERROR] %s\n", message)
}

func (ls *LogService) Warning(message string) {
	ls.mu.Lock()
	defer ls.mu.Unlock()
	ls.logger.Printf("[WARNING] %s\n", message)
}
