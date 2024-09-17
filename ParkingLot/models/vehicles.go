package models

import "ParkingLot/constants"

type Vehicle struct {
	VehicleNumber string
	VehicleType   constants.VehicleType
}

func NewVehicle(vehicleNumber string, vehicleType constants.VehicleType) Vehicle {
	return Vehicle{
		VehicleNumber: vehicleNumber,
		VehicleType:   vehicleType,
	}
}
