package model

import (
	"log"
	"time"
)

/* These struct are used to represent the data we will get from Velib API in Go */
type Response struct {
	Records []Record	`json:"records"`
}

type Record struct {
	Timestamp 		string 		`json:"record_timestamp"`
	Fields			Field 		`json:"fields"`
}

type Field struct {
	NbrOfElectricBike 		int 		`json:"ebike"`
	NbrOfMechanicalBike 	int 		`json:"mechanical"`
	NbrOfAvailableBike 		int 		`json:"numbikesavailable"`
	Capacity 				int 		`json:"capacity"`
	Name 					string		`json:"name"`
	DistanceFrom 			string 		`json:"dist"`
	NbrOfDocksAvailable 	int 		`json:"numdocksavailable"`
	AcceptRental 			string 		`json:"is_renting"`
	AcceptReturn 			string 		`json:"is_returning"`
	Coordinates 			[]float64 	`json:"coordonnees_geo"`
}

/*  Transforming the string corresponding to a timestamp into a human readable time (HH:mm:ss)
 (e.g "2020-12-12T17:39:06+00:00" => "18:23:40") */
func(r *Record) FormatTimestamp() {
	layout := "2006-01-02T15:04:05.000000+01:00"
	t, err := time.Parse(layout, r.Timestamp)
	if err != nil {
		log.Print(err)
	}

	(*r).Timestamp = t.Format( "15:04:05")
}

/* Transforming the string corresponding to the distance to avoid having more than 4 numbers
(e.g "106.3758427632" => "106.3")*/
func (r *Record) FormatDistance() {
	(*r).Fields.DistanceFrom = (*r).Fields.DistanceFrom[:5]
}
