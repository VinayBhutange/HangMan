package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/VinayBhutange/hangman-go/assets"
	"github.com/VinayBhutange/hangman-go/game"
	"github.com/VinayBhutange/hangman-go/utils"
)

func main() {
	// Display title
	fmt.Print(assets.GameTitle())
	fmt.Println(utils.Bold("\nWelcome to Hangman!"))
	fmt.Println("===================")
	
	// Load statistics
	stats, err := game.LoadStatistics()
	if err != nil {
		log.Printf("Warning: Could not load statistics: %v", err)
		stats = game.NewStatistics()
	}
	
	// Load words
	wordList, err := loadWords()
	if err != nil {
		log.Printf("Warning: Could not load words from file: %v", err)
		log.Println("Using default word list instead.")
		wordList = game.GetDefaultWords()
	}
	
	// Show welcome message with statistics
	if stats.GamesPlayed > 0 {
		fmt.Printf("Welcome back! You've played %d games with a %.1f%% win rate.\n\n", 
			stats.GamesPlayed, stats.GetWinRate())
	}
	
	// Main menu loop
	for {
		choice := showMainMenu()
		
		switch choice {
		case "1":
			// Play game
			playHangmanGame(wordList, stats)
		case "2":
			// View statistics
			stats.PrintStatistics()
			utils.WaitForEnter()
		case "3":
			// Settings/Options
			showSettingsMenu(wordList, stats)
		case "4":
			// Exit
			fmt.Println(utils.Info("Thanks for playing Hangman! 👋"))
			printFinalStats(stats)
			return
		default:
			fmt.Println(utils.Error("Invalid choice. Please try again."))
		}
		
		game.ClearScreen()
	}
}

// showMainMenu displays the main menu and returns user choice
func showMainMenu() string {
	fmt.Println(utils.Bold("🎮 MAIN MENU"))
	fmt.Println("=============")
	fmt.Println("1. 🎯 Play Hangman")
	fmt.Println("2. 📊 View Statistics")
	fmt.Println("3. ⚙️  Settings")
	fmt.Println("4. 🚪 Exit")
	fmt.Println()
	
	choice, err := utils.GetUserInput("Enter your choice (1-4): ")
	if err != nil {
		fmt.Println(utils.Error("Error reading input: " + err.Error()))
		return ""
	}
	
	return strings.TrimSpace(choice)
}

// showSettingsMenu displays the settings menu
func showSettingsMenu(wordList *game.WordList, stats *game.Statistics) {
	for {
		fmt.Println(utils.Bold("⚙️ SETTINGS"))
		fmt.Println("============")
		fmt.Println("1. 📝 Add Custom Word")
		fmt.Println("2. 🗑️  Remove Word")
		fmt.Println("3. 📋 List All Words")
		fmt.Println("4. 🔄 Reset Statistics")
		fmt.Println("5. 🔙 Back to Main Menu")
		fmt.Println()
		
		choice, err := utils.GetUserInput("Enter your choice (1-5): ")
		if err != nil {
			fmt.Println(utils.Error("Error reading input: " + err.Error()))
			continue
		}
		
		switch strings.TrimSpace(choice) {
		case "1":
			addCustomWord(wordList)
		case "2":
			removeWord(wordList)
		case "3":
			listWords(wordList)
		case "4":
			resetStatistics(stats)
		case "5":
			return
		default:
			fmt.Println(utils.Error("Invalid choice. Please try again."))
		}
		
		fmt.Println()
	}
}

// playHangmanGame plays a single game session
func playHangmanGame(wordList *game.WordList, stats *game.Statistics) {
	// Get difficulty level
	difficulty, err := getDifficulty()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	// Get words for selected difficulty
	words := wordList.GetWordsByDifficulty(difficulty)
	if len(words) == 0 {
		fmt.Println(utils.Warning("No words available for selected difficulty. Using all words."))
		words = wordList.Words
	}
	
	// Start new game
	hangmanGame := game.NewGame(words)
	
	// Play the game
	won := playGame(hangmanGame)
	
	// Record statistics
	stats.RecordGame(hangmanGame, difficulty)
	
	// Save statistics
	if err := stats.SaveStatistics(); err != nil {
		log.Printf("Warning: Could not save statistics: %v", err)
	}
	
	// Show brief stats
	if won {
		fmt.Printf(utils.Success("Game won! Current streak: %d\n"), stats.CurrentStreak)
	} else {
		fmt.Printf(utils.Error("Game lost. Win rate: %.1f%%\n"), stats.GetWinRate())
	}
	
	utils.WaitForEnter()
}

// loadWords attempts to load words from file, returns error if unsuccessful
func loadWords() (*game.WordList, error) {
	// Try to load from data/words.txt first
	if _, err := os.Stat("data/words.txt"); err == nil {
		return game.LoadWordsFromFile("data/words.txt")
	}
	
	// If that fails, try relative path
	if _, err := os.Stat("./data/words.txt"); err == nil {
		return game.LoadWordsFromFile("./data/words.txt")
	}
	
	return nil, fmt.Errorf("word file not found")
}

