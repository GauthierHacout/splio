package core

import (
	"fmt"
	"time"
)

/* Every minute the data fetched though the velib API will be outputed in the console */
func HandleConsole() {
	ticker := time.NewTicker(1 * time.Minute)
	for _ = range ticker.C {
		response := GetVelibRecords()
		fmt.Print("\n\n----INFORMATION about velibs stations in a 500 meters radius around Splio HQ----")
		fmt.Println("Last Update At : ", response.Records[0].Timestamp)
		for _, rec := range response.Records {
			fields := rec.Fields
			fmt.Println("\n-- Station : ", fields.Name, " --")
			fmt.Println("Capacity : ", fields.Capacity, " | Distance From Splio (m) : ",
				fields.DistanceFrom, " | Rental is possible :", fields.AcceptRental)
			fmt.Println("Number of : Bikes availables(", fields.NbrOfAvailableBike,") | Electrical bikes (",
				fields.NbrOfElectricBike,") | Mecanichal bikes(",fields.NbrOfMechanicalBike,") | Docks available : (",fields.NbrOfDocksAvailable,")")
		}
	}
}

