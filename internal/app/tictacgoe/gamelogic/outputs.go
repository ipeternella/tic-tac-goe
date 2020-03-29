// Prints the game state to the players
package gamelogic

import (
	"fmt"
	"time"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"
)

// Writes to terminal screen buffer
func DisplayScreenMessage(message string, newline bool) string {
	messageForStdOut := message

	if newline {
		messageForStdOut += "\n"
	}

	fmt.Print(messageForStdOut)
	return messageForStdOut
}

func DisplayBoardWithSpaces(board *Board) {
	time.Sleep(300 * time.Millisecond) // waits a little bit before writing board to screen buffer
	DisplayScreenMessage("\n", false)

	PrintBoard(board, settings.BoardHorizontalSeparator)
	DisplayScreenMessage("\n", false)
}
