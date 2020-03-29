// Main entry-point of the game with the game loop.
package main

import (
	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe"
	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"

	"github.com/rs/zerolog/log"
)

// Main game loop
func main() {
	settings.SetupApplication()
	log.Debug().Msg("Starting the game...")

	tictacgoe.Play()
	log.Debug().Msg("Closing the game application...")
}
