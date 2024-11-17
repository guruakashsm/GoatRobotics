package constants

import (
	"github.com/rs/zerolog"
)

// Log Level Constants
const (
	Debug = "Debug"
	Info  = "Info"
	Warn  = "Warn"
	Error = "Error"
)

var LogMap = map[string]zerolog.Level{
	Debug: zerolog.DebugLevel,
	Info:  zerolog.InfoLevel,
	Warn:  zerolog.WarnLevel,
	Error: zerolog.ErrorLevel,
}

// Server Level Default Configuration
const (
	Port    = 8080
	Host    = "localhost"
	BaseURL = "/rpc/GOATROBOTICS/"
	Version = "v1.0.1"
)


// Messaging Configuration
const (
	MaxMessage = 100
)
