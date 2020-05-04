package repositories

import "github.com/NicolasDeveloper/tracker-microservices/pkg/models"

//IDeviceRepository repository
type IDeviceRepository interface {
	GetLoggedDevice(deviceID string) (models.Device, error)
	GetActiveDevice(deviceID string) (models.Device, error)
	CreateDevice(device models.Device) error
	Update(models.Device) error
}
