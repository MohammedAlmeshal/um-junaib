package game

import (
	"math/rand"
	"time"
)

const gameSpeed = 150

type Game struct {
	Board [BoardSize][BoardSize]int
	Snake *Snake
	Food  *Food
}

func NewGame() *Game {
	coord := Coord{0, 1}
	snake := &Snake{
		body:      NewQueue[Coord](BoardSize * BoardSize),
		Occupied:  make(map[Coord]bool),
		direction: RIGHT}

	snake.body.Enqueue(coord)
	snake.Occupied[coord] = true

	food := &Food{}
	food.Spawn()

	return &Game{
		Snake: snake,
		Food:  food,
	}
}

func (game *Game) IsValidTurn(nextDir Direction) bool {
	currentDir := game.Snake.direction
	isOpposite := (currentDir ^ nextDir) == 1
	isSame := currentDir == nextDir
	return !(isOpposite || isSame)
}

func (game *Game) isWon() bool {
	return len(game.Snake.Occupied) == BoardSize*BoardSize
}

func (game *Game) GetScore() int {
	return len(game.Snake.Occupied)
}

func (game *Game) tick(dir Direction) GameStatus {
	snake := game.Snake
	food := game.Food

	snake.direction = dir
	ateFood := snake.head() == food.Coord
	tail := snake.tail()

	snake.move()

	// random shift
	if snake.body.Count > 1 {
		if rand.Intn(100) < 3 {
			snake.shift(food.Coord)
		}
	}

	if snake.borderCollided() || snake.selfCollided() {
		return GameDead
	}

	if ateFood {
		snake.grow(tail)

		if game.isWon() {
			return GameWon
		}

		for {
			food.Spawn()
			if !snake.Occupied[food.Coord] {
				break
			}
		}
	}
	return GameRunning
}

func (game *Game) Run(inputChan <-chan []byte, renderFunc func(*Game), inputHandler func([]byte) (Direction, bool)) GameStatus {
	pendingDir := RIGHT
	speed := gameSpeed * time.Millisecond

	ticker := time.NewTicker(speed)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			status := game.tick(pendingDir)
			if status != GameRunning {
				return status
			}
			renderFunc(game)

		case data, ok := <-inputChan:
			if !ok {
				return GameTerminated
			}
			if nextDir, ok := inputHandler(data); ok && game.IsValidTurn(nextDir) {
				pendingDir = nextDir
			}
		}
	}
}
