package services

import (
	"testing"
	"time"

	"github.com/NicolasDeveloper/tracker-microservices/internal/trip/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/dtos"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
)

func TestTrip(t *testing.T) {

	t.Run("Should call get address name", func(t *testing.T) {
		mapbox, error := NewMapBoxStub()
		repository, error := NewTripRepositoryStub()

		mapbox.On("GetAddressName", float64(10), float64(20)).Return("Teste endereço", nil)

		repository.On("Save").Return(nil)
		repository.On("GetOpenTrip", "123456789").Return(models.Trip{}, nil)

		service := NewTripService(&mapbox, &repository)

		service.Start(getTripDTO())

		result := mapbox.AssertCalled(t, "GetAddressName", float64(10), float64(20))

		if error != nil || result == false {
			t.Error("Save trip with adress empty")
		}
	})

	t.Run("Should save a trip", func(t *testing.T) {
		mapbox, error := NewMapBoxStub()
		repository, error := NewTripRepositoryStub()

		mapbox.On("GetAddressName", float64(10), float64(20)).Return("Teste endereço", nil)
		repository.On("Save").Return(nil)
		repository.On("GetOpenTrip", "123456789").Return(models.Trip{}, nil)

		service := NewTripService(&mapbox, &repository)

		service.Start(getTripDTO())

		result := repository.AssertCalled(t, "Save")

		if error != nil || result == false {
			t.Error("can't save trip")
		}
	})

}

func getTripDTO() dtos.TripDTO {
	return dtos.TripDTO{
		UserID:       "123456789",
		IdleTime:     0,
		SeqID:        1,
		TotalMileage: 2000,
		Tracks: []dtos.TrackDTO{
			{
				CurrentFuel:             0.5,
				CurrentTripMileage:      0.10,
				Latitude:                float64(10),
				Longitude:               float64(20),
				PositionSendingDateTime: time.Now().In(timeconvert.GetLocation()),
				RPM:                     2000,
				Speed:                   50,
				TotalTripMileage:        2000,
			},
		},
	}
}
