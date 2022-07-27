package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ChrisCodeX/REST-API-Go/models"
	"github.com/ChrisCodeX/REST-API-Go/repository"
	"github.com/ChrisCodeX/REST-API-Go/server"
	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
)

// Receive from the client
type InsertPostRequest struct {
	PostContent string `json:"post_content"`
}

// Response to the client
type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from Authorization header
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

		// Check the validation of the token
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Validation of the user
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			// Decode json into into the struct
			var postRequest = InsertPostRequest{}

			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Generate the id to the post
			id, err := ksuid.NewRandom()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			post := models.Post{
				Id:          id.String(),
				PostContent: postRequest.PostContent,
				UserId:      claims.UserId,
			}

			// Insert the post into the database
			err = repository.InsertPost(r.Context(), &post)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Response to the client
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				Id:          post.Id,
				PostContent: post.PostContent,
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
