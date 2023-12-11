package main

import (
	"encoding/json"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a default JSON struct
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// encode the JSON string
	body, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	//format the json for terminal output
	body = append(body, '\n')

	// write the JSON to the response body
	// set the response header to inform client of incoming json
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
