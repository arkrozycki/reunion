package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"

	conf "github.com/arkrozycki/reunion/config"
)

var log zerolog.Logger

// init function
func init() {
	conf.Config()
	switch os.Getenv("LogLevel") {
	case "dev":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen}
		log = zerolog.New(output).With().Timestamp().Caller().Logger()
	case "test":
		// zerolog.SetGlobalLevel(zerolog.PanicLevel)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Debug().Msgf("logger starting with %s level", os.Getenv("LogLevel"))
}

// Get function
func Get() zerolog.Logger {
	return log
}
