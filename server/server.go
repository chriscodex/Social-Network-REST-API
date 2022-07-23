package server

import "github.com/gorilla/mux"

// Items that the server need to connect
type Config struct {
	// Port where it is executed
	Port int
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
