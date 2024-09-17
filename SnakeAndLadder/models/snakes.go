package models

import (
	"fmt"
	"math/rand"
)

type Snake struct {
	Start int
	End   int
}

func (s *Snake) GetStart() int {
	return s.Start
}

func (s *Snake) GetEnd() int {
	return s.End
}

func NewSnakes(size int, board *Board) []Snake {
	snakes := make([]Snake, size)
	for i := 0; i < size; {
		start := rand.Intn(board.GetSize()*board.GetSize()-1) + 2
		end := rand.Intn(start-1) + 1
		if end < start && start != board.GetSize()*board.GetSize() && end != 0 &&
			start != end && !board.IsLadder(start) && !board.IsLadder(end) &&
			!board.IsSnake(start) && !board.IsSnake(end) {
			snakes[i] = Snake{
				Start: start,
				End:   end,
			}
			fmt.Println("Snake", start, end)
			board.MarkSnake(start)
			i++
		}
	}
	return snakes
}
