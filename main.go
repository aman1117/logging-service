package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

type Log struct {
	Message   string
	Level     LogLevel
	Timestamp time.Time
}

type LoggingClient struct {
	BatchedLogs []Log
	BatchSize   int
	server      *LoggingServer
}

type LoggingServer struct {
	Logs []Log
}

func (s *LoggingServer) Push(logs []Log) {
	s.Logs = append(s.Logs, logs...)
}

func (s *LoggingServer) FilterBasedOnLevel(level LogLevel) []Log {
	var filteredLogs []Log
	for _, log := range s.Logs {
		if log.Level == level {
			filteredLogs = append(filteredLogs, log)
		}
	}
	return filteredLogs
}

func (c *LoggingClient) validate(message string, level LogLevel) error {
	if message == "" {
		return errors.New("message is empty")
	}
	if level < INFO || level > ERROR {
		return errors.New("invalid log level")
	}
	return nil
}

func (c *LoggingClient) Log(message string, level LogLevel) {
	err := c.validate(message, level)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	c.BatchedLogs = append(c.BatchedLogs, Log{
		Message:   message,
		Level:     level,
		Timestamp: time.Now(),
	})
	if len(c.BatchedLogs) == c.BatchSize {
		c.Flush()
	}
}

func (c *LoggingClient) Flush() {
	c.server.Push(c.BatchedLogs)
	c.BatchedLogs = []Log{}
}

func main() {
	server := &LoggingServer{}
	client := &LoggingClient{
		BatchSize: 3,
		server:    server,
	}

	// System startup logs
	client.Log("Logging service initialized successfully", INFO)
	client.Log("Database connection established", INFO)
	
	// Warning logs for potential issues
	client.Log("High memory usage detected: 85% of available memory", WARN)
	client.Log("API response time increased by 200ms", WARN)
	
	// Error logs for critical issues
	client.Log("Failed to connect to database: connection timeout", ERROR)
	client.Log("Payment processing failed: insufficient funds", ERROR)
	
	// More info logs
	client.Log("User authentication successful", INFO)
	client.Log("Cache cleared successfully", INFO)

	infoLogs := server.FilterBasedOnLevel(INFO)
	warnLogs := server.FilterBasedOnLevel(WARN)
	errorLogs := server.FilterBasedOnLevel(ERROR)
	
	fmt.Printf("Info logs: %v\n", infoLogs)
	fmt.Printf("Warn logs: %v\n", warnLogs)
	fmt.Printf("Error logs: %v\n", errorLogs)
	
}
