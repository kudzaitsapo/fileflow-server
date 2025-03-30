package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/kudzaitsapo/fileflow-server/cmd/app"
	"github.com/kudzaitsapo/fileflow-server/internal/handlers"
)

type MiddleWare struct {
	Handler func(http.Handler) http.Handler
}

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

// WriteHeader captures the status code before writing it
func (crw *CustomResponseWriter) WriteHeader(code int) {
	crw.StatusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

// Write captures the response body
func (crw *CustomResponseWriter) Write(body []byte) (int, error) {
	crw.Body = body
	return crw.ResponseWriter.Write(body)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			log.Printf("Headers: %s", r.Header)
			errResp := handlers.JsonError{
				Code:    http.StatusUnauthorized,
				Message: "Authorization required to access this resource",
			}

			response := handlers.JsonEnvelope{
				Success: false,
				Error:   errResp,
			}

			// Write error response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Expected format: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Printf("Parts: %s", parts)
			errResp := handlers.JsonError{
				Code:    http.StatusUnauthorized,
				Message: "Invalid authorization header",
			}

			response := handlers.JsonEnvelope{
				Success: false,
				Error:   errResp,
			}

			// Write error response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		token := parts[1]

		// Validate the token
		currentApp := app.GetCurrentApplication()
		authenticator := currentApp.Authenticator
		_, err := authenticator.ValidateToken(token)
		if err != nil {
			log.Printf("Error validating token: %s", err)
			errResp := handlers.JsonError{
				Code:    http.StatusUnauthorized,
				Message: fmt.Sprintf("Error validating token: %s", err),
			}

			response := handlers.JsonEnvelope{
				Success: false,
				Error:   errResp,
			}

			// Write error response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// TODO: Add error handling middleware here
func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		crw := &CustomResponseWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK, // Default status code
		}

		defer func() {
			if err := recover(); err != nil {
				// Log the stack trace
				log.Printf("PANIC: %v\n%s", err, debug.Stack())

				// Create error response
				errResp := handlers.JsonError{
					Code:    http.StatusInternalServerError,
					Message: "An unexpected error occurred",
				}

				response := handlers.JsonEnvelope{
					Success: false,
					Error:   errResp,
				}

				// Write error response
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)
			}
		}()

		next.ServeHTTP(w, r)

		if crw.StatusCode >= 400 {
			// Default error message
			message := "An error occurred"
			errorDetail := http.StatusText(crw.StatusCode)

			// Try to parse existing response if it exists
			var existingResponse map[string]interface{}
			if len(crw.Body) > 0 {
				if err := json.Unmarshal(crw.Body, &existingResponse); err == nil {
					// Extract custom error message if available
					if msg, ok := existingResponse["message"].(string); ok {
						message = msg
					}
					if errDetail, ok := existingResponse["error"].(string); ok {
						errorDetail = errDetail
					}
				}
			}

			// Create standardized error response
			errResp := handlers.JsonError{
				Code:    crw.StatusCode,
				Message: message,
			}

			resp := handlers.JsonEnvelope{
				Success: false,
				Error:   errResp,
			}

			log.Printf("Error on %s: %s", r.URL.Path, errorDetail)

			// Write standardized response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(crw.StatusCode) // Set the original error status code
			json.NewEncoder(w).Encode(resp)
		}
	})
}

func GetMiddlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		LoggingMiddleware,
		CORS,
		ErrorHandlingMiddleware,
	}
}
