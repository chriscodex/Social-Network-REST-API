package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ChrisCodeX/REST-API-Go/server"
)

// Struct that is returned to the client
type HomeResponse struct {
	Message string `json: "message"`
	Status  bool   `json: "status"`
}

/*
* Handler of the Home Endpoint

* @param {server.Server} s server

* @return {http.HandlerFunc} Handler that calls function
 */

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Http request header, it indicates "w" will respond with a json
		w.Header().Set("Content-Type", "application/json")
		// Status of the http request
		w.WriteHeader(http.StatusOK)
		// Create the response
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome the GO API",
			Status:  true,
		})
	}
}
