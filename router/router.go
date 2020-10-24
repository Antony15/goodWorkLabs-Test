package router

import (
	"log"
	"net/http"

	"github.com/Antony15/goodWorkLabs-Test/constants"
	"github.com/Antony15/goodWorkLabs-Test/logger"
	"github.com/Antony15/goodWorkLabs-Test/rhandler"
	"github.com/gorilla/mux"
)

// SetupHandlers Function for handling routes
func SetupHandlers() {
	logger.Log.Println("Server started at " + constants.AppProtocol + constants.AppHost + constants.AppPort)
	log.Println("Server started at " + constants.AppProtocol + constants.AppHost + constants.AppPort)
	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc(constants.IDEndPoint, rhandler.FindLocations).Methods("POST")
	if err := http.ListenAndServe(constants.AppPort, r); err != nil {
		panic(err)
	}
}
