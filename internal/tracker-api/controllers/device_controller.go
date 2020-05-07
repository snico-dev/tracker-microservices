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
	ctx dbcontext.DbContext
	common.Controller
}

//NewDeviceController constructor
func NewDeviceController(ctx dbcontext.DbContext) DeviceController {
	return DeviceController{
		ctx: ctx,
	}
}

//DeviceIsPluged knows if device is plugged
func (c *DeviceController) DeviceIsPluged(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pinCode := params["pincode"]

	deviceRepository, err := sharedrepositories.NewDeviceRepository(c.ctx)
	device, err := deviceRepository.GetActiveDeviceByPINCode(pinCode)

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
