package dtos

import "time"

//TripDTO data transfer object
type TripDTO struct {
	UserID       string
	DeviceID     string
	SeqID        int
	TotalMileage float64
	IdleTime     int
	Tracks       []TrackDTO
}

//TrackDTO data transfer object
type TrackDTO struct {
	Latitude                float64
	Longitude               float64
	TotalTripMileage        float64
	CurrentTripMileage      float64
	PositionSendingDateTime time.Time
	CurrentFuel             float64
	Speed                   float64
	RPM                     float64
}

//GetFirstTrack first track
func (trip *TripDTO) GetFirstTrack() TrackDTO {
	return trip.Tracks[0]
}
