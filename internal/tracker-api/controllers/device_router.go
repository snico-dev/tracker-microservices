package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/common"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
)

//DeviceRouter handle device resources
type DeviceRouter struct {
	routes []common.Route
}

// NewDeviceRouter instance
func NewDeviceRouter(ctx dbcontext.DbContext) common.Bundle {
	ctrl := NewDeviceController(ctx)

	r := []common.Route{
		{
			Method:  http.MethodGet,
			Path:    "/devices/{pincode}/is-plugged",
			Handler: ctrl.DeviceIsPluged,
		},
		{
			Method:  http.MethodGet,
			Path:    "/devices/{pincode}/user",
			Handler: ctrl.GetUserDevice,
		},
	}

	return &DeviceRouter{
		routes: r,
	}
}

// GetRoutes implement interface common.Bundle
func (b *DeviceRouter) GetRoutes() []common.Route {
	return b.routes
}
