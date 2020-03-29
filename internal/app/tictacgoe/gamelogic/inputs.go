// Reads user input
package gamelogic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Stores user input bytes that from stdin, in a buffer that is read
func ReadUserInput(stdIn io.Reader) string {

	// creates an input buffer that receives []byte streams from stdin
	inputBuffer := bufio.NewReader(stdIn)
	text, _ := inputBuffer.ReadString('\n')

	return strings.TrimSuffix(text, "\n")

}

// Asks for input but accepts everything
func AskUserWithoutOptions(question string, newline bool) string {
	DisplayScreenMessage(question, newline)
	rawUserInput := ReadUserInput(os.Stdin)

	return rawUserInput
}

// Asks for input, with valid options, and only leave when valid input is received
func AskUserWithOptions(question string, validOptions []string, repeatQuestion bool, repeatQuestionPrefix string) string {
	fmt.Println(question)

	for {
		rawUserInput := ReadUserInput(os.Stdin)

		// searches for valid option
		for _, validItem := range validOptions {
			if rawUserInput == validItem {
				return rawUserInput // validated
			}
		}

		// invalid option
		if repeatQuestion {
			DisplayScreenMessage(repeatQuestionPrefix+question, false)
		}
	}
}
