// Reads user input
package repl

import (
	"bufio"
	"io"
	"strings"
)

// Stores user input bytes that from stdin, in a buffer that is read
func GetUserInput(stdIn io.Reader) string {

	// creates an input buffer that receives []byte streams from stdin
	inputBuffer := bufio.NewReader(stdIn)
	text, _ := inputBuffer.ReadString('\n')

	return strings.TrimSuffix(text, "\n")

}
