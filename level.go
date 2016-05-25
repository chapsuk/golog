package golog

// Levels constants
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

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
