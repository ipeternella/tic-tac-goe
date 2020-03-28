package gamelogic

import "testing"

func TestIsUserInputValidMinThresholdExceeded(t *testing.T) {
	board := CreateGameBoard(3) // 3x3 board --> min value is 1
	expectedOutput := "Oops! The lowest allowed spot value for your move is 0!"

	// function invocation
	isValid, output := IsUserInputValid(-1, board) // -1 < 0: lower than min threshold!

	if isValid {
		t.Errorf("Expected input should be false. actual: %t, expected: %t", isValid, false)
	}

	if output != expectedOutput {
		t.Errorf("Game output is different than expected. actual: %v, expected: %v", output, expectedOutput)
	}
}

func TestIsUserInputValidMaxThresholdExceeded(t *testing.T) {
	board := CreateGameBoard(5) // 5 x 5 board --> max value is 25
	expectedOutput := "Oops! The highest allowed spot value for your move is 24!"

	// function invocation
	isValid, output := IsUserInputValid(25, board) // 25 > 24 (5*5 -1): exceeded max threshold!

	// result must not be valid
	if isValid {
		t.Errorf("Expected input should be false. actual: %t, expected: %t", isValid, false)
	}

	// output assertion
	if output != expectedOutput {
		t.Errorf("Game output is different than expected. actual: %v, expected: %v", output, expectedOutput)
	}
}

func TestIsUserInputValidForValidInput(t *testing.T) {
	board := CreateGameBoard(5) // 5 x 5 board --> max value is 25
	expectedOutput := ""

	// function invocation
	isValid, output := IsUserInputValid(5, board) // 0 <= 5 <= 24, ok!!

	// result is NOW valid
	if !isValid {
		t.Errorf("Expected input should be false. actual: %t, expected: %t", isValid, true)
	}

	// output assertion
	if output != expectedOutput {
		t.Errorf("Game output is different than expected. actual: %v, expected: %v", output, expectedOutput)
	}
}
