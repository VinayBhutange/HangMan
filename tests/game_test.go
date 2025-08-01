package tests

import (
	"testing"

	"github.com/VinayBhutange/hangman-go/game"
)

func TestNewGame(t *testing.T) {
	words := []string{"GOLANG", "PROGRAMMING", "COMPUTER"}
	g := game.NewGame(words)
	
	if g == nil {
		t.Fatal("NewGame returned nil")
	}
	
	if g.WrongGuesses != 0 {
		t.Errorf("Expected WrongGuesses to be 0, got %d", g.WrongGuesses)
	}
	
	if g.MaxWrongGuesses != 6 {
		t.Errorf("Expected MaxWrongGuesses to be 6, got %d", g.MaxWrongGuesses)
	}
	
	if g.IsGameOver {
		t.Error("Expected IsGameOver to be false")
	}
	
	if g.IsWon {
		t.Error("Expected IsWon to be false")
	}
	
	if len(g.GuessedLetters) != 0 {
		t.Errorf("Expected GuessedLetters to be empty, got %d items", len(g.GuessedLetters))
	}
}

func TestGuessLetter(t *testing.T) {
	words := []string{"GOLANG"}
	g := game.NewGame(words)
	g.Word = "GOLANG" // Set specific word for testing
	
	// Test correct guess
	isCorrect := g.GuessLetter('G')
	if !isCorrect {
		t.Error("Expected 'G' to be correct")
	}
	
	if g.WrongGuesses != 0 {
		t.Errorf("Expected WrongGuesses to be 0, got %d", g.WrongGuesses)
	}
	
	// Test wrong guess
	isCorrect = g.GuessLetter('X')
	if isCorrect {
		t.Error("Expected 'X' to be incorrect")
	}
	
	if g.WrongGuesses != 1 {
		t.Errorf("Expected WrongGuesses to be 1, got %d", g.WrongGuesses)
	}
	
	// Test duplicate guess
	isCorrect = g.GuessLetter('G')
	if isCorrect {
		t.Error("Expected duplicate guess to return false")
	}
}

func TestGetDisplayWord(t *testing.T) {
	words := []string{"GOLANG"}
	g := game.NewGame(words)
	g.Word = "GOLANG"
	
	// No letters guessed
	display := g.GetDisplayWord()
	expected := "_ _ _ _ _ _"
	if display != expected {
		t.Errorf("Expected '%s', got '%s'", expected, display)
	}
	
	// Guess some letters
	g.GuessLetter('G')
	g.GuessLetter('A')
	
	display = g.GetDisplayWord()
	expected = "G _ _ A _ G"  // Fixed: 'G' appears twice in "GOLANG"
	if display != expected {
		t.Errorf("Expected '%s', got '%s'", expected, display)
	}
}

func TestIsWordComplete(t *testing.T) {
	words := []string{"GO"}
	g := game.NewGame(words)
	g.Word = "GO"
	
	// Word not complete
	if g.IsWordComplete() {
		t.Error("Expected word to be incomplete")
	}
	
	// Guess all letters
	g.GuessLetter('G')
	g.GuessLetter('O')
	
	// Word should be complete
	if !g.IsWordComplete() {
		t.Error("Expected word to be complete")
	}
	
	// Game should be won
	if !g.IsWon {
		t.Error("Expected game to be won")
	}
	
	if !g.IsGameOver {
		t.Error("Expected game to be over")
	}
}

func TestGameLoss(t *testing.T) {
	words := []string{"GOLANG"}
	g := game.NewGame(words)
	g.Word = "GOLANG"
	
	// Make maximum wrong guesses
	wrongLetters := []rune{'X', 'Y', 'Z', 'Q', 'W', 'E'}
	
	for _, letter := range wrongLetters {
		g.GuessLetter(letter)
	}
	
	// Game should be over and lost
	if !g.IsGameOver {
		t.Error("Expected game to be over")
	}
	
	if g.IsWon {
		t.Error("Expected game to be lost")
	}
	
	if g.WrongGuesses != 6 {
		t.Errorf("Expected 6 wrong guesses, got %d", g.WrongGuesses)
	}
}

func TestGetRemainingGuesses(t *testing.T) {
	words := []string{"GOLANG"}
	g := game.NewGame(words)
	
	// No wrong guesses yet
	remaining := g.GetRemainingGuesses()
	if remaining != 6 {
		t.Errorf("Expected 6 remaining guesses, got %d", remaining)
	}
	
	// Make some wrong guesses
	g.GuessLetter('X')
	g.GuessLetter('Y')
	
	remaining = g.GetRemainingGuesses()
	if remaining != 4 {
		t.Errorf("Expected 4 remaining guesses, got %d", remaining)
	}
}
