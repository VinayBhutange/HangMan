package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// WordList represents a collection of words for the game
type WordList struct {
	Words []string
}

// LoadWordsFromFile loads words from a text file
func LoadWordsFromFile(filename string) (*WordList, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open word file: %w", err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" && len(word) >= 3 { // Only include words with 3+ letters
			words = append(words, strings.ToUpper(word))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading word file: %w", err)
	}

	if len(words) == 0 {
		return nil, fmt.Errorf("no valid words found in file")
	}

	return &WordList{Words: words}, nil
}

// GetDefaultWords returns a default set of words if no file is available
func GetDefaultWords() *WordList {
	words := []string{
		"GOLANG", "PROGRAMMING", "COMPUTER", "KEYBOARD", "MONITOR",
		"FUNCTION", "VARIABLE", "PACKAGE", "INTERFACE", "STRUCT",
		"SLICE", "MAP", "CHANNEL", "GOROUTINE", "POINTER",
		"ALGORITHM", "DATABASE", "NETWORK", "SERVER", "CLIENT",
		"TERMINAL", "COMMAND", "EXECUTE", "COMPILE", "DEBUG",
		"SYNTAX", "SEMANTIC", "LEXICAL", "PARSER", "COMPILER",
		"RUNTIME", "MEMORY", "STACK", "HEAP", "GARBAGE",
		"COLLECTION", "THREAD", "PROCESS", "OPERATING", "SYSTEM",
		"FILE", "DIRECTORY", "PATH", "EXTENSION", "BINARY",
		"SOURCE", "CODE", "LOGIC", "CONDITION", "LOOP",
		"RECURSIVE", "ITERATION", "ARRAY", "LIST", "TREE",
		"GRAPH", "NODE", "EDGE", "VERTEX", "SEARCH",
		"SORT", "MERGE", "QUICK", "BUBBLE", "INSERTION",
	}
	
	return &WordList{Words: words}
}

// GetRandomWord returns a random word from the word list
func (wl *WordList) GetRandomWord() string {
	if len(wl.Words) == 0 {
		return "GOLANG" // Fallback word
	}
	
	rand.Seed(time.Now().UnixNano())
	return wl.Words[rand.Intn(len(wl.Words))]
}

// GetWordsByLength returns words of a specific length range
func (wl *WordList) GetWordsByLength(minLen, maxLen int) []string {
	var filtered []string
	
	for _, word := range wl.Words {
		if len(word) >= minLen && len(word) <= maxLen {
			filtered = append(filtered, word)
		}
	}
	
	return filtered
}

// GetWordsByDifficulty returns words based on difficulty level
func (wl *WordList) GetWordsByDifficulty(difficulty string) []string {
	switch strings.ToLower(difficulty) {
	case "easy":
		return wl.GetWordsByLength(4, 5)
	case "medium":
		return wl.GetWordsByLength(6, 8)
	case "hard":
		return wl.GetWordsByLength(9, 15)
	default:
		return wl.Words
	}
}

// AddWord adds a new word to the word list
func (wl *WordList) AddWord(word string) {
	word = strings.ToUpper(strings.TrimSpace(word))
	if word != "" && len(word) >= 3 {
		wl.Words = append(wl.Words, word)
	}
}

// RemoveWord removes a word from the word list
func (wl *WordList) RemoveWord(word string) {
	word = strings.ToUpper(strings.TrimSpace(word))
	for i, w := range wl.Words {
		if w == word {
			wl.Words = append(wl.Words[:i], wl.Words[i+1:]...)
			break
		}
	}
}

// GetWordCount returns the total number of words
func (wl *WordList) GetWordCount() int {
	return len(wl.Words)
}
