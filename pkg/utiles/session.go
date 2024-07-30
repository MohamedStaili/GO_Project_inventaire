package utils

import (
	"fmt"
	"net/http"
	"time"

	"MohamedStaili/GO_Project_inventaire/pkg/models"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "session-name")
}

func SetSession(w http.ResponseWriter, r *http.Request, user models.User, UUID string) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID
	session.Values["role"] = user.Role
	session.Values["username"] = user.Username
	session.Values["email"] = user.Email
	session.Values["name"] = user.Username
	session.Values["uuid"] = UUID
	fmt.Printf("Setting session for user ID: %d\n", user.ID)
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(2 * time.Hour / time.Second),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	session.Save(r, w)

	return nil
}

func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}
