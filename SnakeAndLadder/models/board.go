package models

import (
	"SnakeAndLadder/utils"
)

type Board struct {
	board   [][]*Cell
	ladders []Ladder
	snakes  []Snake
}

func NewBoard(size int, ladderSize int, snakeSize int) *Board {
	board := GenerateBoard(size)
	return &Board{
		board:   board,
		snakes:  make([]Snake, 0, snakeSize),
		ladders: make([]Ladder, 0, ladderSize),
	}
}

func GenerateBoard(size int) [][]*Cell {
	board := make([][]*Cell, size)
	direction := 1
	for i := 0; i < size; i++ {
		board[i] = make([]*Cell, size)
		for j := 0; j < size; j++ {
			if direction == 1 {
				board[i][j] = NewCell(i*size + j)
			} else {
				board[i][j] = NewCell((i+1)*size - j)
			}
		}
		direction = -direction
	}
	return board
}

func (b *Board) AddSnakes(size int) {
	b.snakes = NewSnakes(size, b)
}

func (b *Board) AddLadders(size int) {
	b.ladders = NewLadders(size, b)
}

func MarkSnakesAndLadders(board [][]*Cell, snakes []Snake, ladders []Ladder) {
	for _, snakes := range snakes {
		i, j := utils.GetRowAndColumn(len(board), snakes.GetStart())
		board[i][j].SetIsSnake(true)
	}
	for _, ladders := range ladders {
		i, j := utils.GetRowAndColumn(len(board), ladders.GetStart())
		board[i][j].SetIsLadder(true)
	}
}

func (b *Board) MarkSnake(number int) {
	i, j := utils.GetRowAndColumn(b.GetSize(), number)
	b.board[i][j].SetIsSnake(true)
}

func (b *Board) MarkLadder(number int) {
	i, j := utils.GetRowAndColumn(b.GetSize(), number)
	b.board[i][j].SetIsLadder(true)
}

func (b *Board) GetSize() int {
	return len(b.board)
}

func (b *Board) IsSnake(number int) bool {
	i, j := utils.GetRowAndColumn(b.GetSize(), number)
	return b.board[i][j].IsSnake()
}

func (b *Board) IsLadder(number int) bool {
	i, j := utils.GetRowAndColumn(b.GetSize(), number)
	return b.board[i][j].IsLadder()
}

func (b *Board) GetSnakeEnd(position int) int {
	for _, snake := range b.snakes {
		if snake.GetStart() == position {
			return snake.GetEnd()
		}
	}
	return -1
}

func (b *Board) GetLadderEnd(position int) int {
	for _, ladder := range b.ladders {
		if ladder.GetStart() == position {
			return ladder.GetEnd()
		}
	}
	return -1
}
