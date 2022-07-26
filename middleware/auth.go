package middleware

import (
	"net/http"
	"strings"

	"github.com/ChrisCodeX/REST-API-Go/models"
	"github.com/ChrisCodeX/REST-API-Go/server"
	"github.com/golang-jwt/jwt"
)

var (
	// Endpoints that dont need the middleware
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

/*	Function to check if the middleware should check the token of the endpoint

 */
func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

/* Middleware of Check Authentication

 */
func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) {
				// Indicates the endpoint doesn't need the middleware (next handler)
				next.ServeHTTP(w, r)
				return
			}

			// Get the token from Authorization header
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

			// Check the validation of the token
			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
