package main

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

var terminalState *term.State

func setupTerminal() {
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
		restoreTerminal()
		os.Exit(0)
	}()
}

func restoreTerminal() {
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

func startInputReader() <-chan []byte {
	setupTerminal()
	ch := make(chan []byte, 5)
	go inputReader(ch)
	return ch
}

func inputHandler(data []byte) (Direction, bool) {
	if len(data) >= 3 && data[0] == 27 && data[1] == 91 {
		switch data[2] {
		case 65: // A
			return UP, true
		case 66: // B
			return DOWN, true
		case 67: // C
			return RIGHT, true
		case 68: // D
			return LEFT, true
		default:
			return 0, false
		}
	}
	return 0, false
}
