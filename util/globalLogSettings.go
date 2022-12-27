package util

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetGlobalLogSettings sets the global log settings for the application.
// logType: console or json
// logLevel: debug, info, warn, error
func SetGlobalLogSettings(logType, logLevel string) {
	// check if logType is valid
	if logType != "console" && logType != "json" {
		fmt.Printf("invalid log type: %s\n", logType)
		os.Exit(1)
	}

	// set log Writer (JSON or Console)
	if logType == "console" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// set log level
	switch logLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		// invalid log level
		fmt.Printf("invalid log level: %s\n", logLevel)
		// exit with error code 1
		os.Exit(1)
	}
}
