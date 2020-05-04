package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-communication/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/timeconvert"
)

func TestReport(t *testing.T) {
	t.Run("Should parse alarm report package", func(t *testing.T) {
		buffer := []byte{15, 0, 0, 0, 159, 216, 110, 94, 57, 221, 110, 94, 207, 58, 2, 0, 42, 83, 0, 0, 138, 5, 0, 0, 155, 0, 0, 0, 4, 0, 3, 57, 0, 87, 9, 0, 2, 0, 1, 16, 3, 20, 1, 58, 18, 110, 229, 16, 5, 180, 175, 7, 10, 198, 6, 77, 5, 140, 1, 0, 1, 59, 0, 60, 0, 203, 65, 13, 10}

		report := NewReport()

		got, error := report.GetAlarmReport(buffer)

		want := models.AlarmReport{
			AlarmSeq: 15,
			StateData: models.StateData{
				LastAcconTime:      time.Date(2020, time.March, 16, 1, 38, 39, 0, timeconvert.GetLocation()),
				UTCTime:            time.Date(2020, time.March, 16, 1, 58, 17, 0, timeconvert.GetLocation()),
				CurrentFuel:        1.55,
				CurrentTripMileage: 21290,
				TotalFuel:          14.18,
				TotalTripMileage:   146127,
			},
			GpsData: models.GpsData{
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
			},
			AlarmCount: 1,
			AlarmArray: []models.AlarmData{
				{
					NewAlarmFlag:     "0x00",
					AlarmType:        "0x01",
					AlarmDescription: "0x003b",
					AlarmThreshold:   "0x003c",
				},
			},
		}

		if error != nil || reflect.DeepEqual(want, got) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should parse gps report package", func(t *testing.T) {
		buffer := []byte{0, 160, 135, 85, 94, 32, 140, 85, 94, 205, 10, 0, 0, 199, 16, 0, 0, 19, 0, 0, 0, 59, 0, 0, 0, 4, 0, 3, 55, 0, 87, 0, 0, 0, 0, 1, 25, 2, 0, 21, 5, 34, 36, 239, 20, 5, 100, 176, 2, 10, 0, 0, 79, 2, 188, 1, 170, 3, 202, 50, 13, 10}

		report := NewReport()

		got, error := report.GetGpsReport(buffer)

		want := models.GpsReport{
			Flag: "0x00",
			StateData: models.StateData{
				LastAcconTime:      time.Date(2020, time.February, 25, 20, 46, 24, 0, timeconvert.GetLocation()),
				UTCTime:            time.Date(2020, time.February, 25, 21, 05, 36, 0, timeconvert.GetLocation()),
				TotalTripMileage:   2765,
				CurrentTripMileage: 4295,
				TotalFuel:          0.19,
				CurrentFuel:        0.59,
			},
			GpsData: models.GpsData{
				Count: 1,
				GpsItems: []models.GpsItem{
					{
						DateTime:  time.Date(2000, time.February, 25, 21, 05, 34, 0, timeconvert.GetLocation()),
						Latitude:  -23.682783333333333,
						Longitude: -46.65233,
						Speed:     0,
						Direction: 59.1,
						Flag:      "0xbc",
					},
				},
			},
			RpmData: models.RpmData{
				Count: 1,
				Rpm: []models.RpmItem{
					{
						Rpm: 938,
					},
				},
			},
		}

		if error != nil || reflect.DeepEqual(want, got) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should get first package message", func(t *testing.T) {
		buffer := []byte{64, 64, 89, 0, 4, 50, 49, 51, 71, 68, 80, 50, 48, 49, 56, 48, 50, 50, 52, 50, 49, 0, 0, 0, 0, 64, 1, 0, 160, 135, 85, 94, 32, 140, 85, 94, 205, 10, 0, 0, 199, 16, 0, 0, 19, 0, 0, 0, 59, 0, 0, 0, 4, 0, 3, 55, 0, 87, 0, 0, 0, 0, 1, 25, 2, 0, 21, 5, 34, 36, 239, 20, 5, 100, 176, 2, 10, 0, 0, 79, 2, 188, 1, 170, 3, 202, 50, 13, 10}

		report := NewReport()

		got, error := report.GetFirstMessagePack(buffer)
		_, error = report.GetGpsReport(buffer)
		got, error = report.GetSecondMessagePack(buffer, got)

		want := models.Message{
			Head:       "0x4040",
			Length:     89,
			Version:    "0x04",
			DeviceID:   "213GDP2018022421",
			ProtocolID: "0x4001",
			CRC:        "0xca32",
			Tail:       "0x0d0a",
		}

		if error != nil || reflect.DeepEqual(want, got) == false {
			t.Errorf("object not equals")
		}
	})

	t.Run("Should parse query package", func(t *testing.T) {
		buffer := []byte{64, 64, 58, 0, 4, 50, 49, 51, 71, 68, 80, 50, 48, 49, 56, 48, 50, 50, 52, 50, 49, 0, 0, 0, 0, 160, 2, 0, 0, 1, 0, 0, 1, 1, 32, 17, 0, 57, 66, 87, 65, 66, 52, 53, 90, 88, 72, 52, 48, 49, 48, 54, 54, 57, 60, 25, 13}
		buffer = sliceBuffer(buffer)
		report := NewReport()
		got, err := report.GetQuery(buffer)

		want := models.Query{
			CmdSeq:       0,
			RespCount:    1,
			FailCount:    0,
			FailTagArray: []string{},
			RespIndex:    0,
			SuccessCount: 1,
			SuccessTLVArray: []models.TLV{
				{
					Tag:        "0x2001",
					Length:     17,
					ValueArray: []byte{57, 66, 87, 65, 66, 52, 53, 90, 88, 72, 52, 48, 49, 48, 54, 54, 57},
				},
			},
		}

		if err != nil || reflect.DeepEqual(got, want) == false {
			t.Errorf("object not equals")
		}
	})
}

func sliceBuffer(bufferpack []byte) []byte {
	size := len(bufferpack) - 1
	return bufferpack[27:size]
}
