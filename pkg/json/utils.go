package http

import (
	"encoding/json"
	"net/http"
)

// WriteJSON - Serializes JSON and set it as output stream using http.ResponseWriter
func WriteJSON(o interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
