package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseError struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, payload interface{}) error {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(
		payload,
	)
}

func WriteErrorJSON(w http.ResponseWriter, statusCode int, err error) error {
	log.Println(err)
	return WriteJSON(w, statusCode, ResponseError{
		Error: err.Error(),
	})
}
