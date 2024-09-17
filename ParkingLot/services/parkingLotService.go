package services

import (
	"ParkingLot/constants"
	"ParkingLot/errors"
	"ParkingLot/models"
	"sync"
)

var once sync.Once
var parkingLotService *ParkingLotService

type IParkingLotService interface {
}
type ParkingLotService struct {
	parkingLot models.ParkingLot
}

func NewParkingLotService() *ParkingLotService {
	once.Do(func() {
		parkingLotService = &ParkingLotService{
			parkingLot: CreateParkingLot(),
		}
	})
	return parkingLotService
}

func CreateParkingLot() models.ParkingLot {
	gates := make(map[constants.GateType][]models.Gate)
	gates[constants.ENTRY] = models.NewEntryGates()
	gates[constants.EXIT] = models.NewExitGates()
	floors := models.NewFloors(1)
	parkingLot := models.NewParkingLot(1, "ParkingLot1", floors, constants.Active, gates, "Address1")
	return parkingLot
}

func (p *ParkingLotService) ParkVehicle(vechicle models.Vehicle) (*models.Slot, *errors.Error) {
	if p.parkingLot.ParkingLotStatus == constants.Inactive {
		return nil, errors.NewError(400, "PF-001", "Parking Lot is Inactive")
	}
	if p.parkingLot.IsParkingLotFull(vechicle.VehicleType) {
		return nil, errors.NewError(400, "PF-002", "Parking Lot is Full")
	}
	floor, err := p.GetFreeFloor(vechicle.VehicleType)
	if err != nil {
		return nil, err
	}
	slot, err := floor.BookSlot(vechicle.VehicleType)
	if err != nil {
		return nil, err
	}
	return slot, nil

}

func (p *ParkingLotService) GetFreeSlot(vehicleType constants.VehicleType) (*models.Slot, *errors.Error) {
	for _, floor := range p.parkingLot.Floors {
		if !floor.IsFloorFull(vehicleType) {
			slot, err := floor.GetFreeSlot(vehicleType)
			if err != nil {
				return &models.Slot{}, errors.NewError(400, "PF-003", "No Free Slot Available")
			}
			return slot, nil
		}
	}
	return &models.Slot{}, errors.NewError(400, "PF-004", "No Free Slot Available")
}

func (p *ParkingLotService) GetFreeFloor(vehicleType constants.VehicleType) (*models.Floor, *errors.Error) {
	for _, floor := range p.parkingLot.Floors {
		if !floor.IsFloorFull(vehicleType) {
			return &floor, nil
		}
	}
	return nil, errors.NewError(400, "PF-004", "No Free Slot Available")
}

func (p *ParkingLotService) UnparkVehicle(slot models.Slot, vehicle constants.VehicleType) *errors.Error {
	var parkedFloor *models.Floor
	for _, floor := range p.parkingLot.Floors {
		if floor.FloorID == slot.FloorID {
			parkedFloor = &floor
			break
		}
	}
	parkedFloor.UnparkVehicle(slot, vehicle)
	return nil
}

func (p *ParkingLotService) GetEntryGate(gateID int) (*models.Gate, *errors.Error) {
	for _, gate := range p.parkingLot.Gates[constants.ENTRY] {
		if gate.GateID == gateID {
			return &gate, nil
		}
	}
	return nil, errors.NewError(400, "PF-005", "Invalid Gate ID")
}
