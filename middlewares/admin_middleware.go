package middlewares

import (
	"MohamedStaili/GO_Project_inventaire/pkg/config"
	"MohamedStaili/GO_Project_inventaire/pkg/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

// AdminOnlyMiddleware checks if the user has an admin role
func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get JWT cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "No JWT cookie found", http.StatusUnauthorized)
			return
		}
		var SecretKey = "secret"
		// Parse the JWT token using the claims
		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil // Replace with your actual secret key
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok || claims.ExpiresAt < time.Now().Unix() {
			http.Error(w, "Token expired or invalid", http.StatusUnauthorized)
			return
		}

		// Fetch user from the database based on the claim issuer (user ID)
		var user models.User
		db := config.GetDb()
		err = db.Where("id = ?", claims.Issuer).First(&user).Error
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Check if the user has an admin role
		if user.Role != "admin" {
			http.Error(w, "Forbidden:not admin", http.StatusForbidden)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
