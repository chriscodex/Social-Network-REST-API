package server

// Items that the server will need to connect
type Config struct {
	// Port where it will be executed
	Port int
	// Secret key that will be used to generate Tokens
	JWTSecret string
	// Database connection
	DatabaseUrl string
}

// Interface to be considered a server
type Server interface {
	Config() *Config
}
