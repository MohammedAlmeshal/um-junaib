package main

import (
	"time"
)

type Game struct {
	board [BoardSize][BoardSize]int
	snake *Snake
	food  *Food
}

func newGame() *Game {
	coord := Coord{0, 1}
	snake := &Snake{
		body:      newQueue[Coord](BoardSize * BoardSize),
		occupied:  make(map[Coord]bool),
		direction: RIGHT}

	snake.body.enqueue(coord)
	snake.occupied[coord] = true

	food := &Food{}
	food.spawn()

	return &Game{
		snake: snake,
		food:  food,
	}
}

func (game *Game) isValidTurn(nextDir Direction) bool {
	currentDir := game.snake.direction
	isOpposite := (currentDir ^ nextDir) == 1
	isSame := currentDir == nextDir
	return !(isOpposite || isSame)
}

func (game *Game) isWon() bool {
	return len(game.snake.occupied) == BoardSize*BoardSize
}

func (game *Game) getScore() int {
	return len(game.snake.occupied)
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

		if game.isWon() {
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

func (game *Game) run(inputChan <-chan []byte) GameStatus {
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
			if nextDir, ok := inputHandler(data); ok && game.isValidTurn(nextDir) {
				pendingDir = nextDir
			}
		}
	}
}
