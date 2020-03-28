// Main entry-point of the game.
package main

import (
	"os"
	"strconv"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/repl"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/gamelogic"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"

	"github.com/rs/zerolog/log"
)

// Main game loop
func main() {
	// sets the application up
	settings.SetupApplication()

	log.Debug().Msgf("Starting the game...")

	gameBoard := gamelogic.CreateGameBoard(settings.BoardSize)
	gamelogic.PrintBoard(gameBoard, settings.BoardHorizontalSeparator)

	userInput := repl.GetUserInput(os.Stdin)
	userFieldPositionInput, intConversionError := strconv.Atoi(userInput)

	if intConversionError != nil {
		log.Fatal().Msg("Input was not an integer!")
	}

	if isValid, rejectionMsg := gamelogic.IsUserInputValid(userFieldPositionInput, gameBoard); !isValid {
		log.Fatal().Msg(rejectionMsg)
	}

	// updates game and re-prints the board
	gamelogic.UpdateGameState("IG", userFieldPositionInput, gameBoard)
	gamelogic.PrintBoard(gameBoard, settings.BoardHorizontalSeparator)

	log.Debug().Msgf("Closing the game application...")
}
