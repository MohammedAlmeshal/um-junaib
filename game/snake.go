package game

import (
	"math/rand"
)

type Snake struct {
	body      *Queue[Coord]
	Occupied  map[Coord]bool
	direction Direction
}

func (snake *Snake) move() {

	head := snake.head()
	var new_coord Coord
	switch snake.direction {
	case UP:
		new_coord = Coord{head.X - 1, head.Y}
	case RIGHT:
		new_coord = Coord{head.X, head.Y + 1}
	case DOWN:
		new_coord = Coord{head.X + 1, head.Y}
	case LEFT:
		new_coord = Coord{head.X, head.Y - 1}
	}

	coord, _ := snake.body.Dequeue()
	delete(snake.Occupied, coord)

	snake.body.Enqueue(new_coord)
	snake.Occupied[new_coord] = true

}

func (snake *Snake) tail() Coord {
	return snake.body.Data[snake.body.Head]
}
func (snake *Snake) head() Coord {
	return snake.body.Data[(snake.body.Tail-1+snake.body.Size)%snake.body.Size]
}
func (snake *Snake) Head() Coord {
	return snake.head()
}
func (snake *Snake) grow(coord Coord) {
	snake.body.PushFront(coord)
	snake.Occupied[coord] = true
}

func (snake *Snake) borderCollided() bool {
	return snake.borderCollidedAt(snake.head())
}
func (snake *Snake) borderCollidedAt(coord Coord) bool {
	return coord.X >= BoardSize || coord.Y >= BoardSize || coord.X < 0 || coord.Y < 0
}

func (snake *Snake) selfCollided() bool {
	return len(snake.Occupied) != snake.body.Count
}

func (snake *Snake) shift(foodCoord Coord) {
	var dirOptions = map[Direction][]Coord{
		UP:    {{0, -1}, {0, 1}},
		DOWN:  {{0, -1}, {0, 1}},
		LEFT:  {{-1, 0}, {1, 0}},
		RIGHT: {{-1, 0}, {1, 0}},
	}

	shift := dirOptions[snake.direction][rand.Intn(2)]

	// Stage new positions
	newPositions := make([]Coord, snake.body.Count)
	for i := 0; i < snake.body.Count; i++ {
		pos := (snake.body.Head + i) % snake.body.Size
		coord := snake.body.Data[pos]

		newCoord := Coord{coord.X + shift.X, coord.Y + shift.Y}
		if snake.borderCollidedAt(newCoord) || newCoord == foodCoord {
			// Abort without touching body or occupied
			return
		}
		newPositions[i] = newCoord
	}

	// Commit: update body + occupied
	snake.Occupied = make(map[Coord]bool)
	for i := 0; i < snake.body.Count; i++ {
		pos := (snake.body.Head + i) % snake.body.Size
		snake.body.Data[pos] = newPositions[i]
		snake.Occupied[newPositions[i]] = true
	}
}
