package main

import (
	"os"
	"snake/game"
	"snake/terminal"
)

func main() {
	defer terminal.RestoreTerminal()
	inputChan := terminal.StartInputReader()

	// Show start menu
	if result := terminal.StartMenuLoop(inputChan); result == terminal.MenuQuit {
		terminal.RestoreTerminal()
		os.Exit(0)
	}

	// Main game loop
	for {
		g := game.NewGame()
		status := g.Run(inputChan, terminal.RenderBoard, terminal.InputHandler)

		switch status {
		case game.GameDead:
			result := terminal.GameOverMenuLoop(inputChan)
			if result == terminal.MenuQuit {
				terminal.RestoreTerminal()
				os.Exit(0)
			}
		case game.GameWon:
			result := terminal.WinMenuLoop(inputChan)
			if result == terminal.MenuQuit {
				terminal.RestoreTerminal()
				os.Exit(0)
			}
		case game.GameTerminated:
			return
		}
	}
}
