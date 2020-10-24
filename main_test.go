package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Antony15/goodWorkLabs-Test/rhandler"
)

// FindLocations testing function
// Output must contain 3 closest POI’s of each type in the response
func TestFindLocations(t *testing.T) {
	var jsonStr = []byte(`{"latitude":"37.7942","longitude":"-122.4070"}`)

	req, err := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rhandler.FindLocations)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
