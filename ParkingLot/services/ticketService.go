package services

import (
	"ParkingLot/errors"
	"ParkingLot/models"
	"ParkingLot/repositories"
	"time"
)

type ITicketService interface {
	GenerateTicket(models.Vehicle, int) (models.Ticket, error)
}

type TicketService struct {
	ticketRepo        repositories.ITicketRepository
	parkingLotService *ParkingLotService
}

func NewTicketService(ticketRepo repositories.ITicketRepository) *TicketService {
	return &TicketService{
		ticketRepo:        ticketRepo,
		parkingLotService: NewParkingLotService(),
	}
}

func (t *TicketService) GenerateTicket(vehicle models.Vehicle, gateID int) (*models.Ticket, *errors.Error) {
	gate, err := t.parkingLotService.GetEntryGate(gateID)
	if err != nil {
		return nil, err
	}
	slot, err := t.parkingLotService.ParkVehicle(vehicle)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	ticket := models.NewTicket(1, *slot, vehicle, now, *gate)
	t.ticketRepo.InsertTicket(vehicle.VehicleNumber, &ticket)
	return &ticket, nil
}
