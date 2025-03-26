package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, v interface{}) error {
	if r == nil {
		return errors.New("request is nil")
	}
	r.Body = http.MaxBytesReader(nil, r.Body, 1048576)
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}
	defer r.Body.Close()
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&v); err != nil {
		return fmt.Errorf("error encoding data: %w", err)
	}
	return nil
}

// WriteError or http.Error
func WriteError(w http.ResponseWriter, status int, err error) error {
	errResp := struct {
		Error string `json:"error"`
	}{Error: err.Error()}
	return WriteJSON(w, status, errResp)
}
