package main

import (
	"fmt"
	"log"
	"os"

	"hangman-go/assets"
	"hangman-go/game"
	"hangman-go/utils"
)

func main() {
	// Display title
	fmt.Print(assets.GameTitle())
	fmt.Println("\nWelcome to Hangman!")
	fmt.Println("===================")
	
	// Load words
	wordList, err := loadWords()
	if err != nil {
		log.Printf("Warning: Could not load words from file: %v", err)
		log.Println("Using default word list instead.")
		wordList = game.GetDefaultWords()
	}
	
	// Game statistics
	gamesPlayed := 0
	gamesWon := 0
	
	// Main game loop
	for {
		// Get difficulty level
		difficulty, err := getDifficulty()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		
		// Get words for selected difficulty
		words := wordList.GetWordsByDifficulty(difficulty)
		if len(words) == 0 {
			fmt.Println("No words available for selected difficulty. Using all words.")
			words = wordList.Words
		}
		
		// Start new game
		hangmanGame := game.NewGame(words)
		gamesPlayed++
		
		// Play the game
		if playGame(hangmanGame) {
			gamesWon++
		}
		
		// Show statistics
		game.DisplayGameStats(gamesPlayed, gamesWon)
		
		// Ask if player wants to play again
		playAgain, err := utils.GetYesNoInput("Do you want to play again?")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		
		if !playAgain {
			break
		}
		
		game.ClearScreen()
	}
	
	// Farewell message
	fmt.Println("Thanks for playing Hangman! 👋")
	fmt.Printf("Final Statistics - Games Played: %d, Games Won: %d\n", gamesPlayed, gamesWon)
	if gamesPlayed > 0 {
		winRate := float64(gamesWon) / float64(gamesPlayed) * 100
		fmt.Printf("Win Rate: %.1f%%\n", winRate)
	}
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
	fmt.Printf("\nStarting new game with difficulty level!\n")
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
			game.DisplayCorrectGuess(letter)
		} else {
			game.DisplayWrongGuess(letter)
		}
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
			game.DisplayInvalidInput(err.Error())
			continue
		}
		
		// Check if letter was already guessed
		if g.GuessedLetters[letter] {
			game.DisplayAlreadyGuessed(letter)
			continue
		}
		
		return letter, nil
	}
}
