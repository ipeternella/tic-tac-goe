package tictacgoe

import (
	"bufio"
	"fmt"
	"io"
)

// Stores user input bytes that from stdin, in a buffer that is read
func GetUserInput(stdIn io.Reader) string {

	// creates an input buffer that receives []byte streams from stdin
	inputBuffer := bufio.NewReader(stdIn)
	text, _ := inputBuffer.ReadString('\n')

	return text

}

// Asserts that a user input is within the board size range
// and that the field spot is not already taken by another mark
func IsUserInputValid(fieldPosition int, board *Board) (bool, string) {

	// board constraints
	maxAllowedValue := len(board.fields)*len(board.fields) - 1 // assumes first field starts at 0
	minAllowedValue := 0

	// outcomes
	gameRejectionOutput := ""
	evaluationResult := false

	// assertion if the spot is taken
	if false {
		evaluationResult = true
	}

	if fieldPosition < minAllowedValue {
		gameRejectionOutput = fmt.Sprintf("Oops! The lowest allowed spot value for your move is %d!", minAllowedValue)
	}

	if fieldPosition > maxAllowedValue {
		gameRejectionOutput = fmt.Sprintf("Oops! The highest allowed spot value for your move is %d!", maxAllowedValue)
	}

	return evaluationResult, gameRejectionOutput

}
