package golog

import (
	"strings"
)

// Levels constants
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

// StringToLevel return Level by level name
func StringToLevel(lvl string) Level {
	lvl = strings.ToLower(lvl)
	switch lvl {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn", "warning":
		return WarnLevel
	case "error", "err":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	case "panic":
		return PanicLevel
	}
	return DebugLevel
}

// LevelToString convert LogLevel to a string
func LevelToString(level Level) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	case PanicLevel:
		return "PANIC"
	}

	return "UNKNOWN"
}
