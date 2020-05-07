package core

import (
	"log"
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/common"
	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/controllers"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
	"github.com/gorilla/mux"
)

//App initialize app
type App struct {
	router *mux.Router
	dbctx  dbcontext.DbContext
}

//NewApp contructor
func NewApp() *App {
	return &App{}
}

//Initialize configure app
func (a *App) Initialize() *App {
	a.dbctx = dbcontext.NewContext()
	a.dbctx.Connect()
	return a
}

//ConfigEndpoints endpoints
func (a *App) ConfigEndpoints() *App {
	a.router = mux.NewRouter()
	s := a.router.PathPrefix("/api/v1/").Subrouter()

	for _, b := range initBundles(a.dbctx) {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", a.router)
	return a
}

//Run startup app
func (a *App) Run(port string) *App {
	log.Fatal(http.ListenAndServe(port, a.router))
	return a
}

func initBundles(ctx dbcontext.DbContext) []common.Bundle {
	return []common.Bundle{
		controllers.NewTripRouter(ctx),
		controllers.NewDeviceRouter(ctx),
	}
}
