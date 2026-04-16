package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Config holds logger configuration
type Config struct {
	Level  string
	Format string
}

// Init initializes the global logger
func Init(cfg Config) error {
	// Set log level
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	// Set output format
	if cfg.Format == "console" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		// JSON format (default)
		log.Logger = log.Output(os.Stderr)
	}

	return nil
}

// GetLogger returns the global logger instance
func GetLogger() zerolog.Logger {
	return log.Logger
}
