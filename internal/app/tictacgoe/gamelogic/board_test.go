package gamelogic

import (
	"reflect"
	"testing"
)

// Asserts the creation of an initialized game board
func TestCreateGameBoardShouldCreateInitializeBoardFields3x3(t *testing.T) {
	boardSize := 4
	expectedBoard := [][]string{
		{"00", "01", "02", "03"},
		{"04", "05", "06", "07"},
		{"08", "09", "10", "11"},
		{"12", "13", "14", "15"},
	}

	// function invocation
	boardFields := CreateGameBoard(boardSize).fields

	// board assertions for each row
	for i := range expectedBoard {
		// compares each board row with the expected board
		if !reflect.DeepEqual(boardFields[i], expectedBoard[i]) {
			t.Errorf("Board row %d has failed assertion! board: %s, expected: %s", i, boardFields[i], expectedBoard[i])
		}
	}
}

// Asserts appropriate row printing
func TestPrivateGetRowStringBoard5x5(t *testing.T) {
	gameBoard := CreateGameBoard(4)
	expectedFirstRow := " 00 | 01 | 02 | 03 "
	expectedLastRow := " 12 | 13 | 14 | 15 "

	firstRow := getRowString(gameBoard, 0)
	lastRow := getRowString(gameBoard, len(gameBoard.fields)-1)

	if firstRow != expectedFirstRow {
		t.Errorf("First row does not match. Actual: %s, expected: %s", firstRow, expectedFirstRow)
	}

	if lastRow != expectedLastRow {
		t.Errorf("Last row does not match. Actual: %s, expected: %s", lastRow, expectedLastRow)
	}
}

// Asserts proper conversion of a field number to a row, col of the board matrix
// TODO: Improve this test with more clarity
func TestGetBoardRowAndCol5x5(t *testing.T) {
	board := CreateGameBoard(5)

	row, col := GetBoardRowAndCol(0, board)
	expectedRow, expectedCol := 0, 0

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(13, board)
	expectedRow, expectedCol = 2, 3

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(15, board)
	expectedRow, expectedCol = 3, 0

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(6, board)
	expectedRow, expectedCol = 1, 1

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(9, board)
	expectedRow, expectedCol = 1, 4

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

	row, col = GetBoardRowAndCol(24, board)
	expectedRow, expectedCol = 4, 4

	if row != expectedRow || col != expectedCol {
		t.Errorf("(Row, Col) mismatch. actual: (%d, %d), expected: (%d, %d)", row, col, expectedRow, expectedCol)
	}

}
