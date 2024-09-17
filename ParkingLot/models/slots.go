package models

import "ParkingLot/constants"

type Slot struct {
	SlotID       int
	SlotName     string
	VechicleType constants.VehicleType
	SlotStatus   constants.SlotStatus
	FloorID      int
}

func NewSlots(FloorID int, NumOfSlots int, vehicleType constants.VehicleType) []Slot {
	slots := []Slot{}
	for i := 0; i < NumOfSlots; i++ {
		slots = append(slots, Slot{
			SlotID:       i,
			SlotName:     "Slot" + string(i),
			VechicleType: vehicleType,
			SlotStatus:   constants.SlotAvailable,
			FloorID:      FloorID,
		})
	}
	return slots
}

func NewSlot(FloorID int, SlotType constants.VehicleType) Slot {
	// Change SlotID generation logic
	return Slot{
		SlotID:       0,
		SlotName:     "Slot" + string(0),
		VechicleType: SlotType,
		SlotStatus:   constants.SlotAvailable,
		FloorID:      FloorID,
	}
}

func (s *Slot) SetSlotStatus(status constants.SlotStatus) {
	s.SlotStatus = status
}

func (s *Slot) GetSlotStatus() constants.SlotStatus {
	return s.SlotStatus
}
