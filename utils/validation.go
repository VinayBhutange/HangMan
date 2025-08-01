package utils

import (
	"unicode"
)

// IsLetter checks if a rune is a letter (A-Z or a-z)
func IsLetter(r rune) bool {
	return unicode.IsLetter(r)
}

// IsAlphabetic checks if a string contains only alphabetic characters
func IsAlphabetic(s string) bool {
	if len(s) == 0 {
		return false
	}
	
	for _, r := range s {
		if !IsLetter(r) {
			return false
		}
	}
	return true
}

// IsValidWord checks if a word is valid for the game
func IsValidWord(word string) bool {
	// Word must be at least 3 characters
	if len(word) < 3 {
		return false
	}
	
	// Word must contain only letters
	if !IsAlphabetic(word) {
		return false
	}
	
	return true
}

// IsValidDifficulty checks if the difficulty level is valid
func IsValidDifficulty(difficulty string) bool {
	switch difficulty {
	case "easy", "medium", "hard":
		return true
	default:
		return false
	}
}

// ContainsRune checks if a slice of runes contains a specific rune
func ContainsRune(slice []rune, target rune) bool {
	for _, r := range slice {
		if r == target {
			return true
		}
	}
	return false
}

// ValidateLetterGuess validates a letter guess
func ValidateLetterGuess(input string) (rune, error) {
	if len(input) == 0 {
		return 0, NewValidationError("input cannot be empty")
	}
	
	if len(input) > 1 {
		return 0, NewValidationError("please enter only one letter")
	}
	
	letter := rune(input[0])
	
	if !IsLetter(letter) {
		return 0, NewValidationError("please enter a valid letter (A-Z)")
	}
	
	return letter, nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

// NewValidationError creates a new validation error
func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return e.Message
}

// IsValidationError checks if an error is a validation error
func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}
