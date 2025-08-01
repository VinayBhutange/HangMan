package tests

import (
	"testing"
	"time"

	"github.com/VinayBhutange/hangman-go/game"
)

func TestNewStatistics(t *testing.T) {
	stats := game.NewStatistics()

	if stats == nil {
		t.Fatal("NewStatistics returned nil")
	}

	if stats.GamesPlayed != 0 {
		t.Errorf("Expected GamesPlayed to be 0, got %d", stats.GamesPlayed)
	}

	if stats.BestGame != 6 {
		t.Errorf("Expected BestGame to be 6, got %d", stats.BestGame)
	}

	if stats.Difficulties == nil {
		t.Error("Expected Difficulties map to be initialized")
	}

	if stats.WordsGuessed == nil {
		t.Error("Expected WordsGuessed slice to be initialized")
	}
}

func TestRecordGame(t *testing.T) {
	stats := game.NewStatistics()

	// Create a won game
	words := []string{testWordGo} // Use shorter word for easier testing
	g := game.NewGame(words)
	g.Word = testWordGo
	g.GuessLetter('X') // Wrong guess first
	g.GuessLetter('G')
	g.GuessLetter('O')

	// Debug output
	t.Logf("Game state: IsWon=%v, IsGameOver=%v, WrongGuesses=%d", g.IsWon, g.IsGameOver, g.WrongGuesses)
	t.Logf("Word complete: %v", g.IsWordComplete())

	// Verify game is won
	if !g.IsWon {
		t.Fatal("Game should be won")
	}

	// Record the game
	stats.RecordGame(g, "medium")

	// Check statistics
	if stats.GamesPlayed != 1 {
		t.Errorf("Expected GamesPlayed to be 1, got %d", stats.GamesPlayed)
	}

	if stats.GamesWon != 1 {
		t.Errorf("Expected GamesWon to be 1, got %d", stats.GamesWon)
	}

	if stats.CurrentStreak != 1 {
		t.Errorf("Expected CurrentStreak to be 1, got %d", stats.CurrentStreak)
	}

	if stats.BestGame != 1 { // 1 wrong guess
		t.Errorf("Expected BestGame to be 1, got %d", stats.BestGame)
	}

	if stats.Difficulties["medium"] != 1 {
		t.Errorf("Expected medium difficulty count to be 1, got %d", stats.Difficulties["medium"])
	}

	if len(stats.WordsGuessed) != 1 || stats.WordsGuessed[0] != testWordGo {
		t.Errorf("Expected WordsGuessed to contain %q, got %v", testWordGo, stats.WordsGuessed)
	}
}

func TestRecordLostGame(t *testing.T) {
	stats := game.NewStatistics()

	// Create a lost game
	words := []string{"GOLANG"}
	g := game.NewGame(words)
	g.Word = "GOLANG"

	// Make 6 wrong guesses
	wrongLetters := []rune{'X', 'Y', 'Z', 'Q', 'W', 'E'}
	for _, letter := range wrongLetters {
		g.GuessLetter(letter)
	}

	// Verify game is lost
	if !g.IsGameOver || g.IsWon {
		t.Fatal("Game should be lost")
	}

	// Record the game
	stats.RecordGame(g, "easy")

	// Check statistics
	if stats.GamesPlayed != 1 {
		t.Errorf("Expected GamesPlayed to be 1, got %d", stats.GamesPlayed)
	}

	if stats.GamesLost != 1 {
		t.Errorf("Expected GamesLost to be 1, got %d", stats.GamesLost)
	}

	if stats.GamesWon != 0 {
		t.Errorf("Expected GamesWon to be 0, got %d", stats.GamesWon)
	}

	if stats.CurrentStreak != 0 {
		t.Errorf("Expected CurrentStreak to be 0, got %d", stats.CurrentStreak)
	}
}

