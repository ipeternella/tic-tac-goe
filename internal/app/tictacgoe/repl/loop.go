// contains main game loop
package repl

import (
	"fmt"
	"strconv"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/gamelogic"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"
)

// Runs the game loop until the user wants to quit
func StartGameLoop() {

	AskUserWithoutOptions(settings.WelcomeMsg, false)
	AskUserWithoutOptions(fmt.Sprintf(settings.MatchAboutToStartMsg, settings.BoardSize, settings.BoardSize), false)
	DisplayScreenMessage(settings.InitialBoardMsg, false)
	gameBoard := gamelogic.CreateGameBoard(settings.BoardSize)

	// forever game loop
	for {
		DisplayBoardWithSpaces(gameBoard)
		rawUserInput := AskUserWithoutOptions(settings.AskNextMoveMsg, false) // TODO: validate inputs

		// checks if user wants to quit
		if rawUserInput == settings.QuitMark {
			break
		}

		userFieldPositionInput, _ := strconv.Atoi(rawUserInput)
		isValidMove, rejectionMsg := gamelogic.IsUserInputValid(userFieldPositionInput, gameBoard)

		if isValidMove {
			gamelogic.UpdateGameState(settings.Player1Mark, userFieldPositionInput, gameBoard) // TODO: asserts game over
		} else {
			DisplayScreenMessage(rejectionMsg, true)
		}

	}

	DisplayScreenMessage(settings.ByeByeMsg, true)
}
