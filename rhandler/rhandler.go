package rhandler

import (
	"encoding/json"
	"net/http"

	"github.com/Antony15/goodWorkLabs-Test/location"
	"github.com/Antony15/goodWorkLabs-Test/utils"
)

// FindLocations function handler used to find the Parking spots, Charging Stations and Restaurants near the user provided location
func FindLocations(w http.ResponseWriter, r *http.Request) {
	// create header
	w.Header().Add("Content-Type", "application/json")
	loc := location.New()
	ResponseArray := make(map[string]interface{})
	// Decode the request json to location struct
	if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
		// if decoding failed, return error response
		ResponseArray["message"] = "Error : Request Error"
		utils.PrintMessage(w, ResponseArray, http.StatusBadRequest)
		return
	} else {
		// Validate the incoming request for required fields
		if err = loc.ValidateRequest(); err != nil {
			// if missing required fields, retuen error response
			ResponseArray["message"] = err.Error()
			utils.PrintMessage(w, ResponseArray, http.StatusBadRequest)
			return
		} else {
			// call SendRequest function
			if res, ok := loc.SendRequest(); ok {
				// if no error, return response with status code 200
				utils.PrintMessage(w, res, http.StatusOK)
				return
			}
		}
	}
	return
}
