package main

import (
	"os"
)

func main() {
	defer restoreTerminal()
	inputChan := startInputReader()

	// Show start menu
	if result := startMenuLoop(inputChan); result == MenuQuit {
		restoreTerminal()
		os.Exit(0)
	}

	// Main game loop
	for {
		game := newGame()
		status := game.run(inputChan)

		switch status {
		case GameDead:
			result := gameOverMenuLoop(inputChan)
			if result == MenuQuit {
				restoreTerminal()
				os.Exit(0)
			}
		case GameWon:
			result := winMenuLoop(inputChan)
			if result == MenuQuit {
				restoreTerminal()
				os.Exit(0)
			}
		case GameTerminated:
			return
		}
	}
}
