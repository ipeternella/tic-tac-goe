// Prints the game state to the players
package repl

import (
	"fmt"
	"time"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/gamelogic"
	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"
)

// Writes to terminal screen buffer
func DisplayScreenMessage(message string, newline bool) {
	if newline {
		fmt.Println(message)
	}

	fmt.Println(message)
}

func DisplayBoardWithSpaces(board *gamelogic.Board) {
	time.Sleep(300 * time.Millisecond) // waits a little bit before writing board to screen buffer
	DisplayScreenMessage("\n", false)

	gamelogic.PrintBoard(board, settings.BoardHorizontalSeparator)
	DisplayScreenMessage("\n", false)
}
