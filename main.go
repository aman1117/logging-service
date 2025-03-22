package main

import (
	"fmt"
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
}

func (c *LoggingClient) Log(message string, level LogLevel) {
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
	for _, log := range c.BatchedLogs {
		fmt.Printf("Timestamp: %s Level: %d Message: %s\n", log.Timestamp.Format("2006-01-02 15:04:05"), log.Level, log.Message)
	}
	c.BatchedLogs = []Log{}
}

func main() {
	client := &LoggingClient{
		BatchSize: 3,
	}
	
	fmt.Printf("Type of client: %T\n", client)  // Will print: *main.LoggingClient
	fmt.Printf("Value of client: %v\n", client) // Will print: &{[] 3}

	client.Log("Hello, World!", INFO)
	client.Log("Hello, World!", WARN)
	client.Log("Hello, World!", ERROR)
}
