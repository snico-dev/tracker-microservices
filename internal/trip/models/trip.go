package models

import (
	"errors"
	"reflect"
	"time"

	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
	"github.com/beevik/guid"
)

//Trip data type
type Trip struct {
	ID                string     `bson:"_id" json:"id"`
	UserID            string     `bson:"user_id" json:"user_id"`
	DeviceID          string     `bson:"device_id" json:"device_id"`
	SeqID             string     `bson:"seq_id" json:"seq_id"`
	StartAt           time.Time  `bson:"start_at" json:"start_at"`
	UpdateAt          time.Time  `bson:"update_at" json:"update_at"`
	ClosedAt          time.Time  `bson:"closed_at" json:"closed_at"`
	Open              bool       `bson:"open" json:"open"`
	TotalGpsMileage   float64    `bson:"total_gps_mileage" json:"total_gps_mileage"`
	TotalMileage      float64    `bson:"total_mileage" json:"total_mileage"`
	MaxSpeed          float64    `bson:"max_speed" json:"max_speed"`
	MaxRPM            float64    `bson:"max_rpm" json:"max_rpm"`
	IdleTime          float64    `bson:"idle_time" json:"idle_time"`
	CurrentCoordinate Coordenate `bson:"current_coordinate" json:"current_coordinate"`
	StartTrack        Track      `bson:"start_track" json:"start_track"`
	EndTrack          Track      `bson:"end_track" json:"end_track"`
	Tracks            []Track    `bson:"tracks" json:"tracks"`
}

//NewTrip constructor
func NewTrip(
	deviceID string,
	userID string,
	startTrack Track) (Trip, error) {
	if len(userID) == 0 {
		return Trip{}, errors.New("usuário sem identificação")
	}

	if reflect.DeepEqual(startTrack, Track{}) == true {
		return Trip{}, errors.New("deve informar um ponto de partida")
	}

	if reflect.DeepEqual(startTrack.Coordenate, Coordenate{}) == true {
		return Trip{}, errors.New("deve informar a coordenada de partida")
	}

	if len(startTrack.AddressName) == 0 {
		return Trip{}, errors.New("deve informar o endereço de partida")
	}

	guidid := guid.New()
	return Trip{
		ID:                guidid.String(),
		UserID:            userID,
		DeviceID:          deviceID,
		Tracks:            []Track{startTrack},
		Open:              true,
		MaxRPM:            0,
		MaxSpeed:          0,
		StartTrack:        startTrack,
		CurrentCoordinate: startTrack.Coordenate,
		StartAt:           time.Now().In(timeconvert.GetLocation()),
		TotalGpsMileage:   float64(0),
		TotalMileage:      float64(0),
	}, nil
}

//Close entity behavor
func (trip *Trip) Close() error {

	if trip.Open == false {
		return errors.New("trip já está fechada")
	}

	trip.Open = false
	trip.ClosedAt = time.Now().In(timeconvert.GetLocation())

	endTrack := trip.Tracks[len(trip.Tracks)-1]

	trip.RegisterEndTrack(endTrack)

	return nil
}

//IsClosed entity behavor
func (trip *Trip) IsClosed() bool {
	return !trip.Open
}

//IsOpen entity behavor
func (trip *Trip) IsOpen() bool {
	return trip.Open
}

//IsFirstLocation entity behavor
func (trip *Trip) IsFirstLocation() bool {
	return len(trip.Tracks) == 0
}

//GetDistanceInMilageFrom get distance
func (trip *Trip) GetDistanceInMilageFrom(track Track) float64 {
	return trip.Tracks[len(trip.Tracks)-1].GetDistanceInMilageFrom(track.Coordenate)
}

//AddTrack entity behavor
func (trip *Trip) AddTrack(track Track) error {
	if reflect.DeepEqual(track, Track{}) == true {
		return errors.New("o track não pode estar vazio")
	}

	if trip.Open == false {
		return errors.New("a viagem não foi encerrada")
	}

	trip.TotalGpsMileage += trip.GetDistanceInMilageFrom(track)
	trip.TotalMileage = track.CurrentTripMileage
	trip.CurrentCoordinate = track.Coordenate
	trip.UpdateAt = time.Now().In(timeconvert.GetLocation())
	trip.Tracks = append(trip.Tracks, track)

	if track.Speed > trip.MaxSpeed {
		trip.MaxSpeed = track.Speed
	}

	if track.RPM > trip.MaxRPM {
		trip.MaxRPM = track.RPM
	}

	return nil
}

//RegisterEndTrack entity behavor
func (trip *Trip) RegisterEndTrack(track Track) error {
	if reflect.DeepEqual(track, Track{}) == true {
		return errors.New("o track não pode ser vazia")
	}

	trip.EndTrack = track
	return nil
}

//RegisterStartTrack entity behavor
func (trip *Trip) RegisterStartTrack(track Track) error {
	if reflect.DeepEqual(track, Track{}) == true {
		return errors.New("o track não pode ser vazia")
	}

	trip.StartTrack = track
	return nil
}

//Time trip duration
func (trip *Trip) Time() float64 {
	diff := trip.ClosedAt.Sub(trip.StartAt)
	return diff.Minutes()
}
