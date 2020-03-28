// Package board contains the board struct used to hold the game's state. Moreover,
// functions that operate on the board are also available.
package gamelogic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type Board struct {
	fields [][]string
}

// Creates a new game board of a given boardSize. Boards are square matrices.
func CreateGameBoard(boardSize int) *Board {
	log.Debug().Msgf("Creating a new game board with %d rows and cols...", boardSize)

	k := 0

	// initializes slices
	emptyBoard := make([][]string, boardSize)
	for i := range emptyBoard {
		emptyBoard[i] = make([]string, boardSize)
	}

	// fills slice dynamically
	for i := range emptyBoard {
		for j := range emptyBoard[i] {

			// single digit numbers require especial treatment
			if k <= 9 {
				emptyBoard[i][j] = "0" + strconv.FormatInt(int64(k), 10)
			} else {
				emptyBoard[i][j] = strconv.FormatInt(int64(k), 10)
			}
			k++
		}
	}

	log.Debug().Msgf("Initialized board. Returning board...")
	return &Board{fields: emptyBoard}
}

// internal function used to print rows
func getRowString(board *Board, rowNumber int) string {
	var row = board.fields[rowNumber]
	rowForStdOut := " "

	for i := range row {

		// traverses row EXCEPT last element
		if i <= len(row)-2 {
			rowForStdOut += row[i] + " | "
		} else {
			rowForStdOut += row[i] + " "
		}

	}

	return rowForStdOut
}

// Prints the game board with its actual game state
func PrintBoard(board *Board, boardHorizontalSeparator string) {
	totalHorizontalSeparators := 5*len(board.fields) - 1 // 4r + (r - 1) = 5r - 1
	horizontalBars := strings.Repeat(boardHorizontalSeparator, totalHorizontalSeparators)

	// traverses each row
	for rowNumber := range board.fields {
		// prints each row of the board
		fmt.Println(getRowString(board, rowNumber))

		// prints the horizontal line EXCEPT for the last row
		if rowNumber <= len(board.fields)-2 {
			fmt.Println(horizontalBars)
		}
	}
}

// Converts a validated field position, e.g. 1, to a (row, col) for 2d matrices notation
// Example: 1 --> (0, 0)
func GetBoardRowAndCol(validatedFieldPosition int, board *Board) (int, int) {
	totalRows := len(board.fields)
	totalCols := len(board.fields)

	row := validatedFieldPosition / totalRows // truncates decimals (integer division)
	col := validatedFieldPosition % totalCols

	return row, col
}

// Updates the game state with a new player's mark at a given board position
func UpdateGameState(playerMark string, validatedFieldPosition int, board *Board) {
	row, col := GetBoardRowAndCol(validatedFieldPosition, board)

	board.fields[row][col] = playerMark
}
