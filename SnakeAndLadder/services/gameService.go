package services

import (
	"SnakeAndLadder/constants"
	"SnakeAndLadder/errors"
	"SnakeAndLadder/models"
	"fmt"
	"sync"
)

type GameService struct {
	game        *models.Game
	diceService IDiceService
}

var gameService *GameService

var once sync.Once

func NewGameService() *GameService {
	once.Do(func() {
		gameService = &GameService{}
	})
	return gameService
}

func (g *GameService) InitializeGame(size, ladders, snakes int) {
	board := models.NewBoard(size, ladders, snakes)
	board.AddSnakes(snakes)
	board.AddLadders(ladders)
	g.game = models.NewGame(board)
	g.diceService = NewDiceService(1, 6, 1)
}

func (g *GameService) StartGame() {
	g.game.Status = constants.InProgress
}

func (g *GameService) EndGame() {
	g.game.Status = constants.Completed
}

func (g *GameService) PlayGame() {
	g.StartGame()
	fmt.Println("Game started")
	for g.game.Status == constants.InProgress {
		player := g.GetCurrentPlayer()
		err := g.GetPlayerInput(player)
		if err != nil {
			fmt.Println("error getting player input", err)
			fmt.Println("Please try again")
			continue
		}
		newPosition, err := g.MovePlayerToken(player)

		if err == errors.ErrCannotMove {
			fmt.Println("Player", player.GetName(), "cannot move")
			g.SetCurrentPlayer()
			continue
		}
		if err != nil {
			fmt.Println("error moving player token", err)
			g.game.Status = constants.Completed
			break
		}
		if g.IsWinningPosition(newPosition) {
			fmt.Println("Player", player.GetName(), "won the game")
			g.EndGame()
			break
		}
		fmt.Println("Player", player.GetName(), "moved to", newPosition)
		g.SetCurrentPlayer()
	}

}

func (g *GameService) IsWinningPosition(position int) bool {
	return position == g.game.GetEndPosition()
}

func (g *GameService) MovePlayerToken(player *models.Player) (int, error) {
	if g.game.Status == constants.NotStarted {
		return 0, errors.ErrGameNotStarted
	}
	if g.game.Status == constants.Completed {
		return 0, errors.ErrGameAlreadyEnd
	}

	playerToken := player.GetToken()
	diceValue := g.diceService.RollDice()[0]
	fmt.Println("Dice value", diceValue)
	newPostion := playerToken.GetPosition() + diceValue

	if newPostion > g.game.Board.GetSize()*g.game.Board.GetSize() {
		return 0, errors.ErrCannotMove
	}

	if g.game.Board.IsSnake(newPostion) {
		fmt.Println("Player", player.GetName(), "got bitten by snake")
		newPostion = g.game.Board.GetSnakeEnd(newPostion)
	} else if g.game.Board.IsLadder(newPostion) {
		fmt.Println("Player", player.GetName(), "got ladder")
		newPostion = g.game.Board.GetLadderEnd(newPostion)
	}
	playerToken.SetPosition(newPostion)
	return newPostion, nil
}

func (g *GameService) GetCurrentPlayer() *models.Player {
	return g.game.Players[0]
}

func (g *GameService) SetCurrentPlayer() {
	g.game.Players = append(g.game.Players[1:], g.game.Players[0])
}

func (g *GameService) GetPlayerInput(player *models.Player) error {
	var playerInput string
	fmt.Println("Enter `r` to roll the dice for player turn ", player.GetName(), player.GetToken().GetPosition())
	if _, err := fmt.Scan(&playerInput); err != nil {
		return err
	}
	return nil
}

func (g *GameService) AddPlayer(playerName string) {
	if len(g.game.Players) == 4 {
		fmt.Println("maximum 4 players allowed")
		return
	}
	color := g.GetAvailableColor()
	if color == constants.None {
		fmt.Println("No color available")
		return
	}
	player := models.NewPlayer(playerName, models.NewToken(color))
	g.game.Players = append(g.game.Players, player)

}

func (g *GameService) GetAvailableColor() constants.Color {
	for color, assigned := range g.game.TokensAssigned {
		if !assigned {
			g.game.TokensAssigned[color] = true
			return color
		}
	}
	return constants.None
}