// getDifficulty gets the difficulty level from the user
func getDifficulty() (string, error) {
	for {
		difficulty, err := utils.GetDifficultyInput()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		return difficulty, nil
	}
}

// playGame plays a single game and returns true if player won
func playGame(g *game.Game) bool {
	fmt.Printf(utils.Info("Starting new game!\n"))
	fmt.Printf("The word has %d letters.\n\n", len(g.Word))
	
	// Game loop
	for !g.IsGameOver {
		// Display current game state
		game.DisplayGameState(g)
		
		// Get letter input from user
		letter, err := getLetterInput(g)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		
		// Process the guess
		isCorrect := g.GuessLetter(letter)
		
		// Clear screen and show feedback
		game.ClearScreen()
		
		if isCorrect {
			fmt.Println(utils.Success(fmt.Sprintf("Great! '%c' is in the word!", letter)))
		} else {
			fmt.Println(utils.Error(fmt.Sprintf("Sorry, '%c' is not in the word.", letter)))
		}
		fmt.Println()
	}
	
	// Display final game state
	game.DisplayGameState(g)
	
	// Show win/lose message
	if g.IsWon {
		fmt.Print(assets.WinMessage())
		game.DisplayWinMessage(g.Word)
		return true
	} else {
		fmt.Print(assets.LoseMessage())
		game.DisplayLoseMessage(g.Word)
		return false
	}
}

// getLetterInput gets a valid letter input from the user
func getLetterInput(g *game.Game) (rune, error) {
	for {
		letter, err := utils.GetLetterInput()
		if err != nil {
			fmt.Println(utils.Error("Invalid input: " + err.Error()))
			fmt.Println("Please enter a single letter (A-Z)")
			continue
		}
		
		// Check if letter was already guessed
		if g.GuessedLetters[letter] {
			fmt.Println(utils.Warning(fmt.Sprintf("You already guessed '%c'! Try a different letter.", letter)))
			continue
		}
		
		return letter, nil
	}
}

// addCustomWord adds a custom word to the word list
func addCustomWord(wordList *game.WordList) {
	word, err := utils.GetUserInput("Enter a word to add (3+ letters): ")
	if err != nil {
		fmt.Println(utils.Error("Error reading input: " + err.Error()))
		return
	}
	
	if !utils.IsValidWord(word) {
		fmt.Println(utils.Error("Invalid word. Word must be 3+ letters and contain only alphabetic characters."))
		return
	}
	
	wordList.AddWord(word)
	fmt.Println(utils.Success(fmt.Sprintf("Added '%s' to word list!", strings.ToUpper(word))))
}

// removeWord removes a word from the word list
func removeWord(wordList *game.WordList) {
	word, err := utils.GetUserInput("Enter a word to remove: ")
	if err != nil {
		fmt.Println(utils.Error("Error reading input: " + err.Error()))
		return
	}
	
	originalCount := wordList.GetWordCount()
	wordList.RemoveWord(word)
	
	if wordList.GetWordCount() < originalCount {
		fmt.Println(utils.Success(fmt.Sprintf("Removed '%s' from word list!", strings.ToUpper(word))))
	} else {
		fmt.Println(utils.Warning(fmt.Sprintf("Word '%s' not found in word list.", strings.ToUpper(word))))
	}
}

// listWords displays all words in the word list
func listWords(wordList *game.WordList) {
	fmt.Printf(utils.Bold("📋 WORD LIST (%d words)\n"), wordList.GetWordCount())
	fmt.Println("===============")
	
	words := wordList.Words
	for i, word := range words {
		fmt.Printf("%3d. %s\n", i+1, word)
		if (i+1)%20 == 0 && i < len(words)-1 {
			more, err := utils.GetYesNoInput("Show more words?")
			if err != nil || !more {
				break
			}
		}
	}
	
	fmt.Println()
	utils.WaitForEnter()
}

// resetStatistics resets all game statistics
func resetStatistics(stats *game.Statistics) {
	confirm, err := utils.GetYesNoInput("Are you sure you want to reset all statistics? This cannot be undone.")
	if err != nil {
		fmt.Println(utils.Error("Error reading input: " + err.Error()))
		return
	}
	
	if confirm {
		stats.ResetStatistics()
		if err := stats.SaveStatistics(); err != nil {
			fmt.Println(utils.Error("Error saving statistics: " + err.Error()))
		} else {
			fmt.Println(utils.Success("Statistics reset successfully!"))
		}
	} else {
		fmt.Println(utils.Info("Statistics reset cancelled."))
	}
}

// printFinalStats prints final statistics when exiting
func printFinalStats(stats *game.Statistics) {
	if stats.GamesPlayed > 0 {
		fmt.Printf("Final Statistics - Games: %d, Won: %d, Win Rate: %.1f%%\n", 
			stats.GamesPlayed, stats.GamesWon, stats.GetWinRate())
		if stats.LongestStreak > 0 {
			fmt.Printf("Your best winning streak was: %d games!\n", stats.LongestStreak)
		}
	}
}
