package models

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
	"github.com/beevik/guid"
)

func TestTrip(t *testing.T) {
	t.Run("Should calculate gps mileage", func(t *testing.T) {
		trip, error := newTrip()

		trip.AddTrack(getTrack(400, 70, 0, -23.6897796, -46.6433127))

		got := trip.TotalGpsMileage
		wantMin := float64(13)
		wantMax := float64(13.5)

		result := got >= wantMin && got <= wantMax

		if error != nil || result == false {
			t.Errorf("got %f want %f", got, wantMin)
		}
	})

	t.Run("Should create a trip identification", func(t *testing.T) {
		trip, err := newTrip()

		if err != nil || guid.IsGuid(trip.ID) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should create a trip with a identification", func(t *testing.T) {
		trip, err := newTrip()

		if err != nil || guid.IsGuid(trip.ID) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should thrown error when create a trip without a tracker identification", func(t *testing.T) {
		_, err := NewTrip("", "", Track{})

		want := "usuário sem identificação"

		if err == nil || err.Error() != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	})

	t.Run("Should thrown error when create a trip without a start coordenate", func(t *testing.T) {
		_, err := NewTrip("", "123", Track{AddressName: "56"})

		want := "deve informar a coordenada de partida"

		if err == nil || err.Error() != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	})

	t.Run("Should thrown error when create a trip without a address name", func(t *testing.T) {
		_, err := NewTrip("", "123", Track{Coordenate: Coordenate{Latitude: 1, Longitude: 2}})

		want := "deve informar o endereço de partida"

		if err == nil || err.Error() != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	})

	t.Run("Should thrown error when create a trip without a start location", func(t *testing.T) {
		_, err := NewTrip("", "123", Track{})

		want := "deve informar um ponto de partida"

		if err == nil || err.Error() != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	})

	t.Run("Should thrown error when resgiter location in closed trip", func(t *testing.T) {
		trip, error := newTrip()
		trip.Close()

		error = trip.AddTrack(Track{Coordenate: Coordenate{Latitude: 1, Longitude: 2}})
		want := "a viagem não foi encerrada"

		if error == nil || error.Error() != want {
			t.Errorf("got %q want %q", error, want)
		}
	})

	t.Run("Should thrown error when resgiter empty track", func(t *testing.T) {
		trip, error := newTrip()
		trip.Close()

		error = trip.AddTrack(Track{})
		want := "o track não pode estar vazio"

		if error == nil || error.Error() != want {
			t.Errorf("got %q want %q", error, want)
		}
	})

	t.Run("Should get last mileage", func(t *testing.T) {
		trip, error := newTrip()

		trip.AddTrack(getTrack(200, 60, 0, 0, 0))
		trip.AddTrack(getTrack(400, 70, 0, 0, 0))
		trip.AddTrack(getTrack(600, 75, 0, 0, 0))
		trip.AddTrack(getTrack(600, 30, 0, 0, 0))

		got := trip.TotalMileage
		want := float64(600)

		if error != nil || got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Should calc max speed", func(t *testing.T) {
		trip, error := newTrip()

		trip.AddTrack(getTrack(200, 60, 0, 0, 0))
		trip.AddTrack(getTrack(400, 70, 0, 0, 0))
		trip.AddTrack(getTrack(600, 75, 0, 0, 0))
		trip.AddTrack(getTrack(600, 30, 0, 0, 0))

		got := trip.MaxSpeed
		want := float64(75)

		if error != nil || got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Should calc max rpm", func(t *testing.T) {
		trip, error := newTrip()

		trip.AddTrack(getTrack(200, 60, 60, 0, 0))
		trip.AddTrack(getTrack(400, 70, 90, 0, 0))
		trip.AddTrack(getTrack(600, 75, 75, 0, 0))
		trip.AddTrack(getTrack(600, 30, 30, 0, 0))

		got := trip.MaxRPM
		want := float64(90)

		if error != nil || got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Should register end location", func(t *testing.T) {
		got, error := newTrip()
		got.AddTrack(Track{Coordenate: Coordenate{Latitude: 1, Longitude: 1}, AddressName: "Teste de endereço final"})
		got.Close()

		want := Track{Coordenate: Coordenate{Latitude: 1, Longitude: 1}, AddressName: "Teste de endereço final"}

		if error != nil || reflect.DeepEqual(got.EndTrack, want) == false {
			t.Error("object not equals")
		}
	})

	t.Run("Should calculate trip time", func(t *testing.T) {
		trip := Trip{
			StartAt:  time.Now().In(timeconvert.GetLocation()),
			ClosedAt: time.Now().In(timeconvert.GetLocation()).Add(time.Duration(30) * time.Minute),
		}

		got := trip.Time()
		wantMax := float64(31)
		wantMin := float64(30)

		result := got >= wantMin && got <= wantMax

		if result == false {
			t.Errorf("got %f want %f", got, wantMin)
		}
	})

	t.Run("Should be opened when start", func(t *testing.T) {
		trip, error := newTrip()

		got := trip.IsOpen()
		want := true

		if error != nil || got != want {
			t.Errorf("got %q want %q", strconv.FormatBool(got), strconv.FormatBool(want))
		}
	})
}

func TestTripEvents(t *testing.T) {
	t.Run("Should start a trip", func(t *testing.T) {
		got, error := newTrip()
		startAt := got.StartAt

		if error != nil || dateIsCloseFromNow(startAt) == false {
			t.Error("trip could not start")
		}
	})

	t.Run("Should close a trip", func(t *testing.T) {
		got, error := newTrip()
		got.Close()

		closeAt := got.ClosedAt

		if error != nil || dateIsCloseFromNow(closeAt) == false {
			t.Error("trip could not end")
		}
	})
}

func getTrack(currentTripMileage float64, speed float64, rpm float64, lat float64, long float64) Track {
	track, _ := NewTrack(Coordenate{Latitude: lat, Longitude: long}, 0, currentTripMileage, 0, time.Now(), speed, rpm, "Teste de endereço")
	return track
}

func dateIsCloseFromNow(date time.Time) bool {
	now := time.Now().In(timeconvert.GetLocation())
	diff := now.Sub(date)

	return diff <= 1
}

func newTrip() (Trip, error) {
	a, b := NewTrip("", "123456", getTrack(200, 60, 0, -23.6898722, -46.6432278))
	return a, b
}
