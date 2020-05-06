package controllers

import (
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
)

//TripController controller
type TripController struct {
	ctx dbcontext.DbContext
	Controller
}

//Resp teste
type Resp struct {
	value string
}

//NewTripController contructor
func NewTripController(ctx dbcontext.DbContext) TripController {
	return TripController{
		ctx: ctx,
	}
}

//Index get
func (c *TripController) Index(w http.ResponseWriter, r *http.Request) {
	resp := Resp{
		value: "200",
	}

	c.SendJSON(
		w,
		resp,
		http.StatusOK,
	)
}
