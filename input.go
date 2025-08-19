package main

import (
	"os"

	"golang.org/x/term"
)

func inputReader(inputChan chan<- []byte) {
	fd := int(os.Stdin.Fd())
	oldState, _ := term.MakeRaw(fd)
	defer term.Restore(fd, oldState)

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
