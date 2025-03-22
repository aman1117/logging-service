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

func (s *LoggingServer) GetTopLogsBasedOnLevelInLastNSeconds(N time.Duration) []types.Log {
	thresholdTime := time.Now().Add(-N)
	logCounts := map[types.LogLevel]int{}
	var filteredLogs []types.Log


	for _, log := range s.Logs {
		if log.Timestamp.After(thresholdTime) {
			filteredLogs = append(filteredLogs, log)
			logCounts[log.Level]++ 
		}
	}

	var targetLevel types.LogLevel
	maxCount := 0
	for level, count := range logCounts {
		if count > maxCount {
			maxCount = count
			targetLevel = level
		}
	}

	var resultLogs []types.Log
	for _, log := range filteredLogs {
		if log.Level == targetLevel {
			resultLogs = append(resultLogs, log)
		}
	}

	return resultLogs
}
