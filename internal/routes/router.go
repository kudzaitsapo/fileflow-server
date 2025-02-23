package routes

import (
	"net/http"

	"github.com/kudzaitsapo/fileflow-server/internal/handlers"
)

type Route struct {
	Pattern string
	Handler http.Handler
	RequiresAuth bool
}

func CreateRoutes() []Route {
	routes := []Route{}

	routes = append(routes, Route{
		Pattern: "GET /v1/health-check",
		Handler: http.HandlerFunc(handlers.HandleHealthCheck),
		RequiresAuth: true,
	},
	Route{
		Pattern: "POST /v1/auth/login",
		Handler: http.HandlerFunc(handlers.HandleAuthentication),
		RequiresAuth: false,
	})
	return routes
}