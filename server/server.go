package server

import (
	"context"
	"errors"

	"github.com/gorilla/mux"
)

// Items that the server need to connect
type Config struct {
	// Port where it is executed
	Port string
	// Secret key used to generate Tokens
	JWTSecret string
	// Database connection
	DatabaseUrl string
}

// Interface to be considered a server
type Server interface {
	Config() *Config
}

// Element that will handle the server
type Broker struct {
	config *Config
	// It defines the API route
	router mux.Router
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
		router: *mux.NewRouter(),
	}

	return broker, nil
}
