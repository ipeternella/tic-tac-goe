package tictacgoe

import (
	"bytes"
	"testing"
)

func TestGetUserInputFromStdIn(t *testing.T) {

	// expected values
	expectedInput := "Hello world!"

	// Creates a bytes buffer that we can write []byte streams to (~fake user input from stdin)
	var buffer bytes.Buffer

	// Writes to std input
	if _, err := buffer.Write([]byte("Hello world!")); err != nil {
		t.Error("Error when writing to std input of the test process...")
		return
	}

	input := GetUserInput(&buffer)

	if input != expectedInput {
		t.Errorf("Unexpected input. actual: %s, expected: %s", input, expectedInput)
	}

}

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
