package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ChrisCodeX/REST-API-Go/handlers"
	"github.com/ChrisCodeX/REST-API-Go/middleware"
	"github.com/ChrisCodeX/REST-API-Go/server"
	"github.com/ChrisCodeX/REST-API-Go/websocket"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// Create New Server
	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Start the server
	s.Start(BindRoutes)
}

/*
Binder of endpoints

@ param {Server} Server

@ param {Router} Route Handler
*/
func BindRoutes(s server.Server, r *mux.Router) {
	// Creating Hub for upgrade to Websocket
	hub := websocket.NewHub()

	/* HTTP */
	// Assigning Middleware
	r.Use(middleware.CheckAuthMiddleware(s))

	// Endpoint "/"
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	// Endpoint "/signup"
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)

	// Endpoint "/login"
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	// Endpoint "/me"
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)

	/* Endpoints "/post" */
	// Create New Post
	r.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)

	// Get a Post By Id
	r.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)

	// Update a Post By Id
	r.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)

	// Delete a Post By Id
	r.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)

	// Get All Posts
	// This endpoint can receive 2 query parameter: page & size
	r.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)

	/*WebSocket endpoint*/

	// Websocket ready for receive connections
	go hub.Run()

	// Endpoint wich handle the websocket connection
	r.HandleFunc("/ws", hub.HandleWebSocket)
}
