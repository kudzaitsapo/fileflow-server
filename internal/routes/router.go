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
	},
	Route{
		Pattern: "POST /v1/projects",
		Handler: http.HandlerFunc(handlers.HandleProjectCreation),
		RequiresAuth: true,
	},
	Route{
		Pattern: "GET /v1/projects",
		Handler: http.HandlerFunc(handlers.HandleProjectList),
		RequiresAuth: true,
	},
	Route{
		Pattern: "GET /v1/projects/{id}/files",
		Handler: http.HandlerFunc(handlers.HandleFilesList),
		RequiresAuth: true,
	},
	Route{
		Pattern: "POST /v1/files",
		Handler: http.HandlerFunc(handlers.HandleFileUpload),
		RequiresAuth: false,
	},
	Route{
		Pattern: "GET /v1/files/{id}/download",
		Handler: http.HandlerFunc(handlers.HandleFileDownload),
		RequiresAuth: false,
	},
	Route{
		Pattern: "GET /v1/files/{id}/info",
		Handler: http.HandlerFunc(handlers.HandleFileInfo),
		RequiresAuth: true,
	})
	return routes
}