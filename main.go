package main

import (
	"os"
)

func main() {
	defer restoreTerminal()
	inputChan := startInputReader()

	// Show start menu
	if result := ShowStartMenu(inputChan); result == MenuQuit {
		restoreTerminal()
		os.Exit(0)
	}

	// Main game loop
	for {
		game := NewGame()
		status := Run(game, inputChan)

		switch status {
		case GameDead:
			result := ShowGameOverMenu(inputChan)
			if result == MenuQuit {
				restoreTerminal()
				os.Exit(0)
			}
		case GameWon:
			result := ShowWinMenu(inputChan)
			if result == MenuQuit {
				restoreTerminal()
				os.Exit(0)
			}
		case GameTerminated:
			return
		}
	}
}
