package middlewares

import (
	utils "MohamedStaili/GO_Project_inventaire/pkg/utiles"
	"net/http"
)

// AuthRequired vérifie si une session est active
func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := utils.GetSession(r)
		if err != nil || session.Values["authenticated"] == nil || !session.Values["authenticated"].(bool) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
