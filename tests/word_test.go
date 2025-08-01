package tests

import (
	"testing"

	"github.com/VinayBhutange/hangman-go/game"
)

const (
	testWordProgramming = "PROGRAMMING"
)

func TestGetDefaultWords(t *testing.T) {
	wordList := game.GetDefaultWords()

	if wordList == nil {
		t.Fatal("GetDefaultWords returned nil")
	}

	if len(wordList.Words) == 0 {
		t.Error("Expected default words to be non-empty")
	}

	// Check that all words are uppercase
	for _, word := range wordList.Words {
		for _, char := range word {
			if char < 'A' || char > 'Z' {
				t.Errorf("Word '%s' contains non-uppercase character '%c'", word, char)
			}
		}
	}
}

func TestGetRandomWord(t *testing.T) {
	words := []string{"GOLANG", "PROGRAMMING", "COMPUTER"}
	wordList := &game.WordList{Words: words}

	randomWord := wordList.GetRandomWord()

	// Check that the random word is from our list
	found := false
	for _, word := range words {
		if word == randomWord {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Random word '%s' not found in word list", randomWord)
	}
}

func TestGetWordsByLength(t *testing.T) {
	words := []string{"GO", "LANG", "GOLANG", "PROGRAMMING"}
	wordList := &game.WordList{Words: words}

	// Test filtering by length
	shortWords := wordList.GetWordsByLength(2, 4)
	expectedShort := []string{"GO", "LANG"}

	if len(shortWords) != len(expectedShort) {
		t.Errorf("Expected %d short words, got %d", len(expectedShort), len(shortWords))
	}

	// Check contents
	for i, word := range shortWords {
		if word != expectedShort[i] {
			t.Errorf("Expected word '%s', got '%s'", expectedShort[i], word)
		}
	}
}

func TestGetWordsByDifficulty(t *testing.T) {
	words := []string{"EASY", "MEDIUM", "DIFFICULT", "PROGRAMMING"}
	wordList := &game.WordList{Words: words}

	// Test easy difficulty (4-5 letters)
	easyWords := wordList.GetWordsByDifficulty("easy")
	if len(easyWords) != 1 { // Only "EASY" (4 letters)
		t.Errorf("Expected 1 easy word, got %d", len(easyWords))
	}

	// Test medium difficulty (6-8 letters)
	mediumWords := wordList.GetWordsByDifficulty("medium")
	if len(mediumWords) != 1 { // Only "MEDIUM" (6 letters)
		t.Errorf("Expected 1 medium word, got %d", len(mediumWords))
	}

	// Test hard difficulty (9+ letters)
	hardWords := wordList.GetWordsByDifficulty("hard")
	if len(hardWords) != 2 { // "DIFFICULT" (9) and "PROGRAMMING" (11)
		t.Errorf("Expected 2 hard words, got %d", len(hardWords))
	}

	// Test invalid difficulty
	allWords := wordList.GetWordsByDifficulty("invalid")
	if len(allWords) != len(words) {
		t.Errorf("Expected all words for invalid difficulty, got %d", len(allWords))
	}
}

func TestAddWord(t *testing.T) {
	wordList := &game.WordList{Words: []string{"GOLANG"}}
	originalCount := len(wordList.Words)

	// Add valid word
	wordList.AddWord("programming")
	if len(wordList.Words) != originalCount+1 {
		t.Errorf("Expected %d words after adding, got %d", originalCount+1, len(wordList.Words))
	}

	// Check that word was added in uppercase
	if wordList.Words[len(wordList.Words)-1] != testWordProgramming {
		t.Error("Added word should be converted to uppercase")
	}

	// Try to add invalid word (too short)
	wordList.AddWord("GO")
	if len(wordList.Words) != originalCount+1 {
		t.Error("Short word should not be added")
	}

	// Try to add empty word
	wordList.AddWord("")
	if len(wordList.Words) != originalCount+1 {
		t.Error("Empty word should not be added")
	}
}

func TestRemoveWord(t *testing.T) {
	wordList := &game.WordList{Words: []string{"GOLANG", "PROGRAMMING", "COMPUTER"}}
	originalCount := len(wordList.Words)

	// Remove existing word
	wordList.RemoveWord("programming")
	if len(wordList.Words) != originalCount-1 {
		t.Errorf("Expected %d words after removing, got %d", originalCount-1, len(wordList.Words))
	}

	// Check that word was removed
	for _, word := range wordList.Words {
		if word == "PROGRAMMING" {
			t.Error("Word should have been removed")
		}
	}

	// Try to remove non-existent word
	wordList.RemoveWord("NONEXISTENT")
	if len(wordList.Words) != originalCount-1 {
		t.Error("Removing non-existent word should not change count")
	}
}

func TestGetWordCount(t *testing.T) {
	words := []string{"GOLANG", "PROGRAMMING", "COMPUTER"}
	wordList := &game.WordList{Words: words}

	count := wordList.GetWordCount()
	if count != len(words) {
		t.Errorf("Expected word count %d, got %d", len(words), count)
	}
}
