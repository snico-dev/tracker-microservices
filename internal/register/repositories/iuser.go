package repositories

import "github.com/NicolasDeveloper/tracker-microservices/internal/register/models"

//IUserRepository interface
type IUserRepository interface {
	GetByPINCode(pinCode string) (models.User, error)
	GetByID(deviceID string) (models.User, error)
	Create(user models.User) error
}
