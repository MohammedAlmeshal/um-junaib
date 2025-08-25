package main

import (
	"math/rand"
)

type Snake struct {
	body      *Queue[Coord]
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

	coord, _ := snake.body.dequeue()
	delete(snake.occupied, coord)

	snake.body.enqueue(new_coord)
	snake.occupied[new_coord] = true

}

func (snake *Snake) tail() Coord {
	return snake.body.data[snake.body.head]
}
func (snake *Snake) head() Coord {
	return snake.body.data[(snake.body.tail-1+snake.body.size)%snake.body.size]
}
func (snake *Snake) grow(coord Coord) {
	snake.body.pushFront(coord)
	snake.occupied[coord] = true
}

func (snake *Snake) borderCollided() bool {
	return snake.borderCollidedAt(snake.head())
}
func (snake *Snake) borderCollidedAt(coord Coord) bool {
	return coord.x >= BoardSize || coord.y >= BoardSize || coord.x < 0 || coord.y < 0
}

func (snake *Snake) selfCollided() bool {
	return len(snake.occupied) != snake.body.count
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
	newPositions := make([]Coord, snake.body.count)
	for i := 0; i < snake.body.count; i++ {
		pos := (snake.body.head + i) % snake.body.size
		coord := snake.body.data[pos]

		newCoord := Coord{coord.x + shift.x, coord.y + shift.y}
		if snake.borderCollidedAt(newCoord) || newCoord == foodCoord {
			// Abort without touching body or occupied
			return
		}
		newPositions[i] = newCoord
	}

	// Commit: update body + occupied
	snake.occupied = make(map[Coord]bool)
	for i := 0; i < snake.body.count; i++ {
		pos := (snake.body.head + i) % snake.body.size
		snake.body.data[pos] = newPositions[i]
		snake.occupied[newPositions[i]] = true
	}
}
