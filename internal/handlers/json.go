package handlers

import (
	"encoding/json"
	"net/http"
)

type JsonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JsonMeta struct {
	TotalRecords int64 `json:"total_records"`
	Limit        int64 `json:"limit"`
	Offset       int64 `json:"offset"`
}

type JsonEnvelope struct {
	Result  any  `json:"result"`
	Success bool `json:"success"`
	Error   any  `json:"error"`
	Meta    any  `json:"meta"`
}

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
	jsonError := &JsonError{
		Code:    status,
		Message: message,
	}
	jsonResult := &JsonEnvelope{
		Result:  nil,
		Success: false,
		Error:   *jsonError,
		Meta:    nil,
	}
	return WriteJson(w, status, jsonResult)
}

func SendJson(w http.ResponseWriter, status int, data any, meta any) error {
	response := &JsonEnvelope{
		Result:  data,
		Success: true,
		Error:   nil,
		Meta:    meta,
	}
	return WriteJson(w, status, response)
}

func SendJsonWithoutMeta(w http.ResponseWriter, status int, data any) {
	SendJson(w, status, data, nil)
}
