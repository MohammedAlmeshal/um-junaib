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

func waitForInput(inputChan <-chan []byte) ([]byte, bool) {
	for {
		data, ok := <-inputChan
		if !ok {
			return nil, false
		}
		if len(data) > 0 {
			return data, true
		}
	}
}

func StartMenuLoop(inputChan <-chan []byte) MenuResult {
	ShowStartMenu()
	for {
		data, ok := waitForInput(inputChan)
		if !ok {
			RestoreTerminal()
			os.Exit(0)
		}
		
		// Check for ESC key
		if len(data) == 1 && data[0] == 27 {
			return MenuQuit
		}
		
		// Check for 'q' or 'Q'
		if len(data) > 0 && (data[0] == 'q' || data[0] == 'Q') {
			return MenuQuit
		}
		
		// Any other key starts the game
		return MenuStart
	}
}

func GameOverMenuLoop(inputChan <-chan []byte, score int) MenuResult {
	ShowGameOver(score)
	for {
		data, ok := waitForInput(inputChan)
		if !ok {
			RestoreTerminal()
			os.Exit(0)
		}
		
		// Check for ESC key
		if len(data) == 1 && data[0] == 27 {
			return MenuQuit
		}
		
		// Check for 'q' or 'Q'
		if len(data) > 0 && (data[0] == 'q' || data[0] == 'Q') {
			return MenuQuit
		}
		
		// Any other key restarts the game
		return MenuRestart
	}
}

func WinMenuLoop(inputChan <-chan []byte) MenuResult {
	ShowWinScreen()
	for {
		data, ok := waitForInput(inputChan)
		if !ok {
			RestoreTerminal()
			os.Exit(0)
		}
		
		// Check for ESC key
		if len(data) == 1 && data[0] == 27 {
			return MenuQuit
		}
		
		// Check for 'q' or 'Q'
		if len(data) > 0 && (data[0] == 'q' || data[0] == 'Q') {
			return MenuQuit
		}
		
		// Any other key restarts the game
		return MenuRestart
	}
}
