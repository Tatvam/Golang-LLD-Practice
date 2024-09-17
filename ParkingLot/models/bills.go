package models

import "ParkingLot/constants"

type Bill struct {
	BillID     int
	Ticket     Ticket
	Amount     float64
	BillStatus constants.BillStatus
	Gate       Gate
	ExitTime   int64
}
