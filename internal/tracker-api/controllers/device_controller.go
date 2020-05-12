package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/common"
	devicedtos "github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/dtos"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
	sharedrepositories "github.com/NicolasDeveloper/tracker-microservices/pkg/repositories"
	"github.com/gorilla/mux"
)

//DeviceController model
type DeviceController struct {
	deviceRepository sharedrepositories.IDeviceRepository
	ctx              dbcontext.DbContext
	common.Controller
}

//NewDeviceController constructor
func NewDeviceController(ctx dbcontext.DbContext) DeviceController {
	deviceRepository, _ := sharedrepositories.NewDeviceRepository(ctx)

	return DeviceController{
		ctx:              ctx,
		deviceRepository: deviceRepository,
	}
}

//DeviceIsPluged knows if device is plugged
func (c *DeviceController) DeviceIsPluged(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pinCode := params["pincode"]

	device, err := c.deviceRepository.GetActiveDeviceByPINCode(pinCode)

	if err != nil {
		c.HandleError(err, w)
		return
	}

	response := devicedtos.ResponseDTO{
		Result: true,
		Content: devicedtos.DevicePlugedDTO{
			IsPlugged: device.Plugged,
		},
	}

	c.SendJSON(
		w,
		response,
		http.StatusOK,
	)
}

//GetUserDevice get user device
func (c *DeviceController) GetUserDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pinCode := params["pincode"]

	device, err := c.deviceRepository.GetActiveDeviceByPINCode(pinCode)

	if err != nil {
		c.HandleError(err, w)
		return
	}

	response := devicedtos.ResponseDTO{
		Result: true,
		Content: devicedtos.UserDeviceDTO{
			UserID: device.UserID,
		},
	}

	c.SendJSON(
		w,
		response,
		http.StatusOK,
	)
}
