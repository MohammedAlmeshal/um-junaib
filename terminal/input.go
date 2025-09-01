package terminal

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
	"snake/game"
)

var terminalState *term.State

func SetupTerminal() {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return
	}
	terminalState = oldState

	// Handle Ctrl+C and other signals to restore terminal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		RestoreTerminal()
		os.Exit(0)
	}()
}

func RestoreTerminal() {
	if terminalState != nil {
		fd := int(os.Stdin.Fd())
		term.Restore(fd, terminalState)
	}
}

func inputReader(inputChan chan<- []byte) {
	buf := make([]byte, 10)
	for {
		n, err := os.Stdin.Read(buf)

		if err != nil {
			close(inputChan)
			return
		}

		if n > 0 {
			data := make([]byte, n)
			copy(data, buf[:n])
			inputChan <- data

			// Exit on ESC
			if n == 1 && buf[0] == 27 {
				close(inputChan)
				return
			}

		}
	}
}

func StartInputReader() <-chan []byte {
	SetupTerminal()
	ch := make(chan []byte, 5)
	go inputReader(ch)
	return ch
}

func InputHandler(data []byte) (game.Direction, bool) {
	if len(data) >= 3 && data[0] == 27 && data[1] == 91 {
		switch data[2] {
		case 65: // A
			return game.UP, true
		case 66: // B
			return game.DOWN, true
		case 67: // C
			return game.RIGHT, true
		case 68: // D
			return game.LEFT, true
		default:
			return 0, false
		}
	}
	return 0, false
}
