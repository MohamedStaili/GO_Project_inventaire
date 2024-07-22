package utils

import (
	"fmt"
	"net/http"

	"MohamedStaili/GO_Project_inventaire/pkg/models"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "session-name")
}

func SetSession(w http.ResponseWriter, r *http.Request, user models.User) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID
	fmt.Printf("Setting session for user ID: %d\n", user.ID)
	session.Save(r, w)

	return nil
}

func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	session.Values["authenticated"] = false
	session.Save(r, w)

	return nil
}
