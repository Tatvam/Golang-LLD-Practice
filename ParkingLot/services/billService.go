package services

import (
	"ParkingLot/models"
	"ParkingLot/repositories"
)

type IBillService interface {
}

type BillService struct {
	billRepo          repositories.IBillRepository
	parkingLotService *ParkingLotService
	// Add Billing Strategy
}

func NewBillService(billRepo repositories.IBillRepository) *BillService {
	return &BillService{
		billRepo:          billRepo,
		parkingLotService: NewParkingLotService(),
	}
}

func (b *BillService) GenerateBill(vehicle models.Vehicle) {
	// Get ticket
	// Calculate amount
	// Insert bill
	// Return bill
}
