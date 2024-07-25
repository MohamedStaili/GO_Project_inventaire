package middlewares

import (
	utils "MohamedStaili/GO_Project_inventaire/pkg/utiles"
	"net/http"
)

// AdminOnlyMiddleware vérifie si l'utilisateur a un rôle d'administrateur
func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := utils.GetSession(r)
		if err != nil || session.Values["authenticated"] != true || session.Values["role"] != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
