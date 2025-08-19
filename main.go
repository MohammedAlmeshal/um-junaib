package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

const BoardSize = 10

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Enqueue(val T) bool {
	q.data = append(q.data, val)
	return true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val, true
}

type Coord struct{ x, y int }

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Food struct {
	coord Coord
}

func (food *Food) spawn() {
	food.coord.x = rand.Intn(BoardSize)
	food.coord.y = rand.Intn(BoardSize)
}

type Snake struct {
	body      Queue[Coord]
	occupied  map[Coord]bool
	direction Direction
}

func (snake *Snake) move() {

	head := snake.head()
	var new_coord Coord
	switch snake.direction {
	case UP:
		new_coord = Coord{head.x - 1, head.y}
	case RIGHT:
		new_coord = Coord{head.x, head.y + 1}
	case DOWN:
		new_coord = Coord{head.x + 1, head.y}
	case LEFT:
		new_coord = Coord{head.x, head.y - 1}
	}

	coord, _ := snake.body.Dequeue()
	delete(snake.occupied, coord)

	snake.body.Enqueue(new_coord)
	snake.occupied[new_coord] = true

}

func (snake *Snake) tail() Coord {
	return snake.body.data[0]
}
func (snake *Snake) head() Coord {
	return snake.body.data[len(snake.body.data)-1]
}
func (snake *Snake) grow(coord Coord) {
	snake.body.data = append([]Coord{coord}, snake.body.data...)
	snake.occupied[coord] = true
}

func (snake *Snake) boarderCollided() bool {
	head := snake.head()
	return head.x >= BoardSize || head.y >= BoardSize || head.x < 0 || head.y < 0

}

func (snake *Snake) selfCollided() bool {
	// if head shared a coord with the body, the snake collided
	return len(snake.body.data) != len(snake.occupied)

}

type Game struct {
	board [BoardSize][BoardSize]int
	snake *Snake
	food  *Food
}

func NewGame() *Game {
	coord := Coord{0, 1}
	snake := &Snake{occupied: make(map[Coord]bool), direction: RIGHT}
	snake.body.Enqueue(coord)
	snake.occupied[coord] = true

	food := &Food{}
	food.spawn()

	return &Game{
		snake: snake,
		food:  food,
	}
}

func (game *Game) tick(dir Direction) {
	snake := game.snake
	food := game.food

	snake.direction = dir
	ateFood := snake.head() == food.coord
	tail := snake.tail()

	snake.move()

	if snake.boarderCollided() || snake.selfCollided() {
		fmt.Println("\nGAME OVER!")
		return
	}

	if ateFood {
		snake.grow(tail)

		for {
			food.spawn()
			if !snake.occupied[food.coord] {
				break
			}
		}
	}
}

func run(game *Game) {
	pendingDir := RIGHT
	speed := 300 * time.Millisecond

	ticker := time.NewTicker(speed)
	defer ticker.Stop()

	inputChan := startInputReader()

	for {
		select {
		case <-ticker.C:
			game.tick(pendingDir)
			renderBoard(game)

		case data, ok := <-inputChan:
			if !ok {
				fmt.Println("\nInput channel closed!")
				return
			}
			if dir, ok := inputHandler(data); ok && validTurn(game.snake.direction, dir) {
				pendingDir = dir
			}
		}
	}
}

func validTurn(current, next Direction) bool {
	isOpposite := (current ^ next) == 1
	isSame := current == next
	return !(isOpposite || isSame)
}

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

func main() {

	game := NewGame()
	run(game)

}
