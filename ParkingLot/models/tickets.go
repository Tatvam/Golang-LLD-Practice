package models

type Ticket struct {
	TicketID  int
	Slot      Slot
	Vehicle   Vehicle
	EntryTime int64
	Gate      Gate
}

func NewTicket(ticketID int, slot Slot, vehicle Vehicle, entryTime int64, gate Gate) Ticket {
	return Ticket{
		TicketID:  ticketID,
		Slot:      slot,
		Vehicle:   vehicle,
		EntryTime: entryTime,
		Gate:      gate,
	}
}
