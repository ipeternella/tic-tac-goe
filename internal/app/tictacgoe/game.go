// contains main game loop
package tictacgoe

import (
	"fmt"
	"strconv"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/gamelogic"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"
)

// Runs the game loop until the user wants to quit
func Play() {

	gamelogic.AskUserWithoutOptions(settings.WelcomeMsg, false)
	gamelogic.AskUserWithoutOptions(fmt.Sprintf(settings.MatchAboutToStartMsg, settings.BoardSize, settings.BoardSize), false)
	gamelogic.DisplayScreenMessage(settings.InitialBoardMsg, true)
	gameBoard := gamelogic.CreateGameBoard(settings.BoardSize)

	// forever game loop
	for {

		gamelogic.DisplayBoardWithSpaces(gameBoard)
		rawUserInput := gamelogic.AskUserWithoutOptions(settings.AskNextMoveMsg, false)

		// checks if user wants to quit
		if rawUserInput == settings.QuitMark {
			break
		}

		userFieldPositionInput, _ := strconv.Atoi(rawUserInput)
		isValidMove, rejectionMsg := gamelogic.IsUserInputValid(userFieldPositionInput, gameBoard)

		// if the move is valid, update the board
		if isValidMove {

			// player's move
			gamelogic.UpdateGameState(settings.Player1Mark, userFieldPositionInput, gameBoard)
			if gamelogic.IsGameOver(settings.Player1Mark, userFieldPositionInput, gameBoard) {
				break
			}

			// computer's move
			computerFieldPosition := gamelogic.CalculateComputerMove(gameBoard)
			gamelogic.UpdateGameState(settings.Player2Mark, computerFieldPosition, gameBoard)
			if gamelogic.IsGameOver(settings.Player2Mark, computerFieldPosition, gameBoard) {
				break
			}

		} else {
			gamelogic.DisplayScreenMessage(rejectionMsg, true)
		}

	}

	gamelogic.DisplayScreenMessage(settings.ByeByeMsg, true)
}
