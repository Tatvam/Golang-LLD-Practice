package models

import (
	"fmt"

	"golang.org/x/exp/rand"
)

type Ladder struct {
	start int
	end   int
}

func (l *Ladder) GetStart() int {
	return l.start
}

func (l *Ladder) GetEnd() int {
	return l.end
}

func NewLadders(size int, board *Board) []Ladder {
	ladders := make([]Ladder, size)
	for i := 0; i < size; {
		start := rand.Intn(board.GetSize()*board.GetSize() - 1)
		end := start + rand.Intn(board.GetSize()*board.GetSize()-start) - 1
		if end > start && end != board.GetSize()*board.GetSize() && start != 0 &&
			start != end && !board.IsLadder(start) && !board.IsLadder(end) &&
			!board.IsSnake(start) && !board.IsSnake(end) {
			ladders[i] = Ladder{
				start: start,
				end:   end,
			}
			fmt.Println("Ladder", start, end)
			board.MarkLadder(start)
			i++
		}
	}
	return ladders
}
