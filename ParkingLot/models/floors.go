package models

import (
	"ParkingLot/constants"
	"ParkingLot/errors"
)

type Floor struct {
	FloorID        int
	FloorName      string
	AvailableSlots map[constants.VehicleType]int
	Slots          map[constants.VehicleType][]Slot
}

func NewFloors(NumberOfFloors int) []Floor {
	floors := []Floor{}
	for i := 0; i < NumberOfFloors; i++ {
		floors = append(floors, Floor{
			FloorID:   i,
			FloorName: "Floor" + string(i),
			AvailableSlots: map[constants.VehicleType]int{
				constants.HEAVY_VEHICLE:  1,
				constants.MEDIUM_VEHICLE: 1,
				constants.LIGHT_VEHICLE:  1,
			},
			Slots: map[constants.VehicleType][]Slot{
				constants.HEAVY_VEHICLE:  NewSlots(i, 1, constants.HEAVY_VEHICLE),
				constants.MEDIUM_VEHICLE: NewSlots(i, 1, constants.MEDIUM_VEHICLE),
				constants.LIGHT_VEHICLE:  NewSlots(i, 1, constants.LIGHT_VEHICLE),
			},
		})
	}
	return floors
}

func NewFloor(FloorID int) Floor {
	return Floor{
		FloorID:   FloorID,
		FloorName: "Floor" + string(FloorID),
		AvailableSlots: map[constants.VehicleType]int{
			constants.HEAVY_VEHICLE:  10,
			constants.MEDIUM_VEHICLE: 10,
			constants.LIGHT_VEHICLE:  10,
		},
		Slots: map[constants.VehicleType][]Slot{
			constants.HEAVY_VEHICLE:  NewSlots(FloorID, 10, constants.HEAVY_VEHICLE),
			constants.MEDIUM_VEHICLE: NewSlots(FloorID, 10, constants.MEDIUM_VEHICLE),
			constants.LIGHT_VEHICLE:  NewSlots(FloorID, 10, constants.LIGHT_VEHICLE),
		},
	}
}

func (f *Floor) IsFloorFull(vehicleType constants.VehicleType) bool {
	return f.AvailableSlots[vehicleType] == 0
}

func (f *Floor) GetAvailableSlots(vehicleType constants.VehicleType) []Slot {
	return f.Slots[vehicleType]
}

func (f *Floor) GetFreeSlot(vehicleType constants.VehicleType) (*Slot, *errors.Error) {
	for _, slot := range f.Slots[vehicleType] {
		if slot.GetSlotStatus() == constants.SlotAvailable {
			return &slot, nil
		}
	}
	return nil, errors.ErrParkingLotFull
}

func (f *Floor) BookSlot(vehicleType constants.VehicleType) (*Slot, *errors.Error) {
	if f.IsFloorFull(vehicleType) {
		return nil, errors.ErrParkingLotFull
	}
	slot, err := f.GetFreeSlot(vehicleType)
	if err != nil {
		return nil, err
	}
	slot.SetSlotStatus(constants.SlotOccupied)
	f.AvailableSlots[vehicleType]--
	return slot, nil
}

func (f *Floor) UnparkVehicle(slot Slot, vehicleType constants.VehicleType) {
	for _, s := range f.Slots[vehicleType] {
		if s.SlotID == slot.SlotID {
			s.SetSlotStatus(constants.SlotAvailable)
			f.AvailableSlots[vehicleType]++
			break
		}
	}
}
