// Main entry-point of the game.
package main

import (
	board "github.com/IgooorGP/tic-tac-goe"
	settings "github.com/IgooorGP/tic-tac-goe"
)

// Main game loop
func main() {
	gameBoard := board.CreateGameBoard(settings.BoardSize)
	board.PrintBoard(gameBoard)
}
