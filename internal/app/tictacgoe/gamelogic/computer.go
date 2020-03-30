package gamelogic

import (
	"math/rand"
	"time"
)

// Calculates the next move of the PC. This function could add a lot of intelligence
// to the machine. For starters, however, this is generates random fields positions
func CalculateComputerMove(board *Board) int {
	rand.Seed(time.Now().Unix()) // random seed based on time

	totalChoices := len(board.freeFields)
	randomChoiceIndex := rand.Intn(totalChoices)
	computerMove := board.freeFields[randomChoiceIndex]

	return computerMove
}
