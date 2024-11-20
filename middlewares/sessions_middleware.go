package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// AuthRequired middleware checks for JWT authentication
func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Exclude certain pages from authentication check
		nonAuthPaths := []string{"/static/login.html", "/static/css", "/static/img", "/static/favicon"} // Add other public pages if needed
		for _, path := range nonAuthPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				next.ServeHTTP(w, r)
				return
			}
		}
		var SecretKey = "secret"
		// Retrieve and parse the JWT cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "No JWT cookie found", http.StatusUnauthorized)
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil // Replace with your actual secret key
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Could not parse claims", http.StatusUnauthorized)
			return
		}

		// Check if the token is expired
		if exp, ok := claims["exp"].(float64); !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		// Continue with the request
		next.ServeHTTP(w, r)
	})
}
