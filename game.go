package main

import (
	"time"
)

type Game struct {
	board [BoardSize][BoardSize]int
	snake *Snake
	food  *Food
}

func NewGame() *Game {
	coord := Coord{0, 1}
	snake := &Snake{
		body:      NewQueue[Coord](BoardSize*BoardSize),
		occupied:  make(map[Coord]bool),
		direction: RIGHT}

	snake.body.Enqueue(coord)
	snake.occupied[coord] = true

	food := &Food{}
	food.spawn()

	return &Game{
		snake: snake,
		food:  food,
	}
}

func (game *Game) tick(dir Direction) GameStatus {
	snake := game.snake
	food := game.food

	snake.direction = dir
	ateFood := snake.head() == food.coord
	tail := snake.tail()

	snake.move()

	if snake.boarderCollided() || snake.selfCollided() {
		return GameDead
	}

	if ateFood {
		snake.grow(tail)

		// Check if player won (snake fills entire board)
		if len(snake.occupied) == BoardSize*BoardSize {
			return GameWon
		}

		for {
			food.spawn()
			if !snake.occupied[food.coord] {
				break
			}
		}
	}
	return GameRunning
}

func (game *Game) getScore() int {
	return len(game.snake.occupied)
}

func Run(game *Game, inputChan <-chan []byte) GameStatus {
	pendingDir := RIGHT
	speed := 100 * time.Millisecond

	ticker := time.NewTicker(speed)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			status := game.tick(pendingDir)
			if status != GameRunning {
				return status
			}
			renderBoard(game)

		case data, ok := <-inputChan:
			if !ok {
				return GameTerminated
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
