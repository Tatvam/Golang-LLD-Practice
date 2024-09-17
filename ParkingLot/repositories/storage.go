package repositories

import "ParkingLot/models"

type ITicketRepository interface {
	InsertTicket(string, *models.Ticket)
	GetTicket(vehicleNumber string) *models.Ticket
}

type IBillRepository interface {
	GenerateBill() (int, error)
	InsertBill(billID int, ticketID int, amount int) error
	GetBill(billID int) (int, int, error)
}
