package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kudzaitsapo/fileflow-server/cmd/app"
)

type MiddleWare struct {
	Handler func(http.Handler) http.Handler
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
			http.Error(w, "Authorization required to access", http.StatusUnauthorized)
			return
		}

		// Expected format: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Printf("Parts: %s", parts)
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		// Validate the token
		currentApp := app.GetCurrentApplication()
		authenticator := currentApp.Authenticator
		_, err := authenticator.ValidateToken(token)
		if err != nil {
			log.Printf("Error validating token: %s", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}



func GetMiddlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		LoggingMiddleware,
		CORS,
	}
}