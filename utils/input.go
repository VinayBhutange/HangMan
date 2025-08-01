package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	difficultyEasy   = "easy"
	difficultyMedium = "medium"
	difficultyHard   = "hard"
)

// GetUserInput reads a line of input from the user
func GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// GetLetterInput gets a single letter input from the user
func GetLetterInput() (rune, error) {
	input, err := GetUserInput("Enter a letter: ")
	if err != nil {
		return 0, err
	}

	// Validate input
	if len(input) != 1 {
		return 0, fmt.Errorf("please enter exactly one letter")
	}

	letter := rune(strings.ToUpper(input)[0])

	if !IsLetter(letter) {
		return 0, fmt.Errorf("please enter a valid letter (A-Z)")
	}

	return letter, nil
}

// GetYesNoInput gets a yes/no response from the user
func GetYesNoInput(prompt string) (bool, error) {
	input, err := GetUserInput(prompt + " (y/n): ")
	if err != nil {
		return false, err
	}

	input = strings.ToLower(input)

	switch input {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		return false, fmt.Errorf("please enter 'y' for yes or 'n' for no")
	}
}

// WaitForEnter waits for the user to press Enter
func WaitForEnter() {
	fmt.Print("Press Enter to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n') // Ignore error
}

// GetDifficultyInput gets difficulty level from user
func GetDifficultyInput() (string, error) {
	fmt.Println("Select difficulty level:")
	fmt.Println("1. Easy (4-5 letters)")
	fmt.Println("2. Medium (6-8 letters)")
	fmt.Println("3. Hard (9+ letters)")
	fmt.Println()

	input, err := GetUserInput("Enter your choice (1-3): ")
	if err != nil {
		return "", err
	}

	choice := strings.TrimSpace(input)
	switch choice {
	case "1":
		return difficultyEasy, nil
	case "2":
		return difficultyMedium, nil
	case "3":
		return difficultyHard, nil
	default:
		return "", fmt.Errorf("please enter 1, 2, or 3")
	}
}
