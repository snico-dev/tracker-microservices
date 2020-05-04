package services

import (
	"strings"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-communication/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/convert"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/dtos"
	sharedmodels "github.com/NicolasDeveloper/tracker-microservices/pkg/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/repositories"
)

//ITrackService service
type ITrackService interface {
	ParseGpsReport(bufferpack []byte) (dtos.TripDTO, error)
	ParseAlarmReport(bufferpack []byte) (dtos.TripDTO, error)
	ParseQueryReport(bufferpack []byte) (models.TLVDescription, error)
	IsIgnitionOnAlarm(bufferpack []byte) (bool, error)
	IsIgnitionOffAlarm(bufferpack []byte) (bool, error)
	GetLoggedDevice(deviceID string) (sharedmodels.Device, error)
	DoLogin(deviceID string) (bool, error)
	GetDeviceID(bufferpack []byte) (string, error)
}

type trackService struct {
	trackRepository repositories.IDeviceRepository
}

//NewTrackService contructor
func NewTrackService(trackRepository repositories.IDeviceRepository) ITrackService {
	return &trackService{
		trackRepository: trackRepository,
	}
}

func (service *trackService) ParseGpsReport(bufferpack []byte) (dtos.TripDTO, error) {
	repo := NewReport()

	message, err := repo.GetFirstMessagePack(bufferpack)
	reportData, err := repo.GetGpsReport(bufferpack)
	message, err = repo.GetSecondMessagePack(bufferpack, message)

	track := reportData.GpsData.GpsItems[0]
	stateData := reportData.StateData

	rpm := float64(0)
	if reportData.RpmData.Count > 0 {
		rpm = reportData.RpmData.Rpm[0].Rpm
	}

	trackDto := dtos.TrackDTO{
		CurrentFuel:             stateData.CurrentFuel,
		CurrentTripMileage:      stateData.CurrentTripMileage,
		Latitude:                track.Latitude,
		Longitude:               track.Longitude,
		PositionSendingDateTime: stateData.UTCTime,
		RPM:                     rpm,
		Speed:                   track.Speed,
	}
	return dtos.TripDTO{
		DeviceID:     message.DeviceID,
		TotalMileage: reportData.StateData.TotalTripMileage,
		IdleTime:     0,
		Tracks:       []dtos.TrackDTO{trackDto},
	}, err
}

func (service *trackService) ParseAlarmReport(bufferpack []byte) (dtos.TripDTO, error) {
	repo := NewReport()

	message, err := repo.GetFirstMessagePack(bufferpack)
	reportData, err := repo.GetAlarmReport(bufferpack)
	message, err = repo.GetSecondMessagePack(bufferpack, message)

	if err != nil {
		return dtos.TripDTO{}, err
	}

	track := models.GpsItem{}
	if reportData.GpsData.Count > 0 {
		track = reportData.GpsData.GpsItems[0]
	}

	stateData := reportData.StateData

	rpm := float64(0)

	trackDto := dtos.TrackDTO{
		CurrentFuel:             stateData.CurrentFuel,
		CurrentTripMileage:      stateData.CurrentTripMileage,
		Latitude:                track.Latitude,
		Longitude:               track.Longitude,
		PositionSendingDateTime: stateData.UTCTime,
		RPM:                     rpm,
		Speed:                   track.Speed,
	}

	return dtos.TripDTO{
		DeviceID:     message.DeviceID,
		TotalMileage: reportData.StateData.TotalTripMileage,
		IdleTime:     0,
		Tracks:       []dtos.TrackDTO{trackDto},
	}, err
}

func (service *trackService) IsIgnitionOnAlarm(bufferpack []byte) (bool, error) {
	repo := NewReport()

	reportData, err := repo.GetAlarmReport(bufferpack)

	if reportData.AlarmCount == 0 {
		return false, err
	}

	alarm := reportData.AlarmArray[0]

	if alarm.AlarmType == "0x16" {
		return true, nil
	}

	return false, err
}

func (service *trackService) IsIgnitionOffAlarm(bufferpack []byte) (bool, error) {
	repo := NewReport()

	reportData, err := repo.GetAlarmReport(bufferpack)

	if reportData.AlarmCount == 0 {
		return false, err
	}

	alarm := reportData.AlarmArray[0]

	if alarm.AlarmType == "0x17" {
		return true, nil
	}

	return false, err
}

func (service *trackService) ParseQueryReport(bufferpack []byte) (models.TLVDescription, error) {
	repo := NewReport()

	message, err := repo.GetFirstMessagePack(bufferpack)
	reportData, err := repo.GetQuery(bufferpack)
	message, err = repo.GetSecondMessagePack(bufferpack, message)

	tlvDescription := models.TLVDescription{}

	if err != nil {
		return tlvDescription, err
	}

	if reportData.SuccessCount > 0 {
		tlvs := reportData.SuccessTLVArray
		for _, tlv := range tlvs {
			if tlv.Tag == "0x2001" {
				value, err := convert.FromByteHexToASCII(tlv.ValueArray)
				if err == nil {
					tlvDescription.VinCode = strings.ToUpper(value)
				}
			}
		}
	}

	return tlvDescription, err
}

func (service *trackService) GetDeviceID(bufferpack []byte) (string, error) {
	repo := NewReport()
	message, err := repo.GetFirstMessagePack(bufferpack)

	if err != nil {
		return "", err
	}

	return message.DeviceID, nil
}

func (service *trackService) GetLoggedDevice(deviceID string) (sharedmodels.Device, error) {
	repo := service.trackRepository
	return repo.GetLoggedDevice(deviceID)
}

func (service *trackService) DoLogin(deviceID string) (bool, error) {
	repo := service.trackRepository

	device, err := repo.GetActiveDevice(deviceID)

	if err != nil {
		return false, err
	}

	err = device.DoLogin()

	if err != nil {
		return false, err
	}

	err = repo.Update(device)

	return true, err
}