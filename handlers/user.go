package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ChrisCodeX/REST-API-Go/models"
	"github.com/ChrisCodeX/REST-API-Go/repository"
	"github.com/ChrisCodeX/REST-API-Go/server"
	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	HASH_COST             = 8
	TOKEN_TIME_EXPIRATION = 2 * time.Hour * 24
)

// Items necessary for the registration of a user
type SignUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response to the client on Signup
type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

// Response to the client on Login
type LoginResponse struct {
	Token string `json:"token"`
}

// Register Handler Function
func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request into the request struct
		var request = SignUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Encrypt password (HASH)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASH_COST)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Generate a random id
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert request data in user struct
		var user = models.User{
			Email:    request.Email,
			Password: string(hashedPassword),
			Id:       id.String(),
		}

		// Send the struct to be stored in the database
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Response to the client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
		// Message on the server
		log.Println("User registered successfully")
	}
}

func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the json request into the request struct
		var request = SignUpLoginRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)

		// Client error
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Send the email to get the user
		user, err := repository.GetUserByEmail(r.Context(), request.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		// It response the same error for security
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			log.Printf("failed to compare hash and password: %s", err)
			return
		}

		claims := models.AppClaims{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(TOKEN_TIME_EXPIRATION).Unix(),
			},
		}

		// Generate the token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Sign the token
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("failed to sign the token: %s", err)
			return
		}

		// Response to the client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResponse{
			Token: tokenString,
		})
	}
}
