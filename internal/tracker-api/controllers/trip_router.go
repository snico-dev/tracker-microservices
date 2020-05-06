package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/core"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
)

// TripRouter handle trip resources
type TripRouter struct {
	routes []core.Route
}

// NewTripRouter instance
func NewTripRouter(ctx dbcontext.DbContext) core.Bundle {
	ctrl := NewTripController(ctx)

	r := []core.Route{
		{
			Method:  http.MethodGet,
			Path:    "/trips",
			Handler: ctrl.Index,
		},
		// core.Route{
		// 	Method:  http.MethodPost,
		// 	Path:    "/kitties",
		// 	Handler: kc.Create,
		// },
		// core.Route{
		// 	Method:  http.MethodDelete,
		// 	Path:    "/kitties/{id}",
		// 	Handler: kc.Delete,
		// },
	}

	return &TripRouter{
		routes: r,
	}
}

// GetRoutes implement interface core.Bundle
func (b *TripRouter) GetRoutes() []core.Route {
	return b.routes
}
