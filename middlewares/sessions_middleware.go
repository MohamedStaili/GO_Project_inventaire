package middlewares

import (
	utils "MohamedStaili/GO_Project_inventaire/pkg/utiles"
	"net/http"
	"strings"
)

// AuthRequired vérifie si une session est active
func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Exclure certaines pages de la vérification d'authentification
		nonAuthPaths := []string{"/static/login.html", "/static/css", "/static/img", "/static/favicon"} // Ajoutez d'autres pages publiques si nécessaire
		for _, path := range nonAuthPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				next.ServeHTTP(w, r)
				return
			}
		}

		session, err := utils.GetSession(r)
		if err != nil || session.Values["authenticated"] == nil || !session.Values["authenticated"].(bool) {
			http.Redirect(w, r, "/static/login.html", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
