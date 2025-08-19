package main

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Snake struct {
	body      Queue[Coord]
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

	coord, _ := snake.body.Dequeue()
	delete(snake.occupied, coord)

	snake.body.Enqueue(new_coord)
	snake.occupied[new_coord] = true

}

func (snake *Snake) tail() Coord {
	return snake.body.data[0]
}
func (snake *Snake) head() Coord {
	return snake.body.data[len(snake.body.data)-1]
}
func (snake *Snake) grow(coord Coord) {
	snake.body.data = append([]Coord{coord}, snake.body.data...)
	snake.occupied[coord] = true
}

func (snake *Snake) boarderCollided() bool {
	head := snake.head()

	return head.x >= BoardSize || head.y >= BoardSize || head.x < 0 || head.y < 0

}

func (snake *Snake) selfCollided() bool {
	// if head shared a coord with the body, the snake collided
	return len(snake.body.data) != len(snake.occupied)

}
