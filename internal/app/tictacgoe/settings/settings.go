// Package with configurations and settings variables used throughout the project.
package settings

import (
	"github.com/rs/zerolog"
)

// Game properties
const (
	BoardHorizontalSeparator string = "-"
	BoardSize                int    = 3

	Player1Mark string = "XX"
	Player2Mark string = "OO"
	QuitMark    string = "q"
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
