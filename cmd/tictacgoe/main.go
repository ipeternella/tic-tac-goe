package main

import (
	board "github.com/IgooorGP/tic-tac-goe"
)

func main() {
	gameBoard := board.CreateBoard()
	board.PrintBoard(gameBoard)
}
