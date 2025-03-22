package types

import "time"

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
