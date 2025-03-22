package main

import "time"

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