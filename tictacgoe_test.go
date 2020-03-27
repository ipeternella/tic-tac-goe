package tictacgoe

import (
	"reflect"
	"testing"
)

// Asserts the creation of an initialized game board
func TestCreateGameBoardShouldCreateInitializeBoardFields3x3(t *testing.T) {
	boardSize := 3
	expectedBoard := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
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
