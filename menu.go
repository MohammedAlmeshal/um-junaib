package main

import (
	"os"
)

type MenuResult int

const (
	MenuStart MenuResult = iota
	MenuRestart
	MenuQuit
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

func ShowStartMenu(inputChan <-chan []byte) MenuResult {
	showStartMenu()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			restoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 's', 'S':
			return MenuStart
		case 'q', 'Q':
			return MenuQuit
		default:
			showStartMenu()
		}
	}
}

func ShowGameOverMenu(inputChan <-chan []byte) MenuResult {
	showGameOver()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			restoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 'r', 'R':
			return MenuRestart
		case 'q', 'Q':
			return MenuQuit
		default:
			showGameOver()
		}
	}
}

func ShowWinMenu(inputChan <-chan []byte) MenuResult {
	showWinScreen()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			restoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 'r', 'R':
			return MenuRestart
		case 'q', 'Q':
			return MenuQuit
		default:
			showWinScreen()
		}
	}
}
