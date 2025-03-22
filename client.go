package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type LoggingClient struct {
	BatchedLogs []Log
	BatchSize   int
	server      *LoggingServer
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