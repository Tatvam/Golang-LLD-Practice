package models

import "ParkingLot/constants"

type Gate struct {
	GateID   int
	GateType constants.GateType
}

func NewEntryGates() []Gate {
	return []Gate{
		{
			GateID:   1,
			GateType: constants.ENTRY,
		},
		{
			GateID:   2,
			GateType: constants.ENTRY,
		},
	}
}

func NewExitGates() []Gate {
	return []Gate{
		{
			GateID:   3,
			GateType: constants.EXIT,
		},
		{
			GateID:   4,
			GateType: constants.EXIT,
		},
	}
}
