package client

import (
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/aman1117/logging-service/server"
	"github.com/aman1117/logging-service/types"
)

type LoggingClient struct {
	BatchedLogs []types.Log
	BatchSize   int
	server      *server.LoggingServer
}

func NewLoggingClient(server *server.LoggingServer, batchSize int) *LoggingClient {
	return &LoggingClient{
		server:      server,
		BatchedLogs: []types.Log{},
		BatchSize:   batchSize,
	}
}

func (c *LoggingClient) validate(message string, level types.LogLevel) error {
	if message == "" {
		return errors.New("message is empty")
	}
	if level < types.INFO || level > types.ERROR {
		return errors.New("invalid log level")
	}
	return nil
}

func (c *LoggingClient) Log(message string, level types.LogLevel) {
	err := c.validate(message, level)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	c.BatchedLogs = append(c.BatchedLogs, types.Log{
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
	c.BatchedLogs = []types.Log{}
}