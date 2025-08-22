package main

import (
	"fmt"
	"os"
)

func waitForInput(inputChan <-chan []byte) byte {
	for {
		data := <-inputChan
		if len(data) > 0 {
			return data[0]
		}
	}
}

func startMenu(inputChan <-chan []byte) bool {
	showStartMenu()
	for {
		input := waitForInput(inputChan)
		switch input {
		case 's', 'S':
			return true
		case 'q', 'Q':
			os.Exit(0)
		default:
			showStartMenu()
		}
	}
}

func gameOverMenu(inputChan <-chan []byte) bool {
	showGameOver()
	for {
		input := waitForInput(inputChan)
		switch input {
		case 'r', 'R':
			return true
		case 'q', 'Q':
			fmt.Print("QUIT")
			os.Exit(0)
		default:
			showGameOver()
		}
	}
}

func winMenu(inputChan <-chan []byte) bool {
	showWinScreen()
	for {
		input := waitForInput(inputChan)
		switch input {
		case 'r', 'R':
			return true
		case 'q', 'Q':
			fmt.Print("QUIT")
			os.Exit(0)
		default:
			showWinScreen()
		}
	}
}

func main() {
	inputChan := startInputReader()

	// Show start menu
	if !startMenu(inputChan) {
		return
	}

	// Main game loop
	for {
		game := NewGame()
		status := Run(game, inputChan)

		switch status {
		case GameDead:
			if !gameOverMenu(inputChan) {
				return
			}
		case GameWon:
			if !winMenu(inputChan) {
				return
			}
		case GameTerminated:
			return
		}
	}
}
