package repositories

import "ParkingLot/models"

type BillRepository struct {
	bills map[int]models.Bill // map of billID to bill
}
