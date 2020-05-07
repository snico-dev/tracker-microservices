package services

import (
	"reflect"
	"testing"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-communication/models"
	sharedmodels "github.com/NicolasDeveloper/tracker-microservices/pkg/models"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/repositories"
)

type deviceRepositoryStub struct {
}

func newDeviceRepositoryStub() repositories.IDeviceRepository {
	return &deviceRepositoryStub{}
}

func (repo *deviceRepositoryStub) GetLoggedDevice(id string) (sharedmodels.Device, error) {
	return sharedmodels.Device{}, nil
}

//GetActiveDevice search for device in database
func (repo *deviceRepositoryStub) GetActiveDevice(deviceID string) (sharedmodels.Device, error) {
	return sharedmodels.Device{}, nil
}

//GetActiveDeviceByCode search for device in database
func (repo *deviceRepositoryStub) GetActiveDeviceByCode(deviceID string) (sharedmodels.Device, error) {
	return sharedmodels.Device{}, nil
}

//GetActiveDeviceByPINCode search for device in database
func (repo *deviceRepositoryStub) GetActiveDeviceByPINCode(pinCode string) (sharedmodels.Device, error) {
	return sharedmodels.Device{}, nil
}

//CreateDevice search for device in database
func (repo *deviceRepositoryStub) CreateDevice(device sharedmodels.Device) error {
	return nil
}

//Update update device
func (repo *deviceRepositoryStub) Update(device sharedmodels.Device) error {
	return nil
}

func TestTrack(t *testing.T) {
	t.Run("Should parse alarm report package", func(t *testing.T) {
		buffer := []byte{64, 64, 58, 0, 4, 50, 49, 51, 71, 68, 80, 50, 48, 49, 56, 48, 50, 50, 52, 50, 49, 0, 0, 0, 0, 160, 2, 0, 0, 1, 0, 0, 1, 1, 32, 17, 0, 57, 66, 87, 65, 66, 52, 53, 90, 88, 72, 52, 48, 49, 48, 54, 54, 57, 60, 25, 13}

		service := NewTrackService(newDeviceRepositoryStub())
		got, err := service.ParseQueryReport(buffer)
		want := models.TLVDescription{
			VinCode: "9BWAB45ZXH4010669",
		}
		if err != nil || reflect.DeepEqual(got, want) == false {
			t.Errorf("object not equals")
		}
	})
}
