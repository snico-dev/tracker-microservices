package models

import (
	"time"

	"github.com/NicolasDeveloper/tracker-microservices/pkg/geocoordenate"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
)

//Track model type
type Track struct {
	Coordenate              Coordenate `bson:"coordenate"`
	AddressName             string     `bson:"address_name"`
	CreateAt                time.Time  `bson:"create_at"`
	CurrentTripMileage      float64    `bson:"current_trip_mileage"`
	TotalTripMileage        float64    `bson:"total_trip_mileage"`
	CurrentFuel             float64    `bson:"current_fuel"`
	PositionSendingDateTime time.Time  `bson:"position_sending_date_time"`
	Speed                   float64    `bson:"speed"`
	RPM                     float64    `bson:"rpm"`
}

//NewTrack constructor
func NewTrack(
	coordenate Coordenate,
	currentFuel float64,
	currentTripMileage float64,
	totalTripMileage float64,
	positionSendingDateTime time.Time,
	speed float64,
	rpm float64,
	addressName string) (Track, error) {
	return Track{
		Coordenate:              coordenate,
		CreateAt:                time.Now().In(timeconvert.GetLocation()),
		CurrentFuel:             currentFuel,
		CurrentTripMileage:      currentTripMileage,
		TotalTripMileage:        totalTripMileage,
		PositionSendingDateTime: positionSendingDateTime,
		Speed:                   speed,
		RPM:                     rpm,
		AddressName:             addressName,
	}, nil
}

//GetDistanceInMilageFrom distance between two coordanates
func (t *Track) GetDistanceInMilageFrom(coordenate Coordenate) float64 {
	lat, long := t.Coordenate.Latitude, t.Coordenate.Longitude
	return geocoordenate.Distance(lat, long, coordenate.Latitude, coordenate.Longitude)
}
