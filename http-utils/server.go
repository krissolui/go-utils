package httputils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadJSON[T any](req *http.Request) (*T, error) {
	var payload T

	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func ResponseJSON(w http.ResponseWriter, response any, status int, headers ...http.Header) {
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	payload, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(payload)
}

func ResponseError(w http.ResponseWriter, response any, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	ResponseJSON(w, response, statusCode)
}
