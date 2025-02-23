package routes

import (
	"net/http"

	"github.com/kudzaitsapo/fileflow-server/internal/handlers"
)

type Route struct {
	Pattern string
	Handler http.HandlerFunc
}

func CreateRoutes() []Route {
	routes := []Route{}

	routes = append(routes, Route{
		Pattern: "GET /v1/health-check",
		Handler: handlers.HandleHealthCheck,

	})

	return routes
}