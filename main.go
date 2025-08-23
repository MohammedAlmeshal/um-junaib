package main

import (
	"fmt"
	"os"
)

func waitForInput(inputChan <-chan []byte) (byte, bool) {
	for {
		data, ok := <-inputChan
		if !ok {
			return 0, false
		}
		if len(data) > 0 {
			return data[0], true
		}
	}
}

func startMenu(inputChan <-chan []byte) bool {
	showStartMenu()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			restoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 's', 'S':
			return true
		case 'q', 'Q':
			restoreTerminal()
			os.Exit(0)
		default:
			showStartMenu()
		}
	}
}

func gameOverMenu(inputChan <-chan []byte) bool {
	showGameOver()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			restoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 'r', 'R':
			return true
		case 'q', 'Q':
			fmt.Print("QUIT")
			restoreTerminal()
			os.Exit(0)
		default:
			showGameOver()
		}
	}
}

func winMenu(inputChan <-chan []byte) bool {
	showWinScreen()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			restoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 'r', 'R':
			return true
		case 'q', 'Q':
			restoreTerminal()
			os.Exit(0)
		default:
			showWinScreen()
		}
	}
}

func main() {
	defer restoreTerminal()
	inputChan := startInputReader()

	// Show start menu
	if start := startMenu(inputChan); !start {
		return
	}

	// Main game loop
	for {
		game := NewGame()
		status := Run(game, inputChan)

		switch status {
		case GameDead:
			if restart := gameOverMenu(inputChan); !restart {
				return
			}
		case GameWon:
			if restart := winMenu(inputChan); !restart {
				return
			}
		case GameTerminated:
			return
		}
	}
}
