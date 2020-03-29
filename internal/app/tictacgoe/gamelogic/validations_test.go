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

// Asserts field value has NO digits
func TestHasNoDigits(t *testing.T) {
	value := "IG"

	hasDigits := HasDigits(value)

	if hasDigits {
		t.Errorf("Expected string NOT to have digits. actual: %t, expected: %t", hasDigits, false)
	}
}

// Asserts field value has only digits
func TestHasDigits(t *testing.T) {
	hasDigits := HasDigits("00")

	if !hasDigits {
		t.Errorf("Expected string to have digits. actual: %t, expected: %t", hasDigits, true)
	}

	hasDigits = HasDigits("81")

	if !hasDigits {
		t.Errorf("Expected string to have digits. actual: %t, expected: %t", hasDigits, true)
	}
}

func TestIsFieldPositionTaken(t *testing.T) {
	gameBoard := CreateGameBoard(3)

	gameBoard.fields[0][1] = "IG" // TAKEN by a player's mark!

	isPositionTaken1 := IsFieldPositionTaken(0, gameBoard) // NOT TAKEN!
	isPositionTaken2 := IsFieldPositionTaken(1, gameBoard) // TAKEN!

	// something's wrong if this spot IS taken...
	if isPositionTaken1 {
		t.Errorf("Position taken assertion failed. actual: %t, expected: %t", isPositionTaken1, false)
	}

	// something's wrong if this spot is NOT taken...
	if !isPositionTaken2 {
		t.Errorf("Position taken assertion failed. actual: %t, expected: %t", isPositionTaken2, true)
	}
}
