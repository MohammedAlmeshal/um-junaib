package main

import "math/rand"

type Food struct {
	coord Coord
}

func (food *Food) spawn() {
	food.coord.x = rand.Intn(BoardSize)
	food.coord.y = rand.Intn(BoardSize)
}
