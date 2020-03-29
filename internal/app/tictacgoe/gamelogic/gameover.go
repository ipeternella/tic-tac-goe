// Functions that scan the board to check if the game's over!
package gamelogic

func IsSliceFilledWithSinglePlayerMark(line []string) bool {
	symbolsMap := make(map[string]bool) // map that works as a symbol count!

	// traverses whole slice appending new values to the symbols set
	for _, value := range line {
		symbolsMap[value] = true // if a new symbol is found, then a NEW key is added, otherwise remains as is!

		// slice is NOT homogeneous OR still has digits (not completely filled)
		if len(symbolsMap) > 1 || HasDigits(value) {
			return false
		}
	}

	// slice is homogeneous (symbols set remained as one)
	return true
}

// Evals the board to check for draws and wins/losses -- Naive implementation
func IsGameOver(board *Board) bool {
	// 1. Check for horizontal fills
	// 2. Check for vertical fills (~ same logic)
	// 3. Check for Diagonal 1 fills
	// 4. Check for Diagonal 2 fills (~ same logic)
	// logic will be the same, just will extract different slices from the board.fields!
	return false
}
