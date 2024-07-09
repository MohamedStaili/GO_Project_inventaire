package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody lit le corps de la requête HTTP et désérialise en une structure Go
func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// Fermer le corps de la requête après lecture pour éviter les fuites de ressources
	defer r.Body.Close()

	if err := json.Unmarshal(body, x); err != nil {
		return err
	}

	return nil
}
