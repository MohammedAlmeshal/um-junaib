package main

import (
	_ "embed"
	"fmt"
	"os"

	"golang.org/x/term"
)

const (
	// Display glyphs
	snakeBodyGlyph = "\033[38;2;139;153;601m ◇ \033[0m"
	snakeHeadGlyph = "\033[38;2;139;153;601m o \033[0m"
	foodGlyph      = " 🍓"
	gridGlyph      = "\033[2;90m ◦ \033[0m"

	// Border characters
	horizontalBorder  = "───"
	verticalBorder    = "│"
	topLeftCorner     = " ┌"
	topRightCorner    = "┐"
	bottomLeftCorner  = " └"
	bottomRightCorner = "┘"
)

// Terminal utilities
func getTerminalSize() (width, height int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Default fallback if we can't get terminal size
		return 80, 24
	}
	return width, height
}

func calculateBoardIndent(boardWidth int) string {
	terminalWidth, _ := getTerminalSize()

	// Each cell is 3 characters wide (from your glyphs)
	// Plus 2 for left and right borders
	totalBoardWidth := (boardWidth * 3) + 2

	if totalBoardWidth >= terminalWidth {
		return "" // No indent if board is too wide
	}

	indent := ((terminalWidth - totalBoardWidth) / 2)
	if indent < 0 {
		indent = 0
	}

	return fmt.Sprintf("%*s", indent, "")
}

// Terminal control functions
func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

// Board rendering functions
func renderTopBorder(cols int) {
	indent := calculateBoardIndent(cols)
	fmt.Print(indent + topLeftCorner)
	for i := 0; i < cols; i++ {
		fmt.Print(horizontalBorder)
	}
	fmt.Println(topRightCorner + "\r ")
}

func renderBottomBorder(cols int) {
	indent := calculateBoardIndent(cols)
	fmt.Print("\r\n" + indent + bottomLeftCorner)
	for i := 0; i < cols; i++ {
		fmt.Print(horizontalBorder)
	}
	fmt.Println(bottomRightCorner)
}

func getCellGlyph(coord Coord, snake *Snake, food *Food) string {
	// Check if snake occupies this position
	if _, occupied := snake.occupied[coord]; occupied {
		if coord == snake.head() {
			return snakeHeadGlyph
		}
		return snakeBodyGlyph
	}

	// Check if food is at this position
	if food.coord.x == coord.x && food.coord.y == coord.y {
		return foodGlyph
	}

	// Empty cell
	return gridGlyph
}

func renderBoard(game *Game) {
	board := game.board
	snake := game.snake
	food := game.food
	cols := len(board[0])
	indent := calculateBoardIndent(cols)

	clearScreen()
	hideCursor()

	fmt.Printf(indent+"Score: %d \r\n", game.getScore())

	// Render top border
	renderTopBorder(cols)

	// Render board rows
	for row := range board {
		if row != 0 {
			fmt.Print("\r\n ")
		}

		fmt.Print(indent + verticalBorder) // Left border

		// Render each cell in the row
		for col := range board[row] {
			coord := Coord{row, col}
			glyph := getCellGlyph(coord, snake, food)
			fmt.Print(glyph)
		}

		fmt.Print(verticalBorder) // Right border
	}

	// Render bottom border
	renderBottomBorder(cols)
}

func showStartMenu() {
	clearScreen()
	fmt.Print("\r\nSTART 🐍\r\n")
	fmt.Print("Press 's' to start or 'q' to quit: ")
}

func showGameOver() {
	clearScreen()
	fmt.Print("\r\nGAME OVER!\r\n")
	fmt.Print("Press 'r' to restart or 'q' to quit: ")
}

func showWinScreen() {
	clearScreen()
	fmt.Print("\r\nYOU WON! 🎉\r\n")
	fmt.Print("Press 'r' to restart or 'q' to quit: ")
}
