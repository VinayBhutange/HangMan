// Package utils provides utility functions for input handling, validation, and terminal colors.
package utils

import (
	"os"
	"runtime"
)

// Color codes for terminal output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

// ColorSupported checks if the terminal supports color output
func ColorSupported() bool {
	// Disable colors on Windows for now (can be enhanced later)
	if runtime.GOOS == "windows" {
		return false
	}

	// Check if we're in a terminal that likely supports colors
	term := os.Getenv("TERM")
	return term != "" && term != "dumb"
}

// Colorize adds color to text if colors are supported
func Colorize(text, color string) string {
	if !ColorSupported() {
		return text
	}
	return color + text + ColorReset
}

// Red returns red colored text
func Red(text string) string {
	return Colorize(text, ColorRed)
}

// Green returns green colored text
func Green(text string) string {
	return Colorize(text, ColorGreen)
}

// Yellow returns yellow colored text
func Yellow(text string) string {
	return Colorize(text, ColorYellow)
}

// Blue returns blue colored text
func Blue(text string) string {
	return Colorize(text, ColorBlue)
}

// Purple returns purple colored text
func Purple(text string) string {
	return Colorize(text, ColorPurple)
}

// Cyan returns cyan colored text
func Cyan(text string) string {
	return Colorize(text, ColorCyan)
}

// Bold returns bold text
func Bold(text string) string {
	return Colorize(text, ColorBold)
}

// Success returns green colored success message
func Success(text string) string {
	return Green("✅ " + text)
}

// Error returns red colored error message
func Error(text string) string {
	return Red("❌ " + text)
}

// Warning returns yellow colored warning message
func Warning(text string) string {
	return Yellow("⚠️ " + text)
}

// Info returns blue colored info message
func Info(text string) string {
	return Blue("ℹ️ " + text)
}
