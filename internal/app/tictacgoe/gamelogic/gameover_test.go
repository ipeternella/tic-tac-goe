package gamelogic

import "testing"

// asserts false for slice with digits (players haven't filled it yet)
func TestIsSliceFilledWithSinglePlayerMarkSliceFalseForSliceWithDigits(t *testing.T) {
	boardSlice := []string{"00", "01", "02"} // Slice with digits == false

	// function invocation
	result := IsSliceFilledWithSinglePlayerMark(boardSlice)

	// if returns true, then there's a problem!
	if result {
		t.Errorf("Board slice check should be false. actual :%t, expected: %t", result, false)
	}
}

// asserts false for slice with marks of both players and digits
func TestIsSliceFilledWithSinglePlayerMarkSliceFalseForSliceWithoutSingleMarkAndDigits(t *testing.T) {
	boardSlice := []string{"IG", "01", "PC"} // Slice has two marks: 'IG' and 'PC' and digits --> False

	// function invocation
	result := IsSliceFilledWithSinglePlayerMark(boardSlice)

	// if returns true, then there's a problem!
	if result {
		t.Errorf("Board slice check should be false. actual :%t, expected: %t", result, false)
	}
}

// asserts false for slice with marks of both players
func TestIsSliceFilledWithSinglePlayerMarkSliceFalseForSliceWithoutSingleMark(t *testing.T) {
	boardSlice := []string{"IG", "IG", "PC"} // Slice has two marks: 'IG' and 'PC' --> false

	// function invocation
	result := IsSliceFilledWithSinglePlayerMark(boardSlice)

	// if returns true, then there's a problem!
	if result {
		t.Errorf("Board slice check should be false. actual :%t, expected: %t", result, false)
	}
}

// asserts false for slice with marks of both players
func TestIsSliceFilledWithSinglePlayerMarkSliceTrueForSliceWithSingleMark(t *testing.T) {
	boardSlice := []string{"IG", "IG", "IG"} // Slice has two marks: 'IG' and 'PC' --> false

	// function invocation
	result := IsSliceFilledWithSinglePlayerMark(boardSlice)

	// if returns false, then there's a problem!
	if !result {
		t.Errorf("Board slice check should be false. actual :%t, expected: %t", result, true)
	}
}
