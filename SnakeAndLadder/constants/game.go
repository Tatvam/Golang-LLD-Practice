package constants

type GameStatus int

const (
	NotStarted GameStatus = iota
	InProgress
	Completed
)
