package terminal

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

func StartMenuLoop(inputChan <-chan []byte) MenuResult {
	ShowStartMenu()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			RestoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 's', 'S':
			return MenuStart
		case 'q', 'Q':
			return MenuQuit
		default:
			ShowStartMenu()
		}
	}
}

func GameOverMenuLoop(inputChan <-chan []byte) MenuResult {
	ShowGameOver()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			RestoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 'r', 'R':
			return MenuRestart
		case 'q', 'Q':
			return MenuQuit
		default:
			ShowGameOver()
		}
	}
}

func WinMenuLoop(inputChan <-chan []byte) MenuResult {
	ShowWinScreen()
	for {
		input, ok := waitForInput(inputChan)
		if !ok {
			RestoreTerminal()
			os.Exit(0)
		}
		switch input {
		case 'r', 'R':
			return MenuRestart
		case 'q', 'Q':
			return MenuQuit
		default:
			ShowWinScreen()
		}
	}
}
