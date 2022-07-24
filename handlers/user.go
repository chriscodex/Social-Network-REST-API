package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ChrisCodeX/REST-API-Go/server"
)

// Items necessary for the registration of a user
type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response
type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
