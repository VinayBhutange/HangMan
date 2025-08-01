package game

import (
	"testing"
)

func TestNewGameInternal(t *testing.T) {
	words := []string{"GOLANG", "PROGRAMMING", "COMPUTER"}
	g := NewGame(words)
	
	if g == nil {
		t.Fatal("NewGame returned nil")
	}
	
	if g.Word == "" {
		t.Error("Game word should not be empty")
	}
	
	if g.WrongGuesses != 0 {
		t.Errorf("Expected WrongGuesses to be 0, got %d", g.WrongGuesses)
	}
	
	if g.MaxWrongGuesses != 6 {
		t.Errorf("Expected MaxWrongGuesses to be 6, got %d", g.MaxWrongGuesses)
	}
}

func TestGuessLetterInternal(t *testing.T) {
	words := []string{"GOLANG"}
	g := NewGame(words)
	g.Word = "GOLANG" // Set known word for testing
	
	// Test correct guess
	correct := g.GuessLetter('G')
	if !correct {
		t.Error("Expected G to be correct")
	}
	
	// Test incorrect guess
	incorrect := g.GuessLetter('Z')
	if incorrect {
		t.Error("Expected Z to be incorrect")
	}
	
	if g.WrongGuesses != 1 {
		t.Errorf("Expected WrongGuesses to be 1, got %d", g.WrongGuesses)
	}
}

func TestGetDisplayWordInternal(t *testing.T) {
	words := []string{"GOLANG"}
	g := NewGame(words)
	g.Word = "GOLANG"
	
	// Initially should show underscores
	display := g.GetDisplayWord()
	expected := "_ _ _ _ _ _"
	if display != expected {
		t.Errorf("Expected '%s', got '%s'", expected, display)
	}
	
	// After guessing G
	g.GuessLetter('G')
	display = g.GetDisplayWord()
	expected = "G _ _ _ _ G"
	if display != expected {
		t.Errorf("Expected '%s', got '%s'", expected, display)
	}
}

func TestIsWordCompleteInternal(t *testing.T) {
	words := []string{"GO"}
	g := NewGame(words)
	g.Word = "GO"
	
	// Initially not complete
	if g.IsWordComplete() {
		t.Error("Word should not be complete initially")
	}
	
	// Guess all letters
	g.GuessLetter('G')
	g.GuessLetter('O')
	
	if !g.IsWordComplete() {
		t.Error("Word should be complete after guessing all letters")
	}
}

func TestGetDefaultWordsInternal(t *testing.T) {
	wordList := GetDefaultWords()
	
	if wordList == nil {
		t.Fatal("GetDefaultWords returned nil")
	}
	
	if len(wordList.Words) == 0 {
		t.Error("Default word list should not be empty")
	}
	
	// Check that all words are uppercase and have minimum length
	for _, word := range wordList.Words {
		if len(word) < 3 {
			t.Errorf("Word '%s' is too short (minimum 3 characters)", word)
		}
		
		// Check if word is uppercase
		for _, char := range word {
			if char < 'A' || char > 'Z' {
				t.Errorf("Word '%s' contains non-uppercase character '%c'", word, char)
			}
		}
	}
}

func TestStatisticsInternal(t *testing.T) {
	stats := NewStatistics()
	
	if stats == nil {
		t.Fatal("NewStatistics returned nil")
	}
	
	if stats.GamesPlayed != 0 {
		t.Errorf("Expected GamesPlayed to be 0, got %d", stats.GamesPlayed)
	}
	
	if stats.GamesWon != 0 {
		t.Errorf("Expected GamesWon to be 0, got %d", stats.GamesWon)
	}
}
