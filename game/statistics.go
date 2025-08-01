package game

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Statistics represents game statistics
type Statistics struct {
	GamesPlayed    int            `json:"games_played"`
	GamesWon       int            `json:"games_won"`
	GamesLost      int            `json:"games_lost"`
	TotalGuesses   int            `json:"total_guesses"`
	CorrectGuesses int            `json:"correct_guesses"`
	WrongGuesses   int            `json:"wrong_guesses"`
	BestGame       int            `json:"best_game"`      // Fewest wrong guesses in a won game
	CurrentStreak  int            `json:"current_streak"` // Current winning streak
	LongestStreak  int            `json:"longest_streak"` // Longest winning streak
	LastPlayed     time.Time      `json:"last_played"`
	WordsGuessed   []string       `json:"words_guessed"` // Recently guessed words
	Difficulties   map[string]int `json:"difficulties"`  // Games played per difficulty
}

// NewStatistics creates a new statistics instance
func NewStatistics() *Statistics {
	return &Statistics{
		WordsGuessed: make([]string, 0),
		Difficulties: make(map[string]int),
		BestGame:     6, // Start with worst possible score
	}
}

// LoadStatistics loads statistics from file
func LoadStatistics() (*Statistics, error) {
	statsFile := getStatsFilePath()

	// If file doesn't exist, return new statistics
	if _, err := os.Stat(statsFile); os.IsNotExist(err) {
		return NewStatistics(), nil
	}

	//nolint:gosec // G304: File path is controlled by the application
	data, err := os.ReadFile(statsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read statistics file: %w", err)
	}

	var stats Statistics
	if err := json.Unmarshal(data, &stats); err != nil {
		return nil, fmt.Errorf("failed to parse statistics: %w", err)
	}

	// Initialize maps if they're nil (for backward compatibility)
	if stats.Difficulties == nil {
		stats.Difficulties = make(map[string]int)
	}
	if stats.WordsGuessed == nil {
		stats.WordsGuessed = make([]string, 0)
	}

	return &stats, nil
}

// SaveStatistics saves statistics to file
func (s *Statistics) SaveStatistics() error {
	statsFile := getStatsFilePath()

	// Create directory if it doesn't exist
	dir := filepath.Dir(statsFile)
	if err := os.MkdirAll(dir, 0o750); err != nil { // More restrictive permissions
		return fmt.Errorf("failed to create stats directory: %w", err)
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal statistics: %w", err)
	}

	// Write to file with restricted permissions
	if err := os.WriteFile(statsFile, data, 0o600); err != nil { // More restrictive permissions
		return fmt.Errorf("failed to write statistics file: %w", err)
	}

	return nil
}

// RecordGame records the results of a completed game
func (s *Statistics) RecordGame(g *Game, difficulty string) {
	s.GamesPlayed++
	s.LastPlayed = time.Now()
	s.TotalGuesses += len(g.GuessedLetters)
	s.WrongGuesses += g.WrongGuesses
	s.CorrectGuesses += len(g.GuessedLetters) - g.WrongGuesses

	// Record difficulty
	s.Difficulties[difficulty]++

	// Add word to recently guessed (keep last 10)
	s.WordsGuessed = append(s.WordsGuessed, g.Word)
	if len(s.WordsGuessed) > 10 {
		s.WordsGuessed = s.WordsGuessed[1:]
	}

	if g.IsWon {
		s.GamesWon++
		s.CurrentStreak++

		// Update best game (fewest wrong guesses)
		if g.WrongGuesses < s.BestGame {
			s.BestGame = g.WrongGuesses
		}

		// Update longest streak
		if s.CurrentStreak > s.LongestStreak {
			s.LongestStreak = s.CurrentStreak
		}
	} else {
		s.GamesLost++
		s.CurrentStreak = 0 // Reset streak on loss
	}
}

// GetWinRate returns the win rate as a percentage
func (s *Statistics) GetWinRate() float64 {
	if s.GamesPlayed == 0 {
		return 0.0
	}
	return float64(s.GamesWon) / float64(s.GamesPlayed) * 100
}

// GetAverageGuesses returns the average number of guesses per game
func (s *Statistics) GetAverageGuesses() float64 {
	if s.GamesPlayed == 0 {
		return 0.0
	}
	return float64(s.TotalGuesses) / float64(s.GamesPlayed)
}

// GetGuessAccuracy returns the accuracy of guesses as a percentage
func (s *Statistics) GetGuessAccuracy() float64 {
	if s.TotalGuesses == 0 {
		return 0.0
	}
	return float64(s.CorrectGuesses) / float64(s.TotalGuesses) * 100
}

// PrintStatistics prints formatted statistics
func (s *Statistics) PrintStatistics() {
	fmt.Println("📊 GAME STATISTICS")
	fmt.Println("==================")
	fmt.Printf("Games Played: %d\n", s.GamesPlayed)
	fmt.Printf("Games Won: %d\n", s.GamesWon)
	fmt.Printf("Games Lost: %d\n", s.GamesLost)
	fmt.Printf("Win Rate: %.1f%%\n", s.GetWinRate())
	fmt.Printf("Current Streak: %d\n", s.CurrentStreak)
	fmt.Printf("Longest Streak: %d\n", s.LongestStreak)

	if s.GamesWon > 0 {
		fmt.Printf("Best Game: %d wrong guesses\n", s.BestGame)
	}

	fmt.Printf("Average Guesses: %.1f\n", s.GetAverageGuesses())
	fmt.Printf("Guess Accuracy: %.1f%%\n", s.GetGuessAccuracy())

	if len(s.Difficulties) > 0 {
		fmt.Println("\nGames by Difficulty:")
		for difficulty, count := range s.Difficulties {
			fmt.Printf("  %s: %d\n", difficulty, count)
		}
	}

	if len(s.WordsGuessed) > 0 {
		fmt.Println("\nRecently Guessed Words:")
		for i := len(s.WordsGuessed) - 1; i >= 0 && i >= len(s.WordsGuessed)-5; i-- {
			fmt.Printf("  %s\n", s.WordsGuessed[i])
		}
	}

	if !s.LastPlayed.IsZero() {
		fmt.Printf("\nLast Played: %s\n", s.LastPlayed.Format("2006-01-02 15:04:05"))
	}

	fmt.Println()
}

// getStatsFilePath returns the path to the statistics file
func getStatsFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ".hangman_stats.json" // Fallback to current directory
	}
	return filepath.Join(homeDir, ".hangman", "stats.json")
}

// ResetStatistics resets all statistics
func (s *Statistics) ResetStatistics() {
	*s = *NewStatistics()
}
