package main

import (
	"fmt"
	"os"
)

func handleMenuInput(inputChan <-chan []byte) bool {
	for {
		data := <-inputChan

		if len(data) == 0 {
			continue
		}

		switch data[0] {
		case 's', 'S':
			return true // Start game
		case 'q', 'Q':
			fmt.Print("QUIT")
			os.Exit(0)
		default:
			// Invalid input, show menu again
			showStartMenu()
		}
	}
}

func handleGameOver(inputChan <-chan []byte) bool {
	showGameOver()

	for {
		data := <-inputChan

		if len(data) == 0 {
			continue
		}

		switch data[0] {
		case 'r', 'R':
			return true // Restart
		case 'q', 'Q':
			fmt.Print("QUIT")
			os.Exit(0)
		}
	}
}

func main() {
	inputChan := startInputReader()

	// Show start menu
	showStartMenu()
	if !handleMenuInput(inputChan) {
		return
	}

	// Main game loop
	for {
		game := NewGame()
		status := Run(game, inputChan)

		switch status {
		case GameDead:
			if !handleGameOver(inputChan) {
				return
			}
		case GameTerminated:
			fmt.Print("QUIT")
			return
		}
	}
}
