package utils

import (
	"testing"
)

func TestIsLetterInternal(t *testing.T) {
	// Test valid letters
	validLetters := []rune{'A', 'Z', 'a', 'z', 'M'}
	for _, letter := range validLetters {
		if !IsLetter(letter) {
			t.Errorf("Expected '%c' to be valid", letter)
		}
	}

	// Test invalid inputs
	invalidLetters := []rune{'1', '!', ' ', '@'}
	for _, letter := range invalidLetters {
		if IsLetter(letter) {
			t.Errorf("Expected '%c' to be invalid", letter)
		}
	}
}

func TestIsAlphabeticInternal(t *testing.T) {
	// Test valid strings
	validInputs := []string{"hello", "WORLD", "GoLang"}
	for _, input := range validInputs {
		if !IsAlphabetic(input) {
			t.Errorf("Expected '%s' to be alphabetic", input)
		}
	}

	// Test invalid strings
	invalidInputs := []string{"", "hello123", "test!", " "}
	for _, input := range invalidInputs {
		if IsAlphabetic(input) {
			t.Errorf("Expected '%s' to be non-alphabetic", input)
		}
	}
}

func TestContainsRuneInternal(t *testing.T) {
	letters := []rune{'A', 'B', 'C'}

	// Test letter that exists
	if !ContainsRune(letters, 'B') {
		t.Error("Expected to find 'B' in the slice")
	}

	// Test letter that doesn't exist
	if ContainsRune(letters, 'Z') {
		t.Error("Did not expect to find 'Z' in the slice")
	}
}

func TestIsValidWordInternal(t *testing.T) {
	// Test valid words
	validWords := []string{"GOLANG", "PROGRAMMING", "TEST"}
	for _, word := range validWords {
		if !IsValidWord(word) {
			t.Errorf("Expected '%s' to be valid", word)
		}
	}

	// Test invalid words
	invalidWords := []string{"", "A", "AB", "test123"}
	for _, word := range invalidWords {
		if IsValidWord(word) {
			t.Errorf("Expected '%s' to be invalid", word)
		}
	}
}

func TestIsValidDifficultyInternal(t *testing.T) {
	// Test valid difficulties
	validDifficulties := []string{"easy", "medium", "hard"}
	for _, difficulty := range validDifficulties {
		if !IsValidDifficulty(difficulty) {
			t.Errorf("Expected '%s' to be valid", difficulty)
		}
	}

	// Test invalid difficulties
	invalidDifficulties := []string{"", "impossible", "normal"}
	for _, difficulty := range invalidDifficulties {
		if IsValidDifficulty(difficulty) {
			t.Errorf("Expected '%s' to be invalid", difficulty)
		}
	}
}

func TestColorizeInternal(t *testing.T) {
	result := Colorize("test", Red(""))
	if result == "" {
		t.Error("Colorize should not return empty string")
	}

	// Test color functions
	if Red("test") == "" {
		t.Error("Red function should not return empty string")
	}

	if Green("test") == "" {
		t.Error("Green function should not return empty string")
	}
}
