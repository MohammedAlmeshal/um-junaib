package game

import "math/rand"

type Food struct {
	Coord Coord
}

func (food *Food) Spawn() {
	food.Coord.X = rand.Intn(BoardSize)
	food.Coord.Y = rand.Intn(BoardSize)
}
