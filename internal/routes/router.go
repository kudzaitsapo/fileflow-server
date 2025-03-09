package routes

import (
	"net/http"
	"sort"
	"strings"

	"github.com/kudzaitsapo/fileflow-server/internal/handlers"
)

type Route struct {
	Pattern string
	Handler http.Handler
	RequiresAuth bool
}

// Define a helper function to create OPTIONS handlers
func createOptionsHandler(allowedMethods []string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust according to your security requirements
        w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ", "))
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // Respond to OPTIONS request
        w.WriteHeader(http.StatusNoContent)
    }
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

	// Process existing routes to collect allowed methods per path
	methodMap := make(map[string]map[string]struct{}) // path -> methods set
	for _, route := range routes {
		parts := strings.SplitN(route.Pattern, " ", 2)
		if len(parts) != 2 {
			continue // handle invalid pattern format if needed
		}
		method, path := parts[0], parts[1]
		
		if _, exists := methodMap[path]; !exists {
			methodMap[path] = make(map[string]struct{})
		}
		methodMap[path][method] = struct{}{}
	}

	for path, methods := range methodMap {
		// Collect allowed methods
		allowedMethods := make([]string, 0, len(methods))
		for m := range methods {
			allowedMethods = append(allowedMethods, m)
		}
		
		// Add OPTIONS method to the allowed list if needed (not required by CORS spec)
		// allowedMethods = append(allowedMethods, "OPTIONS")
		
		// Sort for consistent output
		sort.Strings(allowedMethods)
		
		// Add OPTIONS route
		routes = append(routes, Route {
			Pattern:      "OPTIONS " + path,
			Handler:      createOptionsHandler(allowedMethods),
			RequiresAuth: false,
		})
	}

	return routes
}