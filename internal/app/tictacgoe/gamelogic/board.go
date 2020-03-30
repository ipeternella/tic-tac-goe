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
	freeFields []int
	fields     [][]string
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

	// free spots
	freeFields := make([]int, 0)
	for i := 0; i < boardSize*boardSize; i++ {
		freeFields = append(freeFields, i)
	}

	log.Debug().Msgf("Initialized board. Returning board...")
	return &Board{fields: emptyBoard, freeFields: freeFields}
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

// Retrieves a value from the board given a field position. Field value must have
func GetFieldValue(validatedFieldPosition int, board *Board) string {
	row, col := GetBoardRowAndCol(validatedFieldPosition, board)

	return board.fields[row][col]
}

// Updates the game state with a new player's mark at a given board position
func UpdateGameState(playerMark string, validatedFieldPosition int, board *Board) {
	row, col := GetBoardRowAndCol(validatedFieldPosition, board)
	takenFieldIndex := 0

	// marks board with player mark
	board.fields[row][col] = playerMark

	// finds takenField index for removal
	for i, val := range board.freeFields {
		if val == validatedFieldPosition {
			takenFieldIndex = i
		}
	}

	// pops taken field from freeFields list
	board.freeFields = append(board.freeFields[:takenFieldIndex], board.freeFields[takenFieldIndex+1:]...)
}

// Checks if a given element with a rowIndex and colIndex are on a diagonal of the board matrix
func IsFieldPositionOnBoardDiagonal(rowIndex int, colIndex int) bool {
	return rowIndex == colIndex
}

// Checks if a given element with a rowIndex and colIndex are on a REVERSE diagonal of the board matrix
func IsFieldPositionOnBoardReverseDiagonal(rowIndex int, colIndex int, board *Board) bool {
	boardSize := len(board.fields)

	return rowIndex+colIndex == (boardSize - 1) // math relation that asserts that
}

// Based on a row and col index, returns the whole row diagonal slice, if row and col is
// from a diagonal element. Otherwise, returns an empty string slice.
func GetBoardDiagonalSlice(board *Board) []string {
	diagonalSlice := make([]string, 0) // starts as an empty slice

	// appends only diagonal elements from the board
	for i := 0; i < len(board.fields); i++ {
		diagonalSlice = append(diagonalSlice, board.fields[i][i])
	}

	return diagonalSlice
}

// Based on a row and col index, returns the whole REVERSE row diagonal slice, if row and col is
// from a reverse diagonal element. Otherwise, returns an empty string slice.
func GetBoardReverseDiagonalSlice(board *Board) []string {
	reverseDiagonalSlice := make([]string, 0) // starts as an empty slice

	// appends only diagonal elements from the board
	for i := 0; i < len(board.fields); i++ {

		// naive approach
		for j := 0; j < len(board.fields); j++ {

			if IsFieldPositionOnBoardReverseDiagonal(i, j, board) {
				reverseDiagonalSlice = append(reverseDiagonalSlice, board.fields[i][j])
			}
		}
	}

	return reverseDiagonalSlice
}

// Based on a row index, returns the whole row slice from the board
func GetBoardRowSlice(rowIndex int, board *Board) []string {
	return board.fields[rowIndex]
}

// Based on a column index, returns each element from that column number on the board
func GetBoardColumnSlice(colIndex int, board *Board) []string {
	columnSlice := make([]string, 0) // empty slice of strings

	// traverses each row of the 2D board fields getting the right col member
	for _, row := range board.fields {
		columnSlice = append(columnSlice, row[colIndex])
	}

	return columnSlice
}
