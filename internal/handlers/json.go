package handlers

import (
	"encoding/json"
	"net/http"
)


func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteJsonError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}

	return WriteJson(w, status, &envelope{Error: message})
}

func JsonResponse(w http.ResponseWriter, r *http.Request, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}
	return WriteJson(w, status, &envelope{Data: data})
}