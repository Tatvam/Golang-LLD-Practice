package errors

import "errors"

var (
	ErrCannotMove     = errors.New("cannot move")
	ErrInvalidMove    = errors.New("invalid move")
	ErrGameNotStarted = errors.New("game not started")
	ErrGameAlreadyEnd = errors.New("game already ended")
)
