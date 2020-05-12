package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/common"
	triprepositories "github.com/NicolasDeveloper/tracker-microservices/internal/trip/repositories"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
	"github.com/gorilla/mux"
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
	params := mux.Vars(r)
	userid := params["userid"]

	tripRepository, err := triprepositories.NewTripRepository(c.ctx)

	if err != nil {
		c.HandleError(err, w)
		return
	}

	trips, err := tripRepository.GetTripsByUser(userid)

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
