package settings

import (
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Sets up game properties
func setupGame(boardSize int) error {

	// Sanity checks
	if boardSize < 3 || boardSize > 9 {
		return errors.New("the largest supported board is a 9 x 9 and the smaller is 3 x 3")
	}

	return nil
}

// Setups the logs of the application: level, format, etc.
func setupLogs(loggingLevel string) {
	level := loggingLevelMap[loggingLevel]
	zerolog.SetGlobalLevel(level)
}

// Setups the the application
func SetupApplication() {

	setupLogs(loggingLevel)
	log.Debug().Msg("Setting the application up...")

	// Setup may raise exceptions
	if err := setupGame(BoardSize); err != nil {
		log.Fatal().Msgf("Setup error: %s", err.Error())
		os.Exit(1)
	}

	log.Debug().Msg("Application has been successfully set up!")

}
