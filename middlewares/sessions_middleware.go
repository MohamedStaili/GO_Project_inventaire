package middlewares

import (
	utils "MohamedStaili/GO_Project_inventaire/pkg/utiles"
	"net/http"
	"strings"
)

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

		session, err := utils.GetSession(r)
		if err != nil || session.Values["authenticated"] == nil || !session.Values["authenticated"].(bool) {
			http.Redirect(w, r, "/static/login.html", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
