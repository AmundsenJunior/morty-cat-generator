package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// the App struct hold the app routing instance
type App struct {
	Router *mux.Router
}

// initialize router of app
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

// initialize routes into the mux router
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/", a.GetMortyFact).Methods("GET")
	a.Router.HandleFunc("/health", a.GetHealthStatus).Methods("GET")
}

func (a *App) Run(addr string) {
	logger.Fatal(http.ListenAndServe(addr, a.Router))
}
