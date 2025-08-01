package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ClearScreen clears the terminal screen
func ClearScreen() {
	var cmd *exec.Cmd
	
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// DisplayWelcome shows the welcome message and game instructions
func DisplayWelcome() {
	fmt.Println("🎮 WELCOME TO HANGMAN GAME! 🎮")
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("📖 HOW TO PLAY:")
	fmt.Println("• Guess the hidden word letter by letter")
	fmt.Println("• You have 6 wrong guesses before you lose")
	fmt.Println("• Enter one letter at a time")
	fmt.Println("• Good luck!")
	fmt.Println()
	fmt.Println("Press Enter to start...")
}

// DisplayGameState shows the current state of the game
func DisplayGameState(g *Game) {
	fmt.Println("🎯 HANGMAN GAME")
	fmt.Println("===============")
	fmt.Println()
	
	// Display hangman figure
	DisplayHangman(g.WrongGuesses)
	fmt.Println()
	
	// Display word progress
	fmt.Printf("Word: %s\n", g.GetDisplayWord())
	fmt.Println()
	
	// Display game statistics
	fmt.Printf("Wrong guesses: %d/%d\n", g.WrongGuesses, g.MaxWrongGuesses)
	fmt.Printf("Remaining guesses: %d\n", g.GetRemainingGuesses())
	fmt.Println()
	
	// Display guessed letters
	wrongLetters := g.GetWrongLetters()
	if len(wrongLetters) > 0 {
		fmt.Printf("Wrong letters: %s\n", formatLetters(wrongLetters))
	}
	
	guessedLetters := g.GetGuessedLetters()
	if len(guessedLetters) > 0 {
		fmt.Printf("All guessed letters: %s\n", formatLetters(guessedLetters))
	}
	fmt.Println()
}

// DisplayHangman shows the hangman figure based on wrong guesses
func DisplayHangman(wrongGuesses int) {
	stages := []string{
		// Stage 0: Empty gallows
		`
   +---+
   |   |
       |
       |
       |
       |
=========`,
		// Stage 1: Head
		`
   +---+
   |   |
   O   |
       |
       |
       |
=========`,
		// Stage 2: Body
		`
   +---+
   |   |
   O   |
   |   |
       |
       |
=========`,
		// Stage 3: Left arm
		`
   +---+
   |   |
   O   |
  /|   |
       |
       |
=========`,
		// Stage 4: Right arm
		`
   +---+
   |   |
   O   |
  /|\  |
       |
       |
=========`,
		// Stage 5: Left leg
		`
   +---+
   |   |
   O   |
  /|\  |
  /    |
       |
=========`,
		// Stage 6: Right leg (complete hangman)
		`
   +---+
   |   |
   O   |
  /|\  |
  / \  |
       |
=========`,
	}
	
	if wrongGuesses >= 0 && wrongGuesses < len(stages) {
		fmt.Print(stages[wrongGuesses])
	}
}

// DisplayWinMessage shows the win message
func DisplayWinMessage(word string) {
	fmt.Println("🎉 CONGRATULATIONS! 🎉")
	fmt.Println("======================")
	fmt.Printf("You guessed the word: %s\n", word)
	fmt.Println("You win! 🏆")
	fmt.Println()
}

// DisplayLoseMessage shows the lose message
func DisplayLoseMessage(word string) {
	fmt.Println("💀 GAME OVER 💀")
	fmt.Println("===============")
	fmt.Printf("The word was: %s\n", word)
	fmt.Println("Better luck next time! 😔")
	fmt.Println()
}

// DisplayInvalidInput shows invalid input message
func DisplayInvalidInput(message string) {
	fmt.Printf("❌ Invalid input: %s\n", message)
	fmt.Println("Please enter a single letter (A-Z)")
	fmt.Println()
}

// DisplayAlreadyGuessed shows already guessed message
func DisplayAlreadyGuessed(letter rune) {
	fmt.Printf("⚠️  You already guessed '%c'!\n", letter)
	fmt.Println("Try a different letter.")
	fmt.Println()
}

// DisplayCorrectGuess shows correct guess message
func DisplayCorrectGuess(letter rune) {
	fmt.Printf("✅ Great! '%c' is in the word!\n", letter)
	fmt.Println()
}

// DisplayWrongGuess shows wrong guess message
func DisplayWrongGuess(letter rune) {
	fmt.Printf("❌ Sorry, '%c' is not in the word.\n", letter)
	fmt.Println()
}

// formatLetters formats a slice of runes as a comma-separated string
func formatLetters(letters []rune) string {
	if len(letters) == 0 {
		return ""
	}
	
	var strs []string
	for _, letter := range letters {
		strs = append(strs, string(letter))
	}
	
	return strings.Join(strs, ", ")
}

// DisplayGameStats shows game statistics
func DisplayGameStats(gamesPlayed, gamesWon int) {
	fmt.Println("📊 GAME STATISTICS")
	fmt.Println("==================")
	fmt.Printf("Games played: %d\n", gamesPlayed)
	fmt.Printf("Games won: %d\n", gamesWon)
	if gamesPlayed > 0 {
		winRate := float64(gamesWon) / float64(gamesPlayed) * 100
		fmt.Printf("Win rate: %.1f%%\n", winRate)
	}
	fmt.Println()
}
