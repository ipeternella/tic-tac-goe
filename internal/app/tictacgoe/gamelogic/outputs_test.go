package gamelogic

import (
	"testing"
)

// Asserts display screen message displays strings without a new line at the end
func TestDisplayScreenMessageWithoutNewLine(t *testing.T) {
	message := "Hello world!"
	expectedMessage := message

	// function invocation that sends byte streams to stdout
	messageForStdOut := DisplayScreenMessage(message, false)

	if messageForStdOut != message {
		t.Errorf("Message sent to std out mismatch! actual: %s, expected: %s", messageForStdOut, expectedMessage)
	}

}

// Asserts display screen message displays strings WITH a trailing new line
func TestDisplayScreenMessageWithNewLineAtTheEnd(t *testing.T) {
	message := "Hello world!"
	expectedMessage := message + "\n"

	// function invocation that sends byte streams to stdout
	messageForStdOut := DisplayScreenMessage(message, true)

	if messageForStdOut != expectedMessage {
		t.Errorf("Message sent to std out mismatch! actual: %s, expected: %s", messageForStdOut, expectedMessage)
	}

}
