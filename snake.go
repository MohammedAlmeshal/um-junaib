package main

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

func (snake *Snake) boarderCollided() bool {

	head := snake.head()
	return head.x >= BoardSize || head.y >= BoardSize || head.x < 0 || head.y < 0

}

func (snake *Snake) selfCollided() bool {
	return len(snake.occupied) != snake.body.count

}
