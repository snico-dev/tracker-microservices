package services

import (
	"errors"

	"github.com/NicolasDeveloper/tracker-microservices/internal/trip/acls"
	"github.com/NicolasDeveloper/tracker-microservices/internal/trip/models"
	"github.com/NicolasDeveloper/tracker-microservices/internal/trip/repositories"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/dtos"
)

//ITripService interface
type ITripService interface {
	Start(dto dtos.TripDTO) error
	Increment(dto dtos.TripDTO) error
	Close(dto dtos.TripDTO) error
}

//NewTripService contructor
func NewTripService(mapACL acls.IMapACL, tripRepository repositories.ITripRepository) ITripService {
	return &tripService{
		mapACL:         mapACL,
		tripRepository: tripRepository,
	}
}

type tripService struct {
	mapACL         acls.IMapACL
	tripRepository repositories.ITripRepository
}

func (service *tripService) Start(dto dtos.TripDTO) error {
	startTrack := dto.GetFirstTrack()

	coordenate := models.Coordenate{
		Latitude:  startTrack.Latitude,
		Longitude: startTrack.Longitude,
	}

	addressName, err := service.mapACL.GetAddressName(coordenate.Latitude, coordenate.Longitude)

	track, err := models.NewTrack(
		coordenate,
		startTrack.CurrentFuel,
		startTrack.CurrentTripMileage,
		startTrack.TotalTripMileage,
		startTrack.PositionSendingDateTime,
		startTrack.Speed,
		startTrack.RPM,
		addressName,
	)

	if err != nil {
		return err
	}

	if dto.UserID == "" {
		return errors.New("user identification not found")
	}

	trip, err := service.tripRepository.GetOpenTrip(dto.UserID)

	if err != nil || trip.ID != "" {
		return errors.New("has a trip in progress")
	}

	trip, err = models.NewTrip(dto.DeviceID, dto.UserID, track)

	if err != nil {
		return err
	}

	return service.tripRepository.Save(trip)
}

func (service *tripService) Increment(dto dtos.TripDTO) error {
	currentTrack := dto.GetFirstTrack()

	coordenate := models.Coordenate{
		Latitude:  currentTrack.Latitude,
		Longitude: currentTrack.Longitude,
	}

	trip, err := service.tripRepository.GetOpenTrip(dto.UserID)

	if err != nil ||
		trip.ID == "" ||
		trip.UserID == "" {
		return errors.New("trip not found")
	}

	track, err := models.NewTrack(
		coordenate,
		currentTrack.CurrentFuel,
		currentTrack.CurrentTripMileage,
		currentTrack.TotalTripMileage,
		currentTrack.PositionSendingDateTime,
		currentTrack.Speed,
		currentTrack.RPM,
		"",
	)

	err = trip.AddTrack(track)
	tracks := []models.Track{track}

	if err != nil {
		return err
	}

	return service.tripRepository.UpdateTracks(trip, tracks)
}

func (service *tripService) Close(dto dtos.TripDTO) error {
	currentTrack := dto.GetFirstTrack()

	coordenate := models.Coordenate{
		Latitude:  currentTrack.Latitude,
		Longitude: currentTrack.Longitude,
	}

	trip, err := service.tripRepository.GetOpenTrip(dto.UserID)

	if err != nil || trip.ID == "" {
		return errors.New("trip not found")
	}

	addressName, err := service.mapACL.GetAddressName(coordenate.Latitude, coordenate.Longitude)
	track, err := models.NewTrack(
		coordenate,
		currentTrack.CurrentFuel,
		currentTrack.CurrentTripMileage,
		currentTrack.TotalTripMileage,
		currentTrack.PositionSendingDateTime,
		currentTrack.Speed,
		currentTrack.RPM,
		addressName,
	)

	err = trip.AddTrack(track)
	err = trip.Close()
	tracks := []models.Track{track}

	if err != nil {
		return err
	}

	return service.tripRepository.CloseTrip(trip, tracks)
}
