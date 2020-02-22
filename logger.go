package ossa

import (
	fmt "fmt"
	"log"
)

// Logger logs messages
type Logger interface {
	LogDebugf(txid string, format string, a ...interface{})
	LogErrorf(txid string, format string, a ...interface{})
}

type noLogger struct{}

func (l *noLogger) LogErrorf(txid, format string, a ...interface{}) {}
func (l *noLogger) LogDebugf(txid, format string, a ...interface{}) {}

// StdOutLogger logs to stdout
type StdOutLogger struct{}

// LogErrorf logs at error level
func (l *StdOutLogger) LogErrorf(txid, format string, a ...interface{}) {
	log.Printf("TXID: %s ERROR: %s\n", txid, fmt.Sprintf(format, a...))
}

// LogDebugf logs at debug level
func (l *StdOutLogger) LogDebugf(txid, format string, a ...interface{}) {
	log.Printf("TXID: %s DEBUG: %s\n", txid, fmt.Sprintf(format, a...))
}
