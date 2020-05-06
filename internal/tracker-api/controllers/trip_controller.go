package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/common"
	triprepositories "github.com/NicolasDeveloper/tracker-microservices/internal/trip/repositories"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
)

//TripController controller
type TripController struct {
	ctx dbcontext.DbContext
	common.Controller
}

//NewTripController contructor
func NewTripController(ctx dbcontext.DbContext) TripController {
	return TripController{
		ctx: ctx,
	}
}

//Index get
func (c *TripController) Index(w http.ResponseWriter, r *http.Request) {
	tripRepository, err := triprepositories.NewTripRepository(c.ctx)

	if err != nil {
		c.HandleError(err, w)
		return
	}

	trips, err := tripRepository.GetTripsByUser("f81479b5-0aa4-4619-8b2c-7ad583ab4e3c")

	if err != nil {
		c.HandleError(err, w)
		return
	}

	c.SendJSON(
		w,
		trips,
		http.StatusOK,
	)
}
