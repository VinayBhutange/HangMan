// Package assets contains game art and visual elements for the hangman game.
package assets

// HangmanStages contains all the ASCII art stages for the hangman
var HangmanStages = []string{
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
	// Stage 6: Right leg (game over)
	`
   +---+
   |   |
   O   |
  /|\  |
  / \  |
       |
=========`,
}

// GetHangmanStage returns the hangman ASCII art for a given stage
func GetHangmanStage(stage int) string {
	if stage < 0 || stage >= len(HangmanStages) {
		return HangmanStages[0] // Return empty gallows for invalid stage
	}
	return HangmanStages[stage]
}

// GetMaxStages returns the maximum number of hangman stages
func GetMaxStages() int {
	return len(HangmanStages) - 1 // Subtract 1 because we start from 0
}

// GameTitle returns ASCII art for the game title
func GameTitle() string {
	return `
██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗
██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║
███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║
██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║
██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝
`
}

// WinMessage returns ASCII art for winning
func WinMessage() string {
	return `
██╗   ██╗ ██████╗ ██╗   ██╗    ██╗    ██╗██╗███╗   ██╗██╗
╚██╗ ██╔╝██╔═══██╗██║   ██║    ██║    ██║██║████╗  ██║██║
 ╚████╔╝ ██║   ██║██║   ██║    ██║ █╗ ██║██║██╔██╗ ██║██║
  ╚██╔╝  ██║   ██║██║   ██║    ██║███╗██║██║██║╚██╗██║╚═╝
   ██║   ╚██████╔╝╚██████╔╝    ╚███╔███╔╝██║██║ ╚████║██╗
   ╚═╝    ╚═════╝  ╚═════╝      ╚══╝╚══╝ ╚═╝╚═╝  ╚═══╝╚═╝
`
}

// LoseMessage returns ASCII art for losing
func LoseMessage() string {
	return `
 ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ██╗   ██╗███████╗██████╗ 
██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██║   ██║██╔════╝██╔══██╗
██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║██║   ██║█████╗  ██████╔╝
██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝ ╚████╔╝ ███████╗██║  ██║
 ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝
`
}
