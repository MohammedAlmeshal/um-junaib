package main

import "fmt"

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func renderBoard(game *Game) {
	board := game.board
	snake := game.snake
	food := game.food
	// clear screan
	clearScreen()
	hideCursor()
	for row := range board {
		fmt.Print("\r\n")
		for col := range board[row] {

			_, occupied := snake.occupied[Coord{row, col}]

			if occupied {
				if (Coord{row, col} == snake.head()) {
					fmt.Printf(" 🐍")
				} else {
					fmt.Printf(" 💪")
				}

			} else {
				if food.coord.x == row && food.coord.y == col {
					fmt.Printf(" 🍒")
				} else {
					fmt.Printf(" . ")
				}

			}

		}
	}
	fmt.Print("\n")
}
