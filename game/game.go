// Package game implements the core hangman game logic and state management.
package game

import (
	"math/rand"
	"strings"
	"time"
)

// Game represents the current state of a hangman game
type Game struct {
	Word            string        // The word to guess
	GuessedLetters  map[rune]bool // Letters that have been guessed
	WrongGuesses    int           // Number of wrong guesses made
	MaxWrongGuesses int           // Maximum wrong guesses allowed
	IsGameOver      bool          // Whether the game has ended
	IsWon           bool          // Whether the player has won
}

// NewGame creates a new hangman game with a random word
func NewGame(words []string) *Game {
	if len(words) == 0 {
		panic("No words provided for the game")
	}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Select a random word (using math/rand is fine for games)
	//nolint:gosec // G404: Using weak random number generator is acceptable for game purposes
	word := strings.ToUpper(words[rand.Intn(len(words))])

	return &Game{
		Word:            word,
		GuessedLetters:  make(map[rune]bool),
		WrongGuesses:    0,
		MaxWrongGuesses: 6, // Standard hangman allows 6 wrong guesses
		IsGameOver:      false,
		IsWon:           false,
	}
}

// GuessLetter processes a letter guess and returns true if correct
func (g *Game) GuessLetter(letter rune) bool {
	if g.IsGameOver {
		return false
	}

	// Convert to uppercase for consistency
	letter = rune(strings.ToUpper(string(letter))[0])

	// Check if letter was already guessed
	if g.GuessedLetters[letter] {
		return false // Already guessed, no change in state
	}

	// Mark letter as guessed
	g.GuessedLetters[letter] = true

	// Check if letter is in the word
	isCorrect := strings.ContainsRune(g.Word, letter)

	if !isCorrect {
		g.WrongGuesses++
	}

	// Update game state
	g.updateGameState()

	return isCorrect
}

// GetDisplayWord returns the word with unguessed letters as underscores
func (g *Game) GetDisplayWord() string {
	var result strings.Builder

	for i, letter := range g.Word {
		if i > 0 {
			result.WriteString(" ")
		}

		if g.GuessedLetters[letter] {
			result.WriteRune(letter)
		} else {
			result.WriteString("_")
		}
	}

	return result.String()
}

// IsWordComplete checks if all letters in the word have been guessed
func (g *Game) IsWordComplete() bool {
	for _, letter := range g.Word {
		if !g.GuessedLetters[letter] {
			return false
		}
	}
	return true
}

// GetGuessedLetters returns a slice of all guessed letters
func (g *Game) GetGuessedLetters() []rune {
	var letters []rune
	for letter := range g.GuessedLetters {
		letters = append(letters, letter)
	}
	return letters
}

// GetWrongLetters returns a slice of incorrectly guessed letters
func (g *Game) GetWrongLetters() []rune {
	var wrongLetters []rune
	for letter := range g.GuessedLetters {
		if !strings.ContainsRune(g.Word, letter) {
			wrongLetters = append(wrongLetters, letter)
		}
	}
	return wrongLetters
}

// GetRemainingGuesses returns the number of remaining wrong guesses
func (g *Game) GetRemainingGuesses() int {
	return g.MaxWrongGuesses - g.WrongGuesses
}

// updateGameState checks and updates the game over conditions
func (g *Game) updateGameState() {
	// Check if player won
	if g.IsWordComplete() {
		g.IsGameOver = true
		g.IsWon = true
		return
	}

	// Check if player lost
	if g.WrongGuesses >= g.MaxWrongGuesses {
		g.IsGameOver = true
		g.IsWon = false
		return
	}
}

// Reset resets the game with a new word
func (g *Game) Reset(words []string) {
	if len(words) == 0 {
		panic("No words provided for the game")
	}

	//nolint:gosec // G404: Using weak random number generator is acceptable for game purposes
	word := strings.ToUpper(words[rand.Intn(len(words))])

	g.Word = word
	g.GuessedLetters = make(map[rune]bool)
	g.WrongGuesses = 0
	g.IsGameOver = false
	g.IsWon = false
}
