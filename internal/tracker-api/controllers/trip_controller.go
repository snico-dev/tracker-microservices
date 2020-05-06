package controllers

import "github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/common"

//TripController controller
type tripController struct {
	common.Controller
}

func NewTripController() tripController {
	return &tripController{}
}
