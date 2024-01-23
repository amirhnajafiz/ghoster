package logger

type Level int

const (
	ErrorLevel Level = iota + 1
	InfoLevel
	DebugLevel
)
