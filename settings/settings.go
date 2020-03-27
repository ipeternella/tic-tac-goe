// Package with configurations and settings variables used throughout the project.
package settings

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Game properties
const (
	BoardHorizontalSeparator string = "-"
	BoardSize                int    = 3
)

// Logging configs
var loggingLevelMap = map[string]zerolog.Level{
	"panic": zerolog.PanicLevel,
	"fatal": zerolog.FatalLevel,
	"warn":  zerolog.WarnLevel,
	"info":  zerolog.InfoLevel,
	"debug": zerolog.DebugLevel,
}

const loggingLevel string = "info"

// Sets up game properties
func setupGame() {}

// Setups the logs of the application: level, format, etc.
func setupLogs(loggingLevel string) {
	level := loggingLevelMap[loggingLevel]
	zerolog.SetGlobalLevel(level)
}

// Setups the the application
func SetupApplication() {
	setupLogs(loggingLevel)

	log.Debug().Msg("Booting the application...")
	setupGame()

	log.Debug().Msg("Application has been successfully set up!")
}
