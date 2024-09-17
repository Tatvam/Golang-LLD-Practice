package models

import "SnakeAndLadder/constants"

type Token struct {
	position int
	color    constants.Color
}

func NewToken(color constants.Color) *Token {
	return &Token{
		position: 0,
		color:    color,
	}
}

func (t *Token) GetPosition() int {
	return t.position
}

func (t *Token) GetColor() constants.Color {
	return t.color
}

func (t *Token) SetPosition(position int) {
	t.position = position
}
