package services

type IGame interface {
	InitializeGame(int, int, int, int)
	PlayGame()
	MovePlayerToken(int)
	IsGameCompleted() bool
}
