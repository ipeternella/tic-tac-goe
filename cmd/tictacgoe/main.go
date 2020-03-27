// Main entry-point of the game.
package main

import (
	tictacgoe "github.com/IgooorGP/tic-tac-goe"
	"github.com/IgooorGP/tic-tac-goe/settings"
	"github.com/rs/zerolog/log"
)

// Main game loop
func main() {
	// sets the application up
	settings.SetupApplication()

	log.Info().Msgf("Starting the game...")

	gameBoard := tictacgoe.CreateGameBoard(settings.BoardSize)
	tictacgoe.PrintBoard(gameBoard)

	log.Info().Msgf("Closing the game application...")
}
