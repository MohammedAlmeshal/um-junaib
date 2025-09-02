package terminal

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"snake/game"

	"golang.org/x/term"
)

const (
	// Display glyphs
	SnakeBodyGlyph = "\033[38;2;238;177;121m ❂ \033[0m"
	SnakeHeadGlyph = "\033[38;2;220;64;72m ❖ \033[0m"
	FoodGlyph      = " 🍒"
	GridGlyph      = "\033[2;90m ◦ \033[0m"

	// Border characters
	HorizontalBorder  = "───"
	VerticalBorder    = "│"
	TopLeftCorner     = " ┌"
	TopRightCorner    = "┐"
	BottomLeftCorner  = " └"
	BottomRightCorner = "┘"
)

//go:embed art.txt
var SnakeArt string

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
	fmt.Print(indent + TopLeftCorner)
	for i := 0; i < cols; i++ {
		fmt.Print(HorizontalBorder)
	}
	fmt.Println(TopRightCorner + "\r ")
}

func renderBottomBorder(cols int) {
	indent := calculateBoardIndent(cols)
	fmt.Print("\r\n" + indent + BottomLeftCorner)
	for i := 0; i < cols; i++ {
		fmt.Print(HorizontalBorder)
	}
	fmt.Println(BottomRightCorner)
}

func getCellGlyph(coord game.Coord, snake *game.Snake, food *game.Food) string {
	// Check if snake occupies this position
	if _, occupied := snake.Occupied[coord]; occupied {
		if coord == snake.Head() {
			return SnakeHeadGlyph
		}
		return SnakeBodyGlyph
	}

	// Check if food is at this position
	if food.Coord.X == coord.X && food.Coord.Y == coord.Y {
		return FoodGlyph
	}

	// Empty cell
	return GridGlyph
}

func RenderBoard(g *game.Game) {
	board := g.Board
	snake := g.Snake
	food := g.Food
	cols := len(board[0])
	indent := calculateBoardIndent(cols)

	clearScreen()
	hideCursor()

	fmt.Printf(indent+" Score: %d \r\n", g.GetScore())

	// Render top border
	renderTopBorder(cols)

	// Render board rows
	for row := range board {
		if row != 0 {
			fmt.Print("\r\n ")
		}

		fmt.Print(indent + VerticalBorder) // Left border

		// Render each cell in the row
		for col := range board[row] {
			coord := game.Coord{row, col}
			glyph := getCellGlyph(coord, snake, food)
			fmt.Print(glyph)
		}

		fmt.Print(VerticalBorder) // Right border
	}

	// Render bottom border
	renderBottomBorder(cols)
}

func ShowStartMenu() {
	clearScreen()

	// Center the art horizontally
	artLines := strings.Split(SnakeArt, "\n")

	for _, line := range artLines {
		if line == "" {
			fmt.Print("\r\n")
			continue
		}

		fmt.Printf("	%s\r\n", line)
	}

	fmt.Print("\r\n\r\n")
	fmt.Print("        > Press any key to start 🟢\r\n")
	fmt.Print("        > Press 'q' or ESC to quit ❌\r\n")
}

func ShowGameOver(score int) {
	clearScreen()
	fmt.Print("\r\n\r\n        ══════════════════════════════\r\n")
	fmt.Print("        	💀 GAME OVER! 💀\r\n")
	fmt.Print("        ══════════════════════════════\r\n")
	fmt.Printf("        • Score: %d\r\n", score)
	fmt.Print("        - - - - - - - - - - - - - - - \r\n\r\n")

	fmt.Print("        > Press any key to restart 🔄\r\n")
	fmt.Print("        > Press 'q' or ESC to quit ❌\r\n")
}

func ShowWinScreen() {
	clearScreen()
	fmt.Print("\r\n\r\n        ══════════════════════════════\r\n")
	fmt.Print("           🎉 VICTORY! 🎉\r\n")
	fmt.Print("        ══════════════════════════════\r\n\r\n")
	fmt.Print("        > Press any key to restart\r\n")
	fmt.Print("        > Press 'q' or ESC to quit\r\n")
}
