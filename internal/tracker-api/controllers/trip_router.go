package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/common"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
)

// TripRouter handle trip resources
type TripRouter struct {
	routes []common.Route
}

// NewTripRouter instance
func NewTripRouter(ctx dbcontext.DbContext) common.Bundle {
	ctrl := NewTripController(ctx)

	r := []common.Route{
		{
			Method:  http.MethodGet,
			Path:    "/trips/{userid}/user",
			Handler: ctrl.Index,
		},
	}

	return &TripRouter{
		routes: r,
	}
}

// GetRoutes implement interface common.Bundle
func (b *TripRouter) GetRoutes() []common.Route {
	return b.routes
}
