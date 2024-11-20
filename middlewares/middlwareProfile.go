package middlewares

import (
	"MohamedStaili/GO_Project_inventaire/pkg/models"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func AuthProfile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Exclude certain paths from authentication
		nonAuthPaths := []string{"/static/login.html", "/static/css", "/static/img", "/static/favicon"}
		for _, path := range nonAuthPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Retrieve the JWT cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "No JWT cookie found", http.StatusUnauthorized)
			return
		}

		// Parse the JWT token
		token, err := jwt.ParseWithClaims(cookie.Value, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil // Replace "secret" with your actual secret key
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, ok := token.Claims.(*models.CustomClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Check if the token is expired
		if claims.ExpiresAt < time.Now().Unix() {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		// Pass the user ID to the request context
		ctx := context.WithValue(r.Context(), "userID", claims.Issuer) // Issuer is a string here
		
		r = r.WithContext(ctx)

		// Continue with the request
		next.ServeHTTP(w, r)
	})
}
