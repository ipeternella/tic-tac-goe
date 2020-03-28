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
