package lifelog

import (
	"encoding/json"
	"net/http"
	"strings"
)

// GetStringKey replaces the spaces in the given string with '-' inorder to prepare the string URL friendly
func StringKey(s string) string {
	// TODO: need to find what other characters are not url safe and eliminate them
	return strings.ToLower(strings.Replace(s, " ", "-", -1))
}

func WriteResponse(w http.ResponseWriter, statusCode int, contentType string, body interface{}) {
	if contentType == "" {
		contentType = "application/json"
	}
	w.Header().Set("Content-Type", contentType)

	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
