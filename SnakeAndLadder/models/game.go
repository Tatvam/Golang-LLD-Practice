package models

import "SnakeAndLadder/constants"

type Game struct {
	Board          *Board
	Players        []*Player
	Status         constants.GameStatus
	TokensAssigned map[constants.Color]bool
}

func NewGame(board *Board) *Game {
	return &Game{
		Board:   board,
		Players: make([]*Player, 0),
		Status:  constants.NotStarted,
		TokensAssigned: map[constants.Color]bool{
			constants.Red:    false,
			constants.Yellow: false,
			constants.Green:  false,
			constants.Blue:   false,
		},
	}
}

func (g *Game) GetEndPosition() int {
	return g.Board.GetSize() * g.Board.GetSize()
}
