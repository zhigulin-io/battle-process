package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeSuccessResponse(w http.ResponseWriter, response any) {
	marshaled, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(marshaled)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}

func writeErrorResponse(w http.ResponseWriter, code int, err error, msg string) {
	response := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{
		Error:   err.Error(),
		Message: msg,
	}

	marshaled, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(marshaled)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(code)
}
