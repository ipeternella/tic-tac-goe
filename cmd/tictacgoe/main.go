// Main entry-point of the game.
package main

import (
	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/repl"
	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"

	"github.com/rs/zerolog/log"
)

// Main game loop
func main() {
	// sets the application up
	settings.SetupApplication()
	log.Debug().Msgf("Starting the game...")

	repl.StartGameLoop()
	log.Debug().Msgf("Closing the game application...")
}
