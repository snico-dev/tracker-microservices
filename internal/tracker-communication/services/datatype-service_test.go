package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-communication/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/buffer"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
)

func TestDataType(t *testing.T) {
	t.Run("Should interpret gps item data type", func(t *testing.T) {
		bufferpackpack := []byte{16, 3, 20, 1, 58, 18, 110, 229, 16, 5, 180, 175, 7, 10, 198, 6, 77, 5, 140, 1}

		service := NewDataType(buffer.NewBuffer(), timeconvert.NewUnixDate())
		got, error := service.GetGpsItem(bufferpackpack)

		want := models.GpsItem{
			DateTime:  time.Date(2020, time.March, 16, 1, 58, 18, 0, timeconvert.GetLocation()),
			Latitude:  -23.609275,
			Longitude: -46.74330333333333,
			Direction: 135.7,
			Speed:     62.42,
			Flag:      "0x8c",
		}

		if error != nil || reflect.DeepEqual(want, got) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should interpret gps item data type", func(t *testing.T) {
		bufferpack := []byte{1, 16, 3, 20, 1, 58, 18, 110, 229, 16, 5, 180, 175, 7, 10, 198, 6, 77, 5, 140, 1, 0, 1, 59, 0, 60, 0, 203, 65, 13, 10}

		service := NewDataType(buffer.NewBuffer(), timeconvert.NewUnixDate())
		got, error := service.GetGpsData(bufferpack)
		want := models.GpsData{
			Count: 1,
			GpsItems: []models.GpsItem{
				{
					DateTime:  time.Date(2020, time.March, 16, 1, 58, 18, 0, timeconvert.GetLocation()),
					Latitude:  -23.609275,
					Longitude: -46.74330333333333,
					Direction: 135.7,
					Speed:     62.42,
					Flag:      "0x8c",
				},
			},
		}

		if error != nil || reflect.DeepEqual(got, want) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should get state data", func(t *testing.T) {
		bufferpack := []byte{159, 216, 110, 94, 57, 221, 110, 94, 207, 58, 2, 0, 42, 83, 0, 0, 138, 5, 0, 0, 155, 0, 0, 0, 4, 0, 3, 57, 0, 87, 9, 0, 2, 0, 1, 16, 3, 20, 1, 58, 18, 110, 229, 16, 5, 180, 175, 7, 10, 198, 6, 77, 5, 140, 1, 0, 1, 59, 0, 60, 0, 203, 65, 13, 10}

		service := NewDataType(buffer.NewBuffer(), timeconvert.NewUnixDate())
		got, error := service.GetStateData(bufferpack)

		want := models.StateData{
			LastAcconTime:      time.Date(2020, time.March, 16, 1, 38, 39, 0, timeconvert.GetLocation()),
			UTCTime:            time.Date(2020, time.March, 16, 1, 58, 17, 0, timeconvert.GetLocation()),
			CurrentFuel:        1.55,
			CurrentTripMileage: 21290,
			TotalFuel:          14.18,
			TotalTripMileage:   146127,
		}

		if error != nil || reflect.DeepEqual(got, want) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should interpret rpm item", func(t *testing.T) {
		bufferpack := []byte{170, 3, 202, 50, 13, 10}

		service := NewDataType(buffer.NewBuffer(), timeconvert.NewUnixDate())
		got, error := service.GetRpmItem(bufferpack)
		want := models.RpmItem{
			Rpm: 938,
		}
		if error != nil || reflect.DeepEqual(got, want) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should interpret rpm data", func(t *testing.T) {
		bufferpack := []byte{1, 170, 3, 202, 50, 13, 10}

		service := NewDataType(buffer.NewBuffer(), timeconvert.NewUnixDate())
		got, error := service.GetRpmData(bufferpack)
		want := models.RpmData{
			Count: 1,
			Rpm: []models.RpmItem{
				{
					Rpm: 938,
				},
			},
		}
		if error != nil || reflect.DeepEqual(got, want) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should interpret alarm data", func(t *testing.T) {
		bufferpack := []byte{0, 1, 59, 0, 60, 0, 203, 65, 13, 10}

		service := NewDataType(buffer.NewBuffer(), timeconvert.NewUnixDate())
		got, error := service.GetAlarmData(bufferpack)
		want := models.AlarmData{
			NewAlarmFlag:     "0x00",
			AlarmType:        "0x01",
			AlarmDescription: "0x003b",
			AlarmThreshold:   "0x003c",
		}
		if error != nil || reflect.DeepEqual(got, want) == false {
			t.Errorf("object not equals")
		}
	})
}
