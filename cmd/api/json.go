package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
func readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	max_Bytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(max_Bytes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}
	return writeJSON(w, status, envelope{Error: message})
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) error {
	type envelope struct {
		Data any `json:"data"`
	}
	return writeJSON(w, status, envelope{Data: data})
}
