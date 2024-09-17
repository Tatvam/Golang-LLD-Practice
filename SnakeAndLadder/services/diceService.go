package services

import "golang.org/x/exp/rand"

type IDiceService interface {
	RollDice() []int
}

type DiceService struct {
	minValue  int
	maxValue  int
	noOfRolls int
}

func NewDiceService(minValue int, maxValue int, noOfRolls int) *DiceService {
	return &DiceService{
		minValue:  minValue,
		maxValue:  maxValue,
		noOfRolls: noOfRolls,
	}
}

func (d *DiceService) RollDice() []int {
	var diceValues []int
	for i := 0; i < d.noOfRolls; i++ {
		randNumber := d.minValue + rand.Intn(d.maxValue-d.minValue+1)
		diceValues = append(diceValues, randNumber)
	}
	return diceValues
}
