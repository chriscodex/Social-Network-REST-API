package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/ChrisCodeX/REST-API-Go/database"
	"github.com/ChrisCodeX/REST-API-Go/repository"
	"github.com/gorilla/mux"
)

// Items that the server need to connect
type Config struct {
	Port        string // Port where it is executed
	JWTSecret   string // Secret key used to generate Tokens
	DatabaseUrl string // Database connection
}

// Interface to be considered a server
type Server interface {
	Config() *Config
}

// Element that will handle the server
type Broker struct {
	config *Config
	router *mux.Router // It defines the API route
}

// Method that makes the broker a server interface
func (b *Broker) Config() *Config {
	return b.config
}

// Constructor of Server
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	// Validations
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

// Method that makes the server (Broker) able to start
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	log.Println("Server started on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