func TestWinStreak(t *testing.T) {
	stats := game.NewStatistics()

	// Simulate winning 3 games in a row
	for i := 0; i < 3; i++ {
		words := []string{"GO"}
		g := game.NewGame(words)
		g.Word = "GO"
		g.GuessLetter('G')
		g.GuessLetter('O')

		stats.RecordGame(g, "easy")
	}

	if stats.CurrentStreak != 3 {
		t.Errorf("Expected CurrentStreak to be 3, got %d", stats.CurrentStreak)
	}

	if stats.LongestStreak != 3 {
		t.Errorf("Expected LongestStreak to be 3, got %d", stats.LongestStreak)
	}

	// Lose a game
	words := []string{"GOLANG"}
	g := game.NewGame(words)
	g.Word = "GOLANG"
	wrongLetters := []rune{'X', 'Y', 'Z', 'Q', 'W', 'E'}
	for _, letter := range wrongLetters {
		g.GuessLetter(letter)
	}

	stats.RecordGame(g, "hard")

	if stats.CurrentStreak != 0 {
		t.Errorf("Expected CurrentStreak to be 0 after loss, got %d", stats.CurrentStreak)
	}

	if stats.LongestStreak != 3 {
		t.Errorf("Expected LongestStreak to remain 3, got %d", stats.LongestStreak)
	}
}

func TestGetWinRate(t *testing.T) {
	stats := game.NewStatistics()

	// No games played
	if stats.GetWinRate() != 0.0 {
		t.Errorf("Expected win rate to be 0.0 with no games, got %.2f", stats.GetWinRate())
	}

	// Simulate some games
	stats.GamesPlayed = 10
	stats.GamesWon = 7

	expectedRate := 70.0
	if stats.GetWinRate() != expectedRate {
		t.Errorf("Expected win rate to be %.1f, got %.1f", expectedRate, stats.GetWinRate())
	}
}

func TestGetAverageGuesses(t *testing.T) {
	stats := game.NewStatistics()

	// No games played
	if stats.GetAverageGuesses() != 0.0 {
		t.Errorf("Expected average guesses to be 0.0 with no games, got %.2f", stats.GetAverageGuesses())
	}

	// Simulate some games
	stats.GamesPlayed = 5
	stats.TotalGuesses = 25

	expectedAvg := 5.0
	if stats.GetAverageGuesses() != expectedAvg {
		t.Errorf("Expected average guesses to be %.1f, got %.1f", expectedAvg, stats.GetAverageGuesses())
	}
}

func TestGetGuessAccuracy(t *testing.T) {
	stats := game.NewStatistics()

	// No guesses made
	if stats.GetGuessAccuracy() != 0.0 {
		t.Errorf("Expected guess accuracy to be 0.0 with no guesses, got %.2f", stats.GetGuessAccuracy())
	}

	// Simulate some guesses
	stats.TotalGuesses = 20
	stats.CorrectGuesses = 15

	expectedAccuracy := 75.0
	if stats.GetGuessAccuracy() != expectedAccuracy {
		t.Errorf("Expected guess accuracy to be %.1f, got %.1f", expectedAccuracy, stats.GetGuessAccuracy())
	}
}

func TestRecentWordsLimit(t *testing.T) {
	stats := game.NewStatistics()

	// Add more than 10 words
	for i := 0; i < 15; i++ {
		words := []string{"WORD"}
		g := game.NewGame(words)
		g.Word = "WORD"
		stats.RecordGame(g, "easy")
	}

	// Should only keep last 10 words
	if len(stats.WordsGuessed) != 10 {
		t.Errorf("Expected WordsGuessed to have 10 items, got %d", len(stats.WordsGuessed))
	}
}

func TestResetStatistics(t *testing.T) {
	stats := game.NewStatistics()

	// Add some data
	stats.GamesPlayed = 10
	stats.GamesWon = 5
	stats.CurrentStreak = 3
	stats.LastPlayed = time.Now()
	stats.Difficulties["easy"] = 5
	stats.WordsGuessed = []string{"WORD1", "WORD2"}

	// Reset
	stats.ResetStatistics()

	// Verify reset
	if stats.GamesPlayed != 0 {
		t.Errorf("Expected GamesPlayed to be 0 after reset, got %d", stats.GamesPlayed)
	}

	if stats.GamesWon != 0 {
		t.Errorf("Expected GamesWon to be 0 after reset, got %d", stats.GamesWon)
	}

	if stats.CurrentStreak != 0 {
		t.Errorf("Expected CurrentStreak to be 0 after reset, got %d", stats.CurrentStreak)
	}

	if len(stats.Difficulties) != 0 {
		t.Errorf("Expected Difficulties to be empty after reset, got %d items", len(stats.Difficulties))
	}

	if len(stats.WordsGuessed) != 0 {
		t.Errorf("Expected WordsGuessed to be empty after reset, got %d items", len(stats.WordsGuessed))
	}
}
