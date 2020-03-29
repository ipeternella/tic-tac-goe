package gamelogic

import (
	"fmt"
	"regexp"

	"github.com/IgooorGP/tic-tac-goe/internal/app/tictacgoe/settings"
)

// asserts that a given string value only has digits
func HasDigits(value string) bool {
	digitsPattern := regexp.MustCompile("[0-9]+") // one or more digits
	hasDigits := digitsPattern.MatchString(value)

	return hasDigits
}

// Checks if a given field position is already taken by another mark
// Field position must have been validated first.
func IsFieldPositionTaken(validatedFieldPosition int, board *Board) bool {
	fieldValue := GetFieldValue(validatedFieldPosition, board)
	hasDigitsOnly := HasDigits(fieldValue)

	return !hasDigitsOnly // no digits means a player mark is there!
}

// Asserts that a user input is within the board size range
// and that the field spot is not already taken by another mark
func IsUserInputValid(fieldPosition int, board *Board) (bool, string) {

	// board constraints
	maxAllowedValue := len(board.fields)*len(board.fields) - 1 // assumes first field starts at 0
	minAllowedValue := 0

	// outcomes
	gameRejectionOutput := ""
	evaluationResult := true

	// first sanity checks: fieldPosition MUST be within the board valid range
	if fieldPosition < minAllowedValue {
		gameRejectionOutput = fmt.Sprintf(settings.RejectedtMoveBelowMinValueMsg, minAllowedValue)
		evaluationResult = false
	}

	if fieldPosition > maxAllowedValue {
		gameRejectionOutput = fmt.Sprintf(settings.RejectedtMoveAboveMaxValueMsg, maxAllowedValue)
		evaluationResult = false
	}

	// assertion if the spot is taken
	if IsFieldPositionTaken(fieldPosition, board) {
		gameRejectionOutput = fmt.Sprintf(settings.RejectedMoveFieldPositionTaken, fieldPosition)
		evaluationResult = false
	}

	return evaluationResult, gameRejectionOutput

}
