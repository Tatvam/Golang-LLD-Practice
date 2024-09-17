package models

import "ParkingLot/constants"

type ParkingLot struct {
	ParkingLotID     int
	ParkingLotName   string
	Floors           []Floor
	ParkingLotStatus constants.ParkingLotStatus
	Gates            map[constants.GateType][]Gate
	Address          string
}

func NewParkingLot(parkingLotID int, parkingLotName string, floors []Floor, parkingLotStatus constants.ParkingLotStatus, gates map[constants.GateType][]Gate, address string) ParkingLot {
	return ParkingLot{
		ParkingLotID:     parkingLotID,
		ParkingLotName:   parkingLotName,
		Floors:           floors,
		ParkingLotStatus: parkingLotStatus,
		Gates:            gates,
		Address:          address,
	}
}

func (p *ParkingLot) IsParkingLotActive() bool {
	return p.ParkingLotStatus == constants.Active
}

func (p *ParkingLot) IsParkingLotFull(vehicleType constants.VehicleType) bool {
	for _, floor := range p.Floors {
		if !floor.IsFloorFull(vehicleType) {
			return false
		}
	}
	return true
}
