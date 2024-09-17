package repositories

import "ParkingLot/models"

type TicketRepository struct {
	tickets map[string]*models.Ticket // map of vehicle number to ticket
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{
		tickets: make(map[string]*models.Ticket),
	}
}

func (t *TicketRepository) InsertTicket(vehicleNumber string, ticket *models.Ticket) {
	t.tickets[vehicleNumber] = ticket
}

func (t *TicketRepository) GetTicket(vehicleNumber string) *models.Ticket {
	return t.tickets[vehicleNumber]
}
