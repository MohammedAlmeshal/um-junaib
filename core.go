package main

import (
	_ "embed"
	"errors"
)

const BoardSize = 15

type GameStatus int
type Coord struct{ x, y int }

const (
	GameRunning GameStatus = iota
	GameDead
	GameWon
	GameTerminated
)

//go:embed art.txt
var snakeArt string

type Queue[T any] struct {
	data        []T
	head, tail  int
	size, count int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, size),
		size: size,
	}
}

func (q *Queue[T]) Enqueue(val T) error {
	if q.count == q.size {
		return errors.New("queue is full")
	}
	q.data[q.tail] = val
	q.tail = (q.tail + 1) % q.size
	q.count++
	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.count == 0 {
		return zero, errors.New("queue is empty")
	}
	val := q.data[q.head]
	q.head = (q.head + 1) % q.size
	q.count--
	return val, nil
}

func (q *Queue[T]) PushFront(val T) error {
	if q.count == q.size {
		return errors.New("queue is full")
	}
	// move head backwards
	q.head = (q.head - 1 + q.size) % q.size
	q.data[q.head] = val
	q.count++
	return nil
}
