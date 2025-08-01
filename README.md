# Hangman Game in Go 🎮

[![CI/CD Pipeline](https://github.com/[your-username]/hangman-go/actions/workflows/ci.yml/badge.svg)](https://github.com/[your-username]/hangman-go/actions/workflows/ci.yml)
[![Code Quality](https://github.com/[your-username]/hangman-go/actions/workflows/quality.yml/badge.svg)](https://github.com/[your-username]/hangman-go/actions/workflows/quality.yml)
[![codecov](https://codecov.io/gh/[your-username]/hangman-go/branch/main/graph/badge.svg)](https://codecov.io/gh/[your-username]/hangman-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/[your-username]/hangman-go)](https://goreportcard.com/report/github.com/[your-username]/hangman-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker Pulls](https://img.shields.io/docker/pulls/[your-username]/hangman-go)](https://hub.docker.com/r/[your-username]/hangman-go)

A terminal-based Hangman game built with Go to practice fundamental programming concepts and game logic implementation.

## 🎯 Project Overview

This project is a classic Hangman word-guessing game implemented in Go, designed to be played directly in the terminal. It serves as an excellent learning project for Go fundamentals including data structures, control flow, and project organization.

## ✨ Features

### Current Features (MVP)
- 🎲 Random word selection from predefined word list
- 🔤 Letter-by-letter guessing with validation
- 🎨 ASCII art hangman visualization
- 🏆 Win/loss detection and feedback
- 🛡️ Input validation (duplicate prevention, case insensitive)
- 🎮 Clean terminal-based user interface

### Planned Features
- 📊 Multiple difficulty levels (Easy, Medium, Hard)
- 🏅 Scoring system with high score tracking
- 📚 Word categories (Animals, Countries, Programming, etc.)
- 💾 Save/load game functionality
- 📈 Session statistics and game history
- 🌈 Colored terminal output
- 💡 Optional hints system

## 🚀 Getting Started

### Prerequisites
- Go 1.20 or higher
- Terminal/Command prompt

### Installation

#### Option 1: Download Binary (Recommended)
1. Go to the [Releases](https://github.com/[your-username]/hangman-go/releases) page
2. Download the binary for your operating system:
   - `hangman-linux-amd64.tar.gz` for Linux (64-bit)
   - `hangman-windows-amd64.exe.zip` for Windows (64-bit)
   - `hangman-macos-amd64.tar.gz` for macOS (Intel)
   - `hangman-macos-arm64.tar.gz` for macOS (Apple Silicon)
3. Extract and run:
```bash
# Linux/macOS
tar -xzf hangman-*.tar.gz
chmod +x hangman-*
./hangman-*

# Windows
# Extract the zip file and run hangman-windows-amd64.exe
```

#### Option 2: Install with Go
```bash
go install github.com/[your-username]/hangman-go@latest
hangman-go
```

#### Option 3: Clone and Build
```bash
git clone https://github.com/[your-username]/hangman-go.git
cd hangman-go
go mod download
go build -o hangman main.go
./hangman
```

#### Option 4: Docker
```bash
# Run directly
docker run -it --rm [your-username]/hangman-go:latest

# Or use docker-compose for development
docker-compose up hangman
```

### Building from Source
```bash
git clone https://github.com/[your-username]/hangman-go.git
cd hangman-go
go mod download
go build -ldflags="-s -w" -o hangman main.go
```

## 🏗️ Project Structure

```
hangman/
├── main.go                 # Entry point of the application
├── game/
│   ├── game.go            # Core game logic and structures
│   ├── word.go            # Word management and selection
│   └── display.go         # Game display and UI functions
├── utils/
│   ├── input.go           # User input handling utilities
│   └── validation.go      # Input validation functions
├── data/
│   └── words.txt          # Word list file
├── assets/
│   └── hangman_art.go     # ASCII art for hangman stages
├── tests/
│   ├── game_test.go       # Unit tests for game logic
│   └── word_test.go       # Unit tests for word functions
├── Plan.md                # Detailed project plan and roadmap
├── README.md              # This file
├── go.mod                 # Go module file
└── .gitignore             # Git ignore file
```

## 🎮 How to Play

1. Start the game by running the executable
2. A random word will be selected and displayed as underscores
3. Guess letters one at a time
4. Correct guesses reveal the letter's position(s) in the word
5. Wrong guesses add a part to the hangman drawing
6. Win by guessing the complete word before the drawing is finished
7. Lose if the hangman drawing is completed (6 wrong guesses)

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run specific package tests:
```bash
go test ./game
go test ./utils
```

## 🛠️ Development

### Key Go Concepts Practiced
- **Basic Concepts**: Variables, functions, control flow, data types
- **Data Structures**: Slices, maps, structs
- **Advanced Concepts**: Packages, interfaces, error handling, file I/O
- **Testing**: Unit tests and test-driven development
- **Code Organization**: Modular design and separation of concerns

### Development Timeline
- **Week 1**: Basic game mechanics and core logic ✅
- **Week 2**: User interface and input handling 🚧
- **Week 3**: Advanced features and testing 📋
- **Week 4**: Polish, optimization, and documentation 📋

## 🤝 Contributing

Contributions are welcome! Feel free to:
- 🐛 Report bugs
- 💡 Suggest new features
- 🔧 Submit pull requests
- 📖 Improve documentation

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built as a learning project to practice Go fundamentals
- Inspired by the classic Hangman word game
- ASCII art hangman drawings for visual feedback

## 📚 Learning Resources

- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)

---

**Happy Coding! 🎉**

> This project is perfect for Go beginners looking to practice fundamental concepts while building something fun and interactive.
