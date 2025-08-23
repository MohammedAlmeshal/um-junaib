package main

import _ "embed"

const BoardSize = 15

type GameStatus int

const (
	GameRunning GameStatus = iota
	GameDead
	GameWon
	GameTerminated
)

//go:embed art.txt
var snakeArt string

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
