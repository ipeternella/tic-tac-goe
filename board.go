// Package board contains the board struct used to hold the game's state. Moreover,
// functions that operate on the board are also available.
package board

import (
	"fmt"
	"strconv"
)

type Board struct {
	fields [3][3]string
}

/**
Returns a new Board.
*/
func CreateBoard() *Board {
	var emptyBoard [3][3]string
	k := 1

	for i, _ := range emptyBoard {

		for j, _ := range emptyBoard[i] {

			emptyBoard[i][j] = strconv.FormatInt(int64(k), 10)
			k++
		}
	}

	return &Board{fields: emptyBoard}
}

/**
Prints a board.
*/
func PrintBoard(board *Board) {

	for i, _ := range board.fields {

		fmt.Printf(" %s | %s | %s \n", board.fields[i][0], board.fields[i][1], board.fields[i][2])

	}
}
