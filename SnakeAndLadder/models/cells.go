package models

type Cell struct {
	number   int
	isSnake  bool
	isLadder bool
}

func NewCell(number int) *Cell {
	return &Cell{
		number: number,
	}
}

func (c *Cell) GetNumber() int {
	return c.number
}

func (c *Cell) SetIsSnake(isSnake bool) {
	c.isSnake = isSnake
}

func (c *Cell) SetIsLadder(isLadder bool) {
	c.isLadder = isLadder
}

func (c *Cell) IsSnake() bool {
	return c.isSnake
}

func (c *Cell) IsLadder() bool {
	return c.isLadder
}
