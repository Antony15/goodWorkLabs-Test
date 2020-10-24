package utils

import (
	"encoding/json"
	"net/http"
)

// PrintMessage Function for printing response message
func PrintMessage(w http.ResponseWriter, ResponseArray interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ResponseArray)
	return
}
