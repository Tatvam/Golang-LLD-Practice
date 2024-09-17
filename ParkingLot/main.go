package main

import (
	"ParkingLot/constants"
	"ParkingLot/models"
	"ParkingLot/repositories"
	"ParkingLot/services"
	"fmt"
)

func main() {
	// This is the main function
	ticketRepo := repositories.NewTicketRepository()
	ticketService := services.NewTicketService(ticketRepo)
	//billService = services.NewBillService()

	// Park Vehicle
	car1 := models.NewVehicle("KA01HH1234", constants.HEAVY_VEHICLE)
	ticket, err := ticketService.GenerateTicket(car1, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ticket)

	car2 := models.NewVehicle("KA01HH1232", constants.LIGHT_VEHICLE)
	ticket2, err := ticketService.GenerateTicket(car2, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ticket2)
}
