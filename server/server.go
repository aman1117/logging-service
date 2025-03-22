package server

import (
	"time"

	"github.com/aman1117/logging-service/types"
)

type LoggingServer struct {
	Logs []types.Log
}

func (s *LoggingServer) Push(logs []types.Log) {
	s.Logs = append(s.Logs, logs...)
}

func NewLoggingServer() *LoggingServer {
	return &LoggingServer{
		Logs: []types.Log{},
	}
}

func (s *LoggingServer) FilterBasedOnLevel(level types.LogLevel) []types.Log {
	var filteredLogs []types.Log
	for _, log := range s.Logs {
		if log.Level == level {
			filteredLogs = append(filteredLogs, log)
		}
	}
	return filteredLogs
}

func (s *LoggingServer) GetTopLogsInLastNSeconds(d time.Duration) []types.Log {
	var filteredLogs []types.Log
	thresholdTime := time.Now().Add(-d) // Compute once

	for _, log := range s.Logs {
		if log.Timestamp.After(thresholdTime) {
			filteredLogs = append(filteredLogs, log)
		}
	}
	return filteredLogs
}

