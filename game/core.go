package game

import (
	_ "embed"
	"errors"
)

const BoardSize = 15

type GameStatus int
type Coord struct{ X, Y int }

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

const (
	GameRunning GameStatus = iota
	GameDead
	GameWon
	GameTerminated
)

//go:embed art.txt
var SnakeArt string

type Queue[T any] struct {
	Data        []T
	Head, Tail  int
	Size, Count int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		Data: make([]T, size),
		Size: size,
	}
}

func (q *Queue[T]) Enqueue(val T) error {
	if q.Count == q.Size {
		return errors.New("queue is full")
	}
	q.Data[q.Tail] = val
	q.Tail = (q.Tail + 1) % q.Size
	q.Count++
	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.Count == 0 {
		return zero, errors.New("queue is empty")
	}
	val := q.Data[q.Head]
	q.Head = (q.Head + 1) % q.Size
	q.Count--
	return val, nil
}

func (q *Queue[T]) PushFront(val T) error {
	if q.Count == q.Size {
		return errors.New("queue is full")
	}
	// move head backwards
	q.Head = (q.Head - 1 + q.Size) % q.Size
	q.Data[q.Head] = val
	q.Count++
	return nil
}
