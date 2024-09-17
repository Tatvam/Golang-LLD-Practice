package main

import "SnakeAndLadder/services"

func main() {
	// This is the main function

	gameService := services.NewGameService()

	gameService.InitializeGame(10, 5, 5)
	gameService.AddPlayer("Tatvam")
	gameService.AddPlayer("Dadheech")

	gameService.PlayGame()

}
