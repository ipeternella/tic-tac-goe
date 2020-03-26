// Package board contains the board struct used to hold the game's state. Moreover,
// functions that operate on the board are also available.
package board

import (
	"fmt"
	"strconv"
)

type Board struct {
	fields [][]string
}

// Creates a new game board of a given boardSize. Boards are square matrices.
func CreateGameBoard(boardSize int) *Board {
	k := 1

	// initializes slices
	emptyBoard := make([][]string, boardSize)
	for i := range emptyBoard {
		emptyBoard[i] = make([]string, boardSize)
	}

	// fills slice dynamically
	for i := range emptyBoard {
		for j := range emptyBoard[i] {
			emptyBoard[i][j] = strconv.FormatInt(int64(k), 10)
			k++
		}
	}

	return &Board{fields: emptyBoard}
}

// Prints the game board with its actual game state.
func PrintBoard(board *Board) {

	// traverses each row
	for i := range board.fields {
		fmt.Printf(" %s | %s | %s \n", board.fields[i][0], board.fields[i][1], board.fields[i][2])
	}
}
