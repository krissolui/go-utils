package httputils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadJSONReq[T any](req *http.Request) (*T, error) {
	var payload T

	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func sendResponse(w http.ResponseWriter, response any, status int, headers ...http.Header) {
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.WriteHeader(status)

	payload, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(payload)
}

func ResponseJSON(w http.ResponseWriter, response any, headers ...http.Header) {
	w.Header().Set("Content-Type", "application/json")

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	sendResponse(w, response, http.StatusAccepted, headers...)
}

func ResponseError(w http.ResponseWriter, response any, status int) {
	w.Header().Set("Content-Type", "application/problem+json")

	sendResponse(w, response, status)
}
