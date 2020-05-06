package core

import (
	"log"
	"net/http"

	"github.com/NicolasDeveloper/tracker-microservices/internal/tracker-api/controllers"
	"github.com/NicolasDeveloper/tracker-microservices/pkg/database/dbcontext"
	"github.com/gorilla/mux"
	"github.com/laibulle/kitties/app/core"
)

//App initialize app
type App struct {
	router *mux.Router
	dbctx  dbcontext.DbContext
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

	for _, b := range initBundles(db) {
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

func initBundles(ctx dbcontext.DbContext) []core.Bundle {
	return []core.Bundle{controllers.NewTripRouter(ctx)}
}
